SHELL := bash

PROJECT_NAME := gh-ratelimit-metrics-exporter
export PROJECT_NAME

GITHUB_REPOSITORY_OWNER ?= air-hand-org
export GITHUB_REPOSITORY_OWNER

.PHONY: build
build:
	goreleaser release --snapshot --clean

.PHONY: codeql-build
codeql-build:
	go build -v -o /tmp/gh-ratelimit-metrics-exporter-codeql ./cmd/server

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: run
run:
	go run ./cmd/server

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test -v --cover ./...

.PHONY: e2e-test
e2e-test: export PORT ?= 8080
e2e-test: export PROMETHEUS_PORT ?= 9090
e2e-test:
	e2e/run.sh
