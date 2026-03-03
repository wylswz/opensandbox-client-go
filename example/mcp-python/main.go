// Command opensandbox-mcp-python runs an MCP stdio server that executes
// Python code in a single reusable OpenSandbox sandbox.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/wylswz/opensandbox-client-go/pkg/generated/execd"
	"github.com/wylswz/opensandbox-client-go/pkg/generated/sandbox"
	"github.com/wylswz/opensandbox-client-go/pkg/opensandbox"
)

const (
	serverName    = "opensandbox-mcp-python"
	serverVersion = "0.1.0"

	toolExecutePython = "execute_python"

	defaultToolTimeout = 30 * time.Second
	maxToolTimeout     = 2 * time.Minute

	defaultSandboxImage   = "opensandbox/code-interpreter:v1.0.1"
	defaultSandboxTimeout = int32(3600)
)

type executePythonArgs struct {
	Code           string `json:"code" jsonschema:"Python source code to execute."`
	TimeoutSeconds *int64 `json:"timeout_seconds,omitempty" jsonschema:"Execution timeout in seconds (1-120)."`
}

type sandboxRuntime struct {
	clientset *opensandbox.Clientset
	cfg       *opensandbox.Config

	mu        sync.Mutex
	sandboxID string
	execd     *opensandbox.Clientset
	pyContext *execd.CodeContext
}

func newSandboxRuntime(cfg *opensandbox.Config) (*sandboxRuntime, error) {
	clientset, err := opensandbox.NewForConfig(cfg)
	if err != nil {
		return nil, fmt.Errorf("create clientset: %w", err)
	}
	return &sandboxRuntime{
		clientset: clientset,
		cfg:       cfg,
	}, nil
}

