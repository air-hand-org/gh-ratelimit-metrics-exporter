package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"go.uber.org/dig"
	"golang.org/x/sync/errgroup"
)

func init() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	initPrometheus()
}

func registerConstructors() *dig.Container {
	container := dig.New()
	if err := container.Provide(NewStdoutLogger); err != nil {
		log.Fatal(err)
	}
	if err := container.Provide(func() []newClientFunc {
		return newClientFuncs
	}); err != nil {
		log.Fatal(err)
	}
	if err := container.Provide(newClientWithEnv); err != nil {
		log.Fatal(err)
	}
	if err := container.Provide(newGitHubRateLimitsFetcher); err != nil {
		log.Fatal(err)
	}
	return container
}

func main() {
	c := registerConstructors()

	err := c.Invoke(func(rtf RateLimitsFetcher, logger *zerolog.Logger) error {
		signalCtx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer cancel()

		g, ctx := errgroup.WithContext(signalCtx)

		fetch := func() {
			fetchCtx, cancelFetch := context.WithTimeout(ctx, 15*time.Second)
			defer cancelFetch()
			fetchGitHubRateLimit(fetchCtx, rtf, logger)
		}

		// TODO: enable to change port.
		server := &http.Server{
			Addr:    ":8080",
			Handler: newHTTPHandler(),
		}

		g.Go(func() error {
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				return errors.Wrap(err, "failed to start HTTP server")
			}
			return nil
		})

		g.Go(func() error {
			<-ctx.Done()
			logger.Info().Msg("Shutting down server.")

			ctxS, cancelS := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancelS()

			if err := server.Shutdown(ctxS); err != nil {
				return errors.Wrap(err, "failed to shutdown HTTP server")
			}
			logger.Info().Msg("Server shutdown")
			return nil
		})

		g.Go(func() error {
			// fetch GitHub rate limit at the beginning
			fetch()
			// TODO: enable to change interval
			ticker := time.NewTicker(5 * time.Minute)
			defer ticker.Stop()

			for {
				select {
				case <-ctx.Done():
					logger.Info().Msg("Stop fetching GitHub rate limit.")
					return nil
				case <-ticker.C:
					fetch()
				}
			}
		})

		if err := g.Wait(); err != nil {
			logger.Error().Stack().Err(err).Msg("Server stopped with error.")
			return err
		}
		return nil
	})

	if err != nil {
		log.Fatal("Failed to start server\n", err)
	}
}
