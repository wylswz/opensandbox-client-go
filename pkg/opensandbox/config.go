package opensandbox

import (
	"net/http"
)

// Config holds the configuration for connecting to OpenSandbox APIs.
// Similar to k8s rest.Config - it separates loading configuration from using it.
type Config struct {
	// SandboxAPIURL is the base URL for the sandbox lifecycle API (e.g., http://localhost:8080/v1).
	SandboxAPIURL string
	// ExecdAPIURL is the base URL for the execd (code execution) API (e.g., http://localhost:44772).
	ExecdAPIURL string
	// APIKey is used for sandbox lifecycle API authentication (OPEN-SANDBOX-API-KEY header).
	// If empty, OPEN_SANDBOX_API_KEY environment variable is used.
	APIKey string
	// AccessToken is used for execd API authentication (X-EXECD-ACCESS-TOKEN header).
	// Required when using execd operations.
	AccessToken string
	// HTTPClient is the HTTP client to use. If nil, http.DefaultClient is used.
	HTTPClient *http.Client
	// UserAgent is sent with each request.
	UserAgent string
}

// DefaultConfig returns a Config with default local development URLs.
func DefaultConfig() *Config {
	return &Config{
		SandboxAPIURL: "http://localhost:8080/v1",
		ExecdAPIURL:   "http://localhost:44772",
		UserAgent:     "opensandbox-client-go/1.0",
	}
}
