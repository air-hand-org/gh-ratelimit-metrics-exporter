FROM scratch
ARG TARGETPLATFORM
COPY $TARGETPLATFORM/gh-ratelimit-metrics-exporter /app
ENTRYPOINT ["/app"]
