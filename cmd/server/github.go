package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/go-github/v85/github"
	"github.com/rs/zerolog"
)

type newClientFunc func(*zerolog.Logger) *github.Client

var newClientFuncs = []newClientFunc{
	newClientWithGitHubApp,
	newClientWithToken,
}

func newClientWithEnv(funcs []newClientFunc, logger *zerolog.Logger) (*github.Client, error) {
	for _, f := range funcs {
		if gh := f(logger); gh != nil {
			return gh, nil
		}
	}
	return nil, errors.New("no GitHub authentication is configured")
}

type GitHubRateLimitsFetcher struct {
	client *github.Client
	logger *zerolog.Logger
}

// explicit compile error check
var _ RateLimitsFetcher = (*GitHubRateLimitsFetcher)(nil)

func (g *GitHubRateLimitsFetcher) Fetch() (*github.RateLimits, error) {
	rateLimits, res, err := g.client.RateLimit.Get(context.Background())
	if err != nil {
		g.logger.Error().Err(err).Msgf("Err: %v", res)
		return nil, fmt.Errorf("%w", err)
	}
	return rateLimits, nil
}

func newGitHubRateLimitsFetcher(client *github.Client, logger *zerolog.Logger) RateLimitsFetcher {
	return &GitHubRateLimitsFetcher{client: client, logger: logger}
}
