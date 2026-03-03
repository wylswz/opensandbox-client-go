package opensandbox

import (
	"net/http"
	"os"

	"github.com/alibaba/opensandbox-client-go/internal/generated/execd"
	"github.com/alibaba/opensandbox-client-go/internal/generated/sandbox"
)

// Clientset is the main client for OpenSandbox APIs, similar to kubernetes.Clientset.
type Clientset struct {
	sandbox *sandboxClient
	execd   *execdClient
}

// NewForConfig creates a new Clientset for the given config.
func NewForConfig(c *Config) (*Clientset, error) {
	if c == nil {
		c = DefaultConfig()
	}
	cs := &Clientset{}

	// Sandbox client
	sandboxCfg := sandbox.NewConfiguration()
	sandboxCfg.Servers = sandbox.ServerConfigurations{{URL: c.SandboxAPIURL, Description: "Sandbox API"}}
	sandboxCfg.HTTPClient = c.HTTPClient
	if sandboxCfg.HTTPClient == nil {
		sandboxCfg.HTTPClient = http.DefaultClient
	}
	sandboxCfg.UserAgent = c.UserAgent
	apiKey := c.APIKey
	if apiKey == "" {
		apiKey = os.Getenv("OPEN_SANDBOX_API_KEY")
	}
	cs.sandbox = &sandboxClient{
		client: sandbox.NewAPIClient(sandboxCfg),
		apiKey: apiKey,
	}

	// Execd client
	execdCfg := execd.NewConfiguration()
	execdCfg.Servers = execd.ServerConfigurations{{URL: c.ExecdAPIURL, Description: "Execd API"}}
	execdCfg.HTTPClient = c.HTTPClient
	if execdCfg.HTTPClient == nil {
		execdCfg.HTTPClient = http.DefaultClient
	}
	execdCfg.UserAgent = c.UserAgent
	cs.execd = &execdClient{
		client:      execd.NewAPIClient(execdCfg),
		accessToken: c.AccessToken,
	}

	return cs, nil
}

// Sandbox returns the sandbox lifecycle API client.
func (c *Clientset) Sandbox() SandboxInterface {
	return c.sandbox
}

// Execd returns the code execution API client.
func (c *Clientset) Execd() ExecdInterface {
	return c.execd
}

// Ensure Interface is implemented.
var _ Interface = (*Clientset)(nil)