func (r *sandboxRuntime) ensureReady(ctx context.Context) error {
	r.mu.Lock()
	if r.execd != nil && r.pyContext != nil && r.sandboxID != "" {
		r.mu.Unlock()
		return nil
	}
	r.mu.Unlock()

	createCtx, cancel := context.WithTimeout(ctx, 2*time.Minute)
	defer cancel()

	req := &sandbox.CreateSandboxRequest{
		Image:          sandbox.ImageSpec{Uri: defaultSandboxImage},
		Timeout:        defaultSandboxTimeout,
		ResourceLimits: map[string]string{"cpu": "500m", "memory": "512Mi"},
		Entrypoint:     []string{"/opt/opensandbox/code-interpreter.sh"},
		Env:            map[string]string{"PYTHON_VERSION": "3.11"},
	}
	resp, err := r.clientset.Sandbox().Create(createCtx, req)
	if err != nil {
		return fmt.Errorf("create sandbox: %w", err)
	}
	if resp == nil || resp.Id == "" {
		return fmt.Errorf("create sandbox: empty sandbox id")
	}
	sandboxID := resp.Id
	cleanupOnError := true
	defer func() {
		if cleanupOnError {
			_ = r.clientset.Sandbox().Delete(context.Background(), sandboxID)
		}
	}()

	if err := waitForRunning(createCtx, r.clientset, sandboxID); err != nil {
		return err
	}

	execdURL, accessToken, err := resolveExecdEndpoint(createCtx, r.clientset, r.cfg, sandboxID)
	if err != nil {
		return err
	}

	execdCfg := opensandbox.DefaultConfig()
	execdCfg.SandboxAPIURL = r.cfg.SandboxAPIURL
	execdCfg.APIKey = r.cfg.APIKey
	execdCfg.ExecdAPIURL = execdURL
	execdCfg.AccessToken = accessToken

	execdClient, err := opensandbox.NewForConfig(execdCfg)
	if err != nil {
		return fmt.Errorf("create execd client: %w", err)
	}
	if err := execdClient.Execd().Health().Ping(createCtx); err != nil {
		return fmt.Errorf("execd health check failed: %w", err)
	}

	lang := "python"
	pyContext, err := execdClient.Execd().Code().CreateContext(createCtx, execd.CodeContextRequest{Language: &lang})
	if err != nil {
		return fmt.Errorf("create python context: %w", err)
	}
	if pyContext == nil || pyContext.Id == nil || *pyContext.Id == "" {
		return fmt.Errorf("create python context: empty context id")
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	if r.execd != nil && r.pyContext != nil && r.sandboxID != "" {
		// Another request initialized the runtime first.
		_ = r.clientset.Sandbox().Delete(context.Background(), sandboxID)
		return nil
	}
	r.sandboxID = sandboxID
	r.execd = execdClient
	r.pyContext = pyContext
	cleanupOnError = false
	return nil
}

func (r *sandboxRuntime) executePython(ctx context.Context, code string, timeout time.Duration) (string, error) {
	if strings.TrimSpace(code) == "" {
		return "", fmt.Errorf("code cannot be empty")
	}
	if err := r.ensureReady(ctx); err != nil {
		return "", err
	}

	runCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	r.mu.Lock()
	execdClient := r.execd
	pyContext := r.pyContext
	r.mu.Unlock()

	resp, err := execdClient.Execd().Code().RunCode(runCtx, execd.RunCodeRequest{
		Context: pyContext,
		Code:    code,
	})
	if err != nil {
		return "", fmt.Errorf("run code: %w", err)
	}
	if resp == nil {
		return "", nil
	}
	return renderRunCodeOutput(resp), nil
}

func (r *sandboxRuntime) shutdown(ctx context.Context) error {
	r.mu.Lock()
	sandboxID := r.sandboxID
	r.sandboxID = ""
	r.execd = nil
	r.pyContext = nil
	r.mu.Unlock()

	if sandboxID == "" {
		return nil
	}
	return r.clientset.Sandbox().Delete(ctx, sandboxID)
}

func waitForRunning(ctx context.Context, cs *opensandbox.Clientset, sandboxID string) error {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	for {
		sb, err := cs.Sandbox().Get(ctx, sandboxID)
		if err != nil {
			return fmt.Errorf("get sandbox status: %w", err)
		}
		switch sb.Status.State {
		case "Running":
			return nil
		case "Failed":
			msg := ""
			if sb.Status.Message != nil {
				msg = *sb.Status.Message
			}
			return fmt.Errorf("sandbox failed: %s", msg)
		}

		select {
		case <-ctx.Done():
			return fmt.Errorf("wait for sandbox running: %w", ctx.Err())
		case <-ticker.C:
		}
	}
}

func resolveExecdEndpoint(ctx context.Context, cs *opensandbox.Clientset, cfg *opensandbox.Config, sandboxID string) (string, string, error) {
	ep, err := cs.Sandbox().GetEndpointWithProxy(ctx, sandboxID, 44772, false)
	if err != nil {
		return "", "", fmt.Errorf("get execd endpoint: %w", err)
	}
	if ep == nil || ep.Endpoint == "" {
		return "", "", fmt.Errorf("get execd endpoint: empty endpoint")
	}

	raw := strings.TrimSpace(ep.Endpoint)
	if idx := strings.Index(raw, "/proxy/"); idx >= 0 {
		raw = raw[:idx]
	}

	switch {
	case strings.HasPrefix(raw, "http://") || strings.HasPrefix(raw, "https://"):
	case strings.HasPrefix(raw, "/"):
		if cfg.SandboxAPIURL != "" {
			if idx := strings.Index(cfg.SandboxAPIURL, "://"); idx >= 0 {
				base := cfg.SandboxAPIURL[:idx+3]
				rest := cfg.SandboxAPIURL[idx+3:]
				if slash := strings.Index(rest, "/"); slash >= 0 {
					rest = rest[:slash]
				}
				raw = base + rest + raw
			}
		}
		if !strings.HasPrefix(raw, "http://") && !strings.HasPrefix(raw, "https://") {
			raw = "http://localhost:8080" + raw
		}
	default:
		raw = "http://" + raw
	}

	accessToken := os.Getenv(opensandbox.EnvAccessToken)
	if ep.Headers != nil {
		for k, v := range ep.Headers {
			if strings.EqualFold(k, "X-EXECD-ACCESS-TOKEN") {
				accessToken = v
				break
			}
		}
	}
	return raw, accessToken, nil
}

func renderRunCodeOutput(evt *execd.ServerStreamEvent) string {
	var out strings.Builder

	appendLine := func(s string) {
		if s == "" {
			return
		}
		if out.Len() > 0 && !strings.HasSuffix(out.String(), "\n") {
			out.WriteString("\n")
		}
		out.WriteString(s)
	}

	if text := evt.GetText(); text != "" {
		appendLine(text)
	}

	results := evt.GetResults()
	if plain, ok := results["text/plain"]; ok {
		appendLine(stringifyResult(plain))
	}

	if errObj, ok := evt.GetErrorOk(); ok && errObj != nil {
		name := ""
		if errObj.Ename != nil {
			name = *errObj.Ename
		}
		value := ""
		if errObj.Evalue != nil {
			value = *errObj.Evalue
		}
		switch {
		case name != "" && value != "":
			appendLine(name + ": " + value)
		case value != "":
			appendLine(value)
		}
		if len(errObj.Traceback) > 0 {
			appendLine(strings.Join(errObj.Traceback, "\n"))
		}
	}
	return strings.TrimSpace(out.String())
}

func stringifyResult(v any) string {
	switch t := v.(type) {
	case string:
		return t
	case []any:
		parts := make([]string, 0, len(t))
		for _, item := range t {
			parts = append(parts, fmt.Sprint(item))
		}
		return strings.Join(parts, "\n")
	default:
		return fmt.Sprint(t)
	}
}

func parseTimeout(seconds *int64) (time.Duration, error) {
	if seconds == nil {
		return defaultToolTimeout, nil
	}
	if *seconds <= 0 {
		return 0, fmt.Errorf("timeout_seconds must be >= 1")
	}
	d := time.Duration(*seconds) * time.Second
	if d > maxToolTimeout {
		return 0, fmt.Errorf("timeout_seconds must be <= %d", int(maxToolTimeout/time.Second))
	}
	return d, nil
}

func executePythonTool(runtime *sandboxRuntime) mcp.ToolHandlerFor[executePythonArgs, any] {
	return func(ctx context.Context, _ *mcp.CallToolRequest, args executePythonArgs) (*mcp.CallToolResult, any, error) {
		timeout, err := parseTimeout(args.TimeoutSeconds)
		if err != nil {
			return nil, nil, err
		}
		output, err := runtime.executePython(ctx, args.Code, timeout)
		if err != nil {
			return nil, nil, err
		}
		if output == "" {
			output = "(no output)"
		}
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: output},
			},
		}, nil, nil
	}
}

func main() {
	logger := log.New(os.Stderr, serverName+": ", log.LstdFlags)

	cfg := opensandbox.NewConfigFromEnv()
	runtime, err := newSandboxRuntime(cfg)
	if err != nil {
		logger.Fatalf("init runtime: %v", err)
	}

	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		if err := runtime.shutdown(ctx); err != nil {
			logger.Printf("cleanup sandbox: %v", err)
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(sigCh)
	go func() {
		<-sigCh
		_ = runtime.shutdown(context.Background())
		os.Exit(0)
	}()

	server := mcp.NewServer(&mcp.Implementation{
		Name:    serverName,
		Version: serverVersion,
	}, nil)

	mcp.AddTool(server, &mcp.Tool{
		Name:        toolExecutePython,
		Description: "Execute Python code in a single persistent OpenSandbox sandbox.",
	}, executePythonTool(runtime))

	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		logger.Fatalf("mcp server error: %v", err)
	}
}
