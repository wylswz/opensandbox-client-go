package opensandbox

import (
	"net/http"
	"os"
	"strings"
)

const (
	EnvSandboxAPIURL = "OPEN_SANDBOX_SANDBOX_URL"
	EnvExecdAPIURL   = "OPEN_SANDBOX_EXECD_URL"
	EnvAPIKey        = "OPEN_SANDBOX_API_KEY"
	EnvAccessToken   = "OPEN_SANDBOX_EXECD_ACCESS_TOKEN"

	DefaultSandboxAPIURL = "http://localhost:8080/v1"
	DefaultExecdAPIURL   = "http://localhost:44772"
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

func envOr(env, def string) string {
	if v := os.Getenv(env); v != "" {
		return v
	}
	return def
}

func getOrDefault(envMap map[string]string, key, def string) string {
	if v := envMap[key]; v != "" {
		return v
	}
	return def
}

func NewConfigFromEnv() *Config {
	return &Config{
		SandboxAPIURL: envOr(EnvSandboxAPIURL, "http://localhost:8080/v1"),
		ExecdAPIURL:   envOr(EnvExecdAPIURL, "http://localhost:44772"),
		APIKey:        envOr(EnvAPIKey, ""),
		AccessToken:   envOr(EnvAccessToken, ""),
	}
}

func loadEnvFile(path string) (map[string]string, error) {
	env := make(map[string]string)
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if strings.Contains(line, "=") {
			parts := strings.SplitN(line, "=", 2)
			env[parts[0]] = parts[1]
		}
	}
	return env, nil
}

func NewFromEnvFile(path string) (*Config, error) {
	env, err := loadEnvFile(path)
	if err != nil {
		return nil, err
	}
	return &Config{
		SandboxAPIURL: getOrDefault(env, EnvSandboxAPIURL, DefaultSandboxAPIURL),
		ExecdAPIURL:   getOrDefault(env, EnvExecdAPIURL, DefaultExecdAPIURL),
		APIKey:        getOrDefault(env, EnvAPIKey, ""),
		AccessToken:   getOrDefault(env, EnvAccessToken, ""),
	}, nil
}
