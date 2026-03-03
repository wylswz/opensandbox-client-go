.PHONY: generate clean test test-integration

# Generate Go client code from OpenAPI specs.
# Downloads OpenAPI generator and specs from scratch if needed.
generate:
	@chmod +x scripts/generate.sh
	@./scripts/generate.sh

# Run unit tests (excludes integration tests)
test:
	go test -v -count=1 ./...

# Run integration tests. Requires Docker and OpenSandbox server.
# Set OPEN_SANDBOX_SANDBOX_URL to use an external server; otherwise the script starts one.
test-integration:
	@chmod +x scripts/integration_test.sh
	@./scripts/integration_test.sh

# Remove generated code and caches
clean:
	rm -rf pkg/generated .cache
