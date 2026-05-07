SHELL := bash

PROJECT_NAME := gh-ratelimit-metrics-exporter
GITHUB_REPOSITORY_OWNER ?= air-hand-org
export GITHUB_REPOSITORY_OWNER

.PHONY: build
build:
	goreleaser release --snapshot --clean

.PHONY: run
run:
	go run ./app

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test -v --cover ./...

.PHONY: e2e-test
e2e-test: PORT ?= 8080
e2e-test:
	docker stop $(PROJECT_NAME) || true
	docker run -d --rm -p $(PORT):8080 --env GH_TOKEN=$${GH_TOKEN} --name=$(PROJECT_NAME) ghcr.io/$${GITHUB_REPOSITORY_OWNER}/$(PROJECT_NAME):latest-amd64 && \
	wait-for localhost:$(PORT) -t 30  && \
	curl -fsS -X GET http://localhost:$(PORT)/metrics | promtool check metrics --extended
	docker stop $(PROJECT_NAME) || true
