package opensandbox

import (
	"context"

	"github.com/wylswz/opensandbox-client-go/pkg/generated/sandbox"
)

type sandboxClient struct {
	client *sandbox.APIClient
	apiKey string
}

func (c *sandboxClient) authContext(ctx context.Context) context.Context {
	if c.apiKey == "" {
		return ctx
	}
	return context.WithValue(ctx, sandbox.ContextAPIKeys, map[string]sandbox.APIKey{
		"apiKeyAuth": {Key: c.apiKey},
	})
}

func (c *sandboxClient) Create(ctx context.Context, req *sandbox.CreateSandboxRequest) (*sandbox.CreateSandboxResponse, error) {
	resp, _, err := c.client.SandboxesAPI.SandboxesPost(c.authContext(ctx)).
		CreateSandboxRequest(*req).
		Execute()
	return resp, err
}

func (c *sandboxClient) List(ctx context.Context, opts *ListOptions) (*sandbox.ListSandboxesResponse, error) {
	r := c.client.SandboxesAPI.SandboxesGet(c.authContext(ctx))
	if opts != nil {
		if len(opts.State) > 0 {
			r = r.State(opts.State)
		}
		if opts.Metadata != "" {
			r = r.Metadata(opts.Metadata)
		}
		if opts.Page > 0 {
			r = r.Page(opts.Page)
		}
		if opts.PageSize > 0 {
			r = r.PageSize(opts.PageSize)
		}
	}
	resp, _, err := r.Execute()
	return resp, err
}

func (c *sandboxClient) Get(ctx context.Context, sandboxID string) (*sandbox.Sandbox, error) {
	resp, _, err := c.client.SandboxesAPI.SandboxesSandboxIdGet(c.authContext(ctx), sandboxID).Execute()
	return resp, err
}

func (c *sandboxClient) Delete(ctx context.Context, sandboxID string) error {
	_, err := c.client.SandboxesAPI.SandboxesSandboxIdDelete(c.authContext(ctx), sandboxID).Execute()
	return err
}

func (c *sandboxClient) Pause(ctx context.Context, sandboxID string) error {
	_, err := c.client.SandboxesAPI.SandboxesSandboxIdPausePost(c.authContext(ctx), sandboxID).Execute()
	return err
}

func (c *sandboxClient) Resume(ctx context.Context, sandboxID string) error {
	_, err := c.client.SandboxesAPI.SandboxesSandboxIdResumePost(c.authContext(ctx), sandboxID).Execute()
	return err
}

func (c *sandboxClient) RenewExpiration(ctx context.Context, sandboxID string, req *sandbox.RenewSandboxExpirationRequest) (*sandbox.RenewSandboxExpirationResponse, error) {
	resp, _, err := c.client.SandboxesAPI.SandboxesSandboxIdRenewExpirationPost(c.authContext(ctx), sandboxID).
		RenewSandboxExpirationRequest(*req).
		Execute()
	return resp, err
}

func (c *sandboxClient) GetEndpoint(ctx context.Context, sandboxID string, port int32) (*sandbox.Endpoint, error) {
	return c.GetEndpointWithProxy(ctx, sandboxID, port, false)
}

func (c *sandboxClient) GetEndpointWithProxy(ctx context.Context, sandboxID string, port int32, useServerProxy bool) (*sandbox.Endpoint, error) {
	resp, _, err := c.client.SandboxesAPI.SandboxesSandboxIdEndpointsPortGet(c.authContext(ctx), sandboxID, port).
		UseServerProxy(useServerProxy).
		Execute()
	return resp, err
}
