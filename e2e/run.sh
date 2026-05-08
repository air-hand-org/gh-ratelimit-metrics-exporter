#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR=$(dirname $(readlink -f $0))

compose_project="${PROJECT_NAME}-e2e"

cleanup() {
    docker compose -p "${compose_project}" down || true
}

dump_debug() {
    echo "=== /metrics ==="
    curl -fsS -X GET "http://localhost:${PORT}/metrics" || true

    echo "=== PromQL: github_rate_limit_core_remaining ==="
    promtool query instant "http://localhost:${PROMETHEUS_PORT}" 'github_rate_limit_core_remaining' || true
}

retry() {
    local -r max_attempts="$1"
    local -r delay="$2"
    shift 2

    local attempt
    for ((attempt = 1; attempt <= max_attempts; attempt++)); do
        if "$@"; then
            return 0
        fi
        sleep "${delay}"
    done

    return 1
}

check_metrics() {
    curl -fsS -X GET "http://localhost:${PORT}/metrics" | promtool check metrics
}

check_rate_limit_metric() {
    promtool query instant "http://localhost:${PROMETHEUS_PORT}" 'github_rate_limit_core_remaining >= 1' |
        grep -q 'github_rate_limit_core_remaining'
}

trap dump_debug ERR
trap cleanup EXIT

cd $SCRIPT_DIR

cleanup

docker compose -p "${compose_project}" up -d

wait-for "localhost:${PORT}" -t 30
wait-for "localhost:${PROMETHEUS_PORT}" -t 30

curl --retry 30 --retry-delay 1 --retry-all-errors -fsS -X GET "http://localhost:${PORT}/healthz"
retry 30 1 check_metrics
retry 30 1 check_rate_limit_metric
