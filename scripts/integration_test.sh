#!/usr/bin/env bash
#
# Integration test script for OpenSandbox Go client.
# Starts OpenSandbox server via Docker, runs integration tests, then cleans up.
#
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(cd "${SCRIPT_DIR}/.." && pwd)"
CONTAINER_NAME="opensandbox-integration-test"
SERVER_PORT="${OPEN_SANDBOX_SERVER_PORT:-8090}"
SANDBOX_URL="http://localhost:${SERVER_PORT}/v1"

log() { echo "[integration-test] $*"; }

# Check Docker is available
if ! command -v docker &>/dev/null; then
    log "ERROR: Docker is required for integration tests. Install Docker and try again."
    exit 1
fi

# If OPEN_SANDBOX_SANDBOX_URL is set, skip starting server (use external server)
if [[ -n "${OPEN_SANDBOX_SANDBOX_URL:-}" ]]; then
    log "Using external server: $OPEN_SANDBOX_SANDBOX_URL"
    cd "${ROOT_DIR}"
    go test -tags=integration ./pkg/opensandbox/... -v -count=1
    exit 0
fi

# Cleanup on exit
cleanup() {
    local exit_code=$?
    if [[ ${exit_code} -ne 0 ]] && docker ps -a --format '{{.Names}}' 2>/dev/null | grep -q "^${CONTAINER_NAME}$"; then
        log "Tests failed (exit ${exit_code}). Dumping OpenSandbox server logs..."
        docker logs "${CONTAINER_NAME}" 2>&1 | tail -200
        LOG_FILE="${ROOT_DIR}/opensandbox-server.log"
        docker logs "${CONTAINER_NAME}" 2>&1 > "${LOG_FILE}"
        log "Full logs saved to ${LOG_FILE}"
    fi
    log "Stopping OpenSandbox server..."
    docker rm -f "${CONTAINER_NAME}" 2>/dev/null || true
}
trap cleanup EXIT

# Stop any existing container
docker rm -f "${CONTAINER_NAME}" 2>/dev/null || true

# Create config for server (no api_key = auth disabled)
CONFIG_FILE="${ROOT_DIR}/scripts/integration-test-config.toml"
if [[ ! -f "${CONFIG_FILE}" ]]; then
    log "ERROR: Config file not found: ${CONFIG_FILE}"
    exit 1
fi

# Start OpenSandbox server
log "Starting OpenSandbox server on port ${SERVER_PORT}..."
docker run -d \
    --name "${CONTAINER_NAME}" \
    -p "${SERVER_PORT}:${SERVER_PORT}" \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -v "${CONFIG_FILE}:/etc/opensandbox/config.toml:ro" \
    -e SANDBOX_CONFIG_PATH=/etc/opensandbox/config.toml \
    opensandbox/server:latest

# Wait for server to be ready
log "Waiting for server to be ready..."
for i in $(seq 1 30); do
    if curl -sf "http://localhost:${SERVER_PORT}/health" &>/dev/null; then
        log "Server is ready"
        break
    fi
    if [[ $i -eq 30 ]]; then
        log "ERROR: Server did not become ready in time"
        docker logs "${CONTAINER_NAME}" 2>&1 | tail -50
        exit 1
    fi
    sleep 2
done

# Run integration tests
export OPEN_SANDBOX_SANDBOX_URL="${SANDBOX_URL}"
export OPEN_SANDBOX_API_KEY=""  # Auth disabled in config

log "Running integration tests..."
cd "${ROOT_DIR}"
go test -tags=integration ./pkg/opensandbox/... -v -count=1 -timeout 5m

log "Integration tests passed"
