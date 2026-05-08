FROM gcr.io/distroless/static-debian13:nonroot
ARG TARGETPLATFORM
COPY $TARGETPLATFORM/gh-ratelimit-metrics-exporter /app
ENTRYPOINT ["/app"]
