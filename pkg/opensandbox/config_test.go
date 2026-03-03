package opensandbox

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()
	if cfg == nil {
		t.Fatal("DefaultConfig returned nil")
	}
	if cfg.SandboxAPIURL != DefaultSandboxAPIURL {
		t.Errorf("SandboxAPIURL = %q, want %q", cfg.SandboxAPIURL, DefaultSandboxAPIURL)
	}
	if cfg.ExecdAPIURL != DefaultExecdAPIURL {
		t.Errorf("ExecdAPIURL = %q, want %q", cfg.ExecdAPIURL, DefaultExecdAPIURL)
	}
	if cfg.UserAgent != defaultUserAgent {
		t.Errorf("UserAgent = %q, want %q", cfg.UserAgent, defaultUserAgent)
	}
}

func TestNewConfigFromEnv(t *testing.T) {
	// Save and restore env to avoid polluting other tests
	sandboxURL := os.Getenv(EnvSandboxAPIURL)
	execdURL := os.Getenv(EnvExecdAPIURL)
	apiKey := os.Getenv(EnvAPIKey)
	accessToken := os.Getenv(EnvAccessToken)
	defer func() {
		restoreEnv(EnvSandboxAPIURL, sandboxURL)
		restoreEnv(EnvExecdAPIURL, execdURL)
		restoreEnv(EnvAPIKey, apiKey)
		restoreEnv(EnvAccessToken, accessToken)
	}()

	// Test defaults when env is unset
	os.Unsetenv(EnvSandboxAPIURL)
	os.Unsetenv(EnvExecdAPIURL)
	os.Unsetenv(EnvAPIKey)
	os.Unsetenv(EnvAccessToken)
	cfg := NewConfigFromEnv()
	if cfg.SandboxAPIURL != DefaultSandboxAPIURL {
		t.Errorf("SandboxAPIURL = %q, want %q", cfg.SandboxAPIURL, DefaultSandboxAPIURL)
	}
	if cfg.ExecdAPIURL != DefaultExecdAPIURL {
		t.Errorf("ExecdAPIURL = %q, want %q", cfg.ExecdAPIURL, DefaultExecdAPIURL)
	}
	if cfg.APIKey != "" {
		t.Errorf("APIKey = %q, want empty", cfg.APIKey)
	}
	if cfg.AccessToken != "" {
		t.Errorf("AccessToken = %q, want empty", cfg.AccessToken)
	}

	// Test values from env
	os.Setenv(EnvSandboxAPIURL, "https://sandbox.example.com/v1")
	os.Setenv(EnvExecdAPIURL, "https://execd.example.com")
	os.Setenv(EnvAPIKey, "test-api-key")
	os.Setenv(EnvAccessToken, "test-access-token")
	cfg = NewConfigFromEnv()
	if cfg.SandboxAPIURL != "https://sandbox.example.com/v1" {
		t.Errorf("SandboxAPIURL = %q, want https://sandbox.example.com/v1", cfg.SandboxAPIURL)
	}
	if cfg.ExecdAPIURL != "https://execd.example.com" {
		t.Errorf("ExecdAPIURL = %q, want https://execd.example.com", cfg.ExecdAPIURL)
	}
	if cfg.APIKey != "test-api-key" {
		t.Errorf("APIKey = %q, want test-api-key", cfg.APIKey)
	}
	if cfg.AccessToken != "test-access-token" {
		t.Errorf("AccessToken = %q, want test-access-token", cfg.AccessToken)
	}
}

func restoreEnv(key, value string) {
	if value == "" {
		os.Unsetenv(key)
	} else {
		os.Setenv(key, value)
	}
}

