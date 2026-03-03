# OpenSandbox Go Client

Go client for the [OpenSandbox](https://github.com/wylswz/OpenSandbox) API, designed in the style of [kubernetes/client-go](https://github.com/kubernetes/client-go).

## Client structure

The client contains two API groups:

- **Sandbox** – lifecycle management (create, list, get, delete, pause, resume, renew, endpoints)
- **Execd** – code execution (Code, Command, Filesystem, Metrics, Health)

## Installation

```bash
go get github.com/wylswz/opensandbox-client-go
```

## Usage

```go
import (
    "context"
    "log"

    "github.com/wylswz/opensandbox-client-go/internal/generated/execd"
    "github.com/wylswz/opensandbox-client-go/internal/generated/sandbox"
    "github.com/wylswz/opensandbox-client-go/pkg/opensandbox"
)

func main() {
    cfg := opensandbox.DefaultConfig()
    cfg.SandboxAPIURL = "http://localhost:8080/v1"
    cfg.APIKey = "your-api-key" // or set OPEN_SANDBOX_API_KEY env

    cs, err := opensandbox.NewForConfig(cfg)
    if err != nil {
        log.Fatal(err)
    }

    ctx := context.Background()

    // List sandboxes
    list, err := cs.Sandbox().List(ctx, &opensandbox.ListOptions{PageSize: 10})

    // Create sandbox
    resp, err := cs.Sandbox().Create(ctx, &sandbox.CreateSandboxRequest{
        Image:          sandbox.ImageSpec{Uri: "python:3.11"},
        Timeout:        3600,
        ResourceLimits: map[string]string{"cpu": "500m", "memory": "512Mi"},
        Entrypoint:     []string{"python", "-c", "print('hello')"},
    })

    // Execd (code execution) - requires execd URL and access token
    cfg.ExecdAPIURL = "http://localhost:44772"
    cfg.AccessToken = "sandbox-access-token"
    cs, _ = opensandbox.NewForConfig(cfg)
    lang := "python"
    ctxObj, err := cs.Execd().Code().CreateContext(ctx, execd.CodeContextRequest{Language: &lang})
}
```

## Code generation

Generated code lives in `internal/generated/`. Regenerate with:

```bash
make generate
```

This downloads the OpenAPI generator and specs from scratch, then generates Go clients for both APIs.

## Testing

```bash
# Unit tests
make test

# Integration tests (requires Docker)
make test-integration
```

Integration tests start an OpenSandbox server via Docker and exercise sandbox lifecycle (List, Create, Get, Delete, GetEndpoint). To use an external server instead:

```bash
export OPEN_SANDBOX_SANDBOX_URL=http://localhost:8080/v1
export OPEN_SANDBOX_API_KEY=your-api-key  # if auth enabled
make test-integration
```