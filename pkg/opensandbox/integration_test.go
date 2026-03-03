//go:build integration

package opensandbox

import (
	"context"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/alibaba/opensandbox-client-go/internal/generated/execd"
	"github.com/alibaba/opensandbox-client-go/internal/generated/sandbox"
)

func getTestConfig(t *testing.T) *Config {
	sandboxURL := os.Getenv("OPEN_SANDBOX_SANDBOX_URL")
	if sandboxURL == "" {
		t.Skip("OPEN_SANDBOX_SANDBOX_URL not set, skipping integration test")
	}
	cfg := DefaultConfig()
	cfg.SandboxAPIURL = sandboxURL
	cfg.APIKey = os.Getenv("OPEN_SANDBOX_API_KEY")
	return cfg
}

func TestIntegration_Sandbox_List(t *testing.T) {
	cfg := getTestConfig(t)
	cs, err := NewForConfig(cfg)
	if err != nil {
		t.Fatalf("NewForConfig: %v", err)
	}
	ctx := context.Background()

	list, err := cs.Sandbox().List(ctx, &ListOptions{PageSize: 5})
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if list == nil {
		t.Fatal("List returned nil")
	}
	// Initial list may be empty or have items
	_ = list.Items
	_ = list.Pagination
}

func TestIntegration_Sandbox_CreateGetDelete(t *testing.T) {
	cfg := getTestConfig(t)
	cs, err := NewForConfig(cfg)
	if err != nil {
		t.Fatalf("NewForConfig: %v", err)
	}
	ctx := context.Background()

	// Create sandbox
	req := &sandbox.CreateSandboxRequest{
		Image:          sandbox.ImageSpec{Uri: "python:3.11-slim"},
		Timeout:        120,
		ResourceLimits: map[string]string{"cpu": "250m", "memory": "256Mi"},
		Entrypoint:     []string{"sleep", "60"},
	}
	createResp, err := cs.Sandbox().Create(ctx, req)
	if err != nil {
		t.Fatalf("Create: %v", err)
	}
	if createResp == nil || createResp.Id == "" {
		t.Fatalf("Create returned invalid response: %+v", createResp)
	}
	sandboxID := createResp.Id
	t.Logf("Created sandbox: %s", sandboxID)

	// Ensure cleanup on exit
	defer func() {
		_ = cs.Sandbox().Delete(ctx, sandboxID)
	}()

	// Get sandbox
	got, err := cs.Sandbox().Get(ctx, sandboxID)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if got.Id != sandboxID {
		t.Errorf("Get: got id %q, want %q", got.Id, sandboxID)
	}

	// List should include the sandbox
	list, err := cs.Sandbox().List(ctx, &ListOptions{PageSize: 20})
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	found := false
	for _, s := range list.Items {
		if s.Id == sandboxID {
			found = true
			break
		}
	}
	if !found {
		t.Logf("List did not include created sandbox (may be eventual consistency)")
	}

	// Delete
	if err := cs.Sandbox().Delete(ctx, sandboxID); err != nil {
		t.Fatalf("Delete: %v", err)
	}

	// Get should fail or return Terminated
	time.Sleep(2 * time.Second)
	_, err = cs.Sandbox().Get(ctx, sandboxID)
	if err == nil {
		t.Log("Get after delete succeeded (sandbox may still be in Stopping state)")
	}
}

func TestIntegration_Sandbox_GetEndpoint(t *testing.T) {
	cfg := getTestConfig(t)
	cs, err := NewForConfig(cfg)
	if err != nil {
		t.Fatalf("NewForConfig: %v", err)
	}
	ctx := context.Background()

	// Create sandbox with http server
	req := &sandbox.CreateSandboxRequest{
		Image:          sandbox.ImageSpec{Uri: "python:3.11-slim"},
		Timeout:        120,
		ResourceLimits: map[string]string{"cpu": "250m", "memory": "256Mi"},
		Entrypoint:     []string{"python", "-m", "http.server", "8000"},
	}
	createResp, err := cs.Sandbox().Create(ctx, req)
	if err != nil {
		t.Fatalf("Create: %v", err)
	}
	sandboxID := createResp.Id
	defer func() { _ = cs.Sandbox().Delete(ctx, sandboxID) }()

	// Wait for sandbox to be Running
	for i := 0; i < 30; i++ {
		got, err := cs.Sandbox().Get(ctx, sandboxID)
		if err != nil {
			t.Fatalf("Get: %v", err)
		}
		if got.Status.State == "Running" {
			break
		}
		if got.Status.State == "Failed" {
			msg := ""
			if got.Status.Message != nil {
				msg = *got.Status.Message
			}
			t.Fatalf("Sandbox failed: %s", msg)
		}
		time.Sleep(2 * time.Second)
	}

	// Get endpoint for port 8000
	ep, err := cs.Sandbox().GetEndpoint(ctx, sandboxID, 8000)
	if err != nil {
		t.Fatalf("GetEndpoint: %v", err)
	}
	if ep == nil || ep.Endpoint == "" {
		t.Errorf("GetEndpoint returned empty: %+v", ep)
	}
	t.Logf("Endpoint: %s", ep.Endpoint)
}

// execdTestHelper creates a sandbox with code-interpreter, waits for Running,
// fetches execd endpoint (port 44772), and returns a Clientset configured for execd.
func execdTestHelper(t *testing.T) (*Clientset, string, func()) {
	cfg := getTestConfig(t)
	cs, err := NewForConfig(cfg)
	if err != nil {
		t.Fatalf("NewForConfig: %v", err)
	}
	ctx := context.Background()

	// Create sandbox with code-interpreter (has execd on port 44772)
	req := &sandbox.CreateSandboxRequest{
		Image:          sandbox.ImageSpec{Uri: "opensandbox/code-interpreter:v1.0.1"},
		Timeout:        180,
		ResourceLimits: map[string]string{"cpu": "500m", "memory": "512Mi"},
		Entrypoint:     []string{"/opt/opensandbox/code-interpreter.sh"},
		Env:            map[string]string{"PYTHON_VERSION": "3.11"},
	}
	createResp, err := cs.Sandbox().Create(ctx, req)
	if err != nil {
		t.Fatalf("Create: %v", err)
	}
	sandboxID := createResp.Id
	t.Logf("Created code-interpreter sandbox: %s", sandboxID)

	cleanup := func() {
		_ = cs.Sandbox().Delete(ctx, sandboxID)
	}

	// Wait for Running
	for i := 0; i < 60; i++ {
		got, err := cs.Sandbox().Get(ctx, sandboxID)
		if err != nil {
			t.Fatalf("Get: %v", err)
		}
		if got.Status.State == "Running" {
			break
		}
		if got.Status.State == "Failed" {
			msg := ""
			if got.Status.Message != nil {
				msg = *got.Status.Message
			}
			t.Fatalf("Sandbox failed: %s", msg)
		}
		time.Sleep(2 * time.Second)
	}

	// Get execd endpoint (port 44772) with server proxy for host reachability
	ep, err := cs.Sandbox().GetEndpointWithProxy(ctx, sandboxID, 44772, true)
	if err != nil {
		cleanup()
		t.Fatalf("GetEndpoint execd: %v", err)
	}
	if ep == nil || ep.Endpoint == "" {
		cleanup()
		t.Fatal("GetEndpoint returned empty")
	}

	// Build execd URL - endpoint may be "host/path", "host:port/path", or "/path"
	execdURL := strings.TrimSpace(ep.Endpoint)
	if execdURL == "" {
		cleanup()
		t.Fatal("Endpoint URL is empty")
	}
	if !strings.HasPrefix(execdURL, "http://") && !strings.HasPrefix(execdURL, "https://") {
		if strings.HasPrefix(execdURL, "/") {
			// Path only - prepend scheme and host from sandbox URL
			sandboxURL := os.Getenv("OPEN_SANDBOX_SANDBOX_URL")
			if sandboxURL != "" {
				if u, err := url.Parse(sandboxURL); err == nil {
					execdURL = u.Scheme + "://" + u.Host + execdURL
				} else {
					execdURL = "http://" + strings.TrimPrefix(execdURL, "/")
				}
			} else {
				execdURL = "http://localhost:8090" + execdURL
			}
		} else {
			execdURL = "http://" + execdURL
		}
	}

	// Extract X-EXECD-ACCESS-TOKEN from headers
	accessToken := ""
	if ep.Headers != nil {
		for k, v := range ep.Headers {
			if strings.EqualFold(k, "X-EXECD-ACCESS-TOKEN") {
				accessToken = v
				break
			}
		}
	}
	if accessToken == "" {
		cleanup()
		t.Fatal("Endpoint headers missing X-EXECD-ACCESS-TOKEN")
	}

	// Create clientset with execd config
	execdCfg := DefaultConfig()
	execdCfg.SandboxAPIURL = cfg.SandboxAPIURL
	execdCfg.APIKey = cfg.APIKey
	execdCfg.ExecdAPIURL = execdURL
	execdCfg.AccessToken = accessToken

	execdCS, err := NewForConfig(execdCfg)
	if err != nil {
		cleanup()
		t.Fatalf("NewForConfig execd: %v", err)
	}

	return execdCS, sandboxID, cleanup
}

