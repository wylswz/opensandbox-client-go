#!/usr/bin/env bash
#
# Code generation script for OpenSandbox Go client.
# Builds everything from scratch: downloads OpenAPI generator and specs, then generates Go client code.
#
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(cd "${SCRIPT_DIR}/.." && pwd)"
CACHE_DIR="${ROOT_DIR}/.cache"
SPECS_DIR="${ROOT_DIR}/specs"
GENERATED_DIR="${ROOT_DIR}/pkg/generated"

# OpenAPI Generator version (pinned for reproducibility)
OPENAPI_GENERATOR_VERSION="${OPENAPI_GENERATOR_VERSION:-7.14.0}"
OPENAPI_GENERATOR_JAR="openapi-generator-cli-${OPENAPI_GENERATOR_VERSION}.jar"
OPENAPI_GENERATOR_URL="https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/${OPENAPI_GENERATOR_VERSION}/${OPENAPI_GENERATOR_JAR}"

# OpenSandbox spec URLs
OPENSANDBOX_SPEC_BASE="https://raw.githubusercontent.com/alibaba/OpenSandbox/main/specs"
SANDBOX_LIFECYCLE_SPEC="${OPENSANDBOX_SPEC_BASE}/sandbox-lifecycle.yml"
EXECD_API_SPEC="${OPENSANDBOX_SPEC_BASE}/execd-api.yaml"

log() { echo "[generate] $*"; }

# Ensure Java is available
ensure_java() {
    if ! command -v java &>/dev/null; then
        log "ERROR: Java is required to run OpenAPI Generator. Please install Java (JRE 8+) and ensure it's on PATH."
        exit 1
    fi
    log "Using Java: $(java -version 2>&1 | head -1)"
}

# Download OpenAPI Generator JAR
download_openapi_generator() {
    mkdir -p "${CACHE_DIR}"
    local jar_path="${CACHE_DIR}/${OPENAPI_GENERATOR_JAR}"

    if [[ -f "${jar_path}" ]]; then
        log "OpenAPI Generator JAR already cached: ${jar_path}"
        return 0
    fi

    log "Downloading OpenAPI Generator v${OPENAPI_GENERATOR_VERSION}..."
    if ! curl -fsSL -o "${jar_path}" "${OPENAPI_GENERATOR_URL}"; then
        log "ERROR: Failed to download OpenAPI Generator from ${OPENAPI_GENERATOR_URL}"
        exit 1
    fi
    log "Downloaded: ${jar_path}"
}

# Download OpenSandbox spec files
download_specs() {
    mkdir -p "${SPECS_DIR}"

    log "Downloading OpenSandbox specs..."
    curl -fsSL -o "${SPECS_DIR}/sandbox-lifecycle.yml" "${SANDBOX_LIFECYCLE_SPEC}"
    curl -fsSL -o "${SPECS_DIR}/execd-api.yaml" "${EXECD_API_SPEC}"
    log "Specs saved to ${SPECS_DIR}/"
}

# Run OpenAPI Generator
run_generator() {
    local spec_file="$1"
    local output_dir="$2"
    local package_name="$3"

    log "Generating ${package_name} client from ${spec_file}..."

    java -ea -Xms512M -Xmx1024M -jar "${CACHE_DIR}/${OPENAPI_GENERATOR_JAR}" generate \
        --input-spec "${spec_file}" \
        --generator-name go \
        --output "${output_dir}" \
        --git-user-id wylswz \
        --git-repo-id opensandbox-client-go \
        --additional-properties=packageName="${package_name}" \
        --additional-properties=withGoMod=false \
        --additional-properties=isGoSubmodule=false \
        --additional-properties=hideGenerationTimestamp=true \
        --openapi-generator-ignore-list=test,docs,git_push.sh,.travis.yml

    log "Generated: ${output_dir}"
}

# Main
main() {
    log "Starting code generation (from scratch)..."
    cd "${ROOT_DIR}"

    ensure_java
    download_openapi_generator
    download_specs

    # Clean previous generated output
    rm -rf "${GENERATED_DIR}"
    mkdir -p "${GENERATED_DIR}"

    # Generate sandbox lifecycle client
    run_generator \
        "${SPECS_DIR}/sandbox-lifecycle.yml" \
        "${GENERATED_DIR}/sandbox" \
        "sandbox"

    # Generate execd (code execution) client
    run_generator \
        "${SPECS_DIR}/execd-api.yaml" \
        "${GENERATED_DIR}/execd" \
        "execd"

    # Remove generated test files (they have incorrect imports for our module layout)
    rm -rf "${GENERATED_DIR}/sandbox/test" "${GENERATED_DIR}/execd/test"

    log "Done. Generated clients:"
    log "  - ${GENERATED_DIR}/sandbox/  (sandbox lifecycle API)"
    log "  - ${GENERATED_DIR}/execd/    (code execution API)"
}

main "$@"