func TestNewFromEnvFile(t *testing.T) {
	dir := t.TempDir()

	t.Run("valid file", func(t *testing.T) {
		path := filepath.Join(dir, "valid.env")
		content := `OPEN_SANDBOX_SANDBOX_URL=https://sandbox.test/v1
OPEN_SANDBOX_EXECD_URL=https://execd.test
OPEN_SANDBOX_API_KEY=my-api-key
OPEN_SANDBOX_EXECD_ACCESS_TOKEN=my-token
`
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			t.Fatalf("WriteFile: %v", err)
		}
		cfg, err := NewFromEnvFile(path)
		if err != nil {
			t.Fatalf("NewFromEnvFile: %v", err)
		}
		if cfg.SandboxAPIURL != "https://sandbox.test/v1" {
			t.Errorf("SandboxAPIURL = %q, want https://sandbox.test/v1", cfg.SandboxAPIURL)
		}
		if cfg.ExecdAPIURL != "https://execd.test" {
			t.Errorf("ExecdAPIURL = %q, want https://execd.test", cfg.ExecdAPIURL)
		}
		if cfg.APIKey != "my-api-key" {
			t.Errorf("APIKey = %q, want my-api-key", cfg.APIKey)
		}
		if cfg.AccessToken != "my-token" {
			t.Errorf("AccessToken = %q, want my-token", cfg.AccessToken)
		}
	})

	t.Run("partial file uses defaults", func(t *testing.T) {
		path := filepath.Join(dir, "partial.env")
		content := `OPEN_SANDBOX_API_KEY=partial-key
`
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			t.Fatalf("WriteFile: %v", err)
		}
		cfg, err := NewFromEnvFile(path)
		if err != nil {
			t.Fatalf("NewFromEnvFile: %v", err)
		}
		if cfg.SandboxAPIURL != DefaultSandboxAPIURL {
			t.Errorf("SandboxAPIURL = %q, want default %q", cfg.SandboxAPIURL, DefaultSandboxAPIURL)
		}
		if cfg.ExecdAPIURL != DefaultExecdAPIURL {
			t.Errorf("ExecdAPIURL = %q, want default %q", cfg.ExecdAPIURL, DefaultExecdAPIURL)
		}
		if cfg.APIKey != "partial-key" {
			t.Errorf("APIKey = %q, want partial-key", cfg.APIKey)
		}
	})

	t.Run("empty file uses defaults", func(t *testing.T) {
		path := filepath.Join(dir, "empty.env")
		if err := os.WriteFile(path, []byte(""), 0644); err != nil {
			t.Fatalf("WriteFile: %v", err)
		}
		cfg, err := NewFromEnvFile(path)
		if err != nil {
			t.Fatalf("NewFromEnvFile: %v", err)
		}
		if cfg.SandboxAPIURL != DefaultSandboxAPIURL {
			t.Errorf("SandboxAPIURL = %q, want default %q", cfg.SandboxAPIURL, DefaultSandboxAPIURL)
		}
		if cfg.ExecdAPIURL != DefaultExecdAPIURL {
			t.Errorf("ExecdAPIURL = %q, want default %q", cfg.ExecdAPIURL, DefaultExecdAPIURL)
		}
	})

	t.Run("nonexistent file returns error", func(t *testing.T) {
		path := filepath.Join(dir, "nonexistent.env")
		cfg, err := NewFromEnvFile(path)
		if err == nil {
			t.Fatal("NewFromEnvFile expected error for nonexistent file")
		}
		if cfg != nil {
			t.Errorf("NewFromEnvFile should return nil config on error, got %v", cfg)
		}
	})

	t.Run("value with equals sign", func(t *testing.T) {
		path := filepath.Join(dir, "equals.env")
		content := `OPEN_SANDBOX_API_KEY=key=with=equals
`
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			t.Fatalf("WriteFile: %v", err)
		}
		cfg, err := NewFromEnvFile(path)
		if err != nil {
			t.Fatalf("NewFromEnvFile: %v", err)
		}
		if cfg.APIKey != "key=with=equals" {
			t.Errorf("APIKey = %q, want key=with=equals", cfg.APIKey)
		}
	})
}