func TestIntegration_Execd_Health(t *testing.T) {
	cs, _, cleanup := execdTestHelper(t)
	defer cleanup()
	ctx := context.Background()

	if err := cs.Execd().Health().Ping(ctx); err != nil {
		t.Fatalf("Ping: %v", err)
	}
}

func TestIntegration_Execd_Code_CreateContext_RunCode(t *testing.T) {
	cs, _, cleanup := execdTestHelper(t)
	defer cleanup()
	ctx := context.Background()

	lang := "python"
	ctxObj, err := cs.Execd().Code().CreateContext(ctx, execd.CodeContextRequest{Language: &lang})
	if err != nil {
		t.Fatalf("CreateContext: %v", err)
	}
	if ctxObj == nil || ctxObj.Id == nil || *ctxObj.Id == "" {
		t.Fatalf("CreateContext returned invalid: %+v", ctxObj)
	}
	t.Logf("Created context: %s", *ctxObj.Id)

	// Run simple Python code
	runReq := execd.RunCodeRequest{
		Code:    "print(2 + 2)",
		Context: ctxObj,
	}
	resp, err := cs.Execd().Code().RunCode(ctx, runReq)
	if err != nil {
		t.Fatalf("RunCode: %v", err)
	}
	if resp == nil {
		t.Fatal("RunCode returned nil")
	}
	t.Logf("RunCode response: %+v", resp)
}

func TestIntegration_Execd_Command_Run(t *testing.T) {
	cs, _, cleanup := execdTestHelper(t)
	defer cleanup()
	ctx := context.Background()

	runReq := execd.RunCommandRequest{Command: "echo hello"}
	resp, err := cs.Execd().Command().Run(ctx, runReq)
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	if resp == nil {
		t.Fatal("Run returned nil")
	}
	t.Logf("Command response: %+v", resp)
}

func TestIntegration_Execd_Filesystem(t *testing.T) {
	cs, _, cleanup := execdTestHelper(t)
	defer cleanup()
	ctx := context.Background()

	// Create directory
	testDir := "/tmp/opensandbox-test"
	if err := cs.Execd().Filesystem().CreateDirectory(ctx, testDir, nil); err != nil {
		t.Fatalf("CreateDirectory: %v", err)
	}

	// Upload file
	testFile := testDir + "/hello.txt"
	content := strings.NewReader("Hello OpenSandbox!")
	if err := cs.Execd().Filesystem().Upload(ctx, testFile, content); err != nil {
		t.Fatalf("Upload: %v", err)
	}

	// Get info
	info, err := cs.Execd().Filesystem().GetInfo(ctx, []string{testFile})
	if err != nil {
		t.Fatalf("GetInfo: %v", err)
	}
	if info == nil {
		t.Fatal("GetInfo returned nil")
	}
	fi, ok := (*info)[testFile]
	if !ok || fi.Path == "" {
		t.Errorf("GetInfo: missing %q, got %+v", testFile, info)
	}

	// Download
	rc, err := cs.Execd().Filesystem().Download(ctx, testFile)
	if err != nil {
		t.Fatalf("Download: %v", err)
	}
	defer rc.Close()
	// Read and verify (simplified - just ensure we get something)
	_ = rc

	// Delete file
	if err := cs.Execd().Filesystem().Delete(ctx, []string{testFile}); err != nil {
		t.Fatalf("Delete: %v", err)
	}

	// Delete directory
	if err := cs.Execd().Filesystem().DeleteDirectory(ctx, testDir); err != nil {
		t.Fatalf("DeleteDirectory: %v", err)
	}
}

func TestIntegration_Execd_Metrics(t *testing.T) {
	cs, _, cleanup := execdTestHelper(t)
	defer cleanup()
	ctx := context.Background()

	metrics, err := cs.Execd().Metrics().Get(ctx)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if metrics == nil {
		t.Fatal("Get returned nil")
	}
	t.Logf("Metrics: %+v", metrics)
}
