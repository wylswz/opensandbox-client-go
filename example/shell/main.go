// Command opensandbox-shell is a CLI for listing and attaching to sandbox shells.
//
// Usage:
//
//	opensandbox-shell list
//	opensandbox-shell attach <sandbox-id>
//
// Environment variables:
//
//	OPEN_SANDBOX_SANDBOX_URL         Sandbox API URL (default: http://localhost:8080/v1)
//	OPEN_SANDBOX_API_KEY             API key
//	OPEN_SANDBOX_EXECD_ACCESS_TOKEN  Execd access token (fallback if not in endpoint headers)
package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"path"
	"strings"
	"syscall"
	"text/tabwriter"
	"time"

	"github.com/wylswz/opensandbox-client-go/internal/generated/execd"
	"github.com/wylswz/opensandbox-client-go/internal/generated/sandbox"
	"github.com/wylswz/opensandbox-client-go/internal/sse"
	"github.com/wylswz/opensandbox-client-go/pkg/opensandbox"
)

const (
	// defaultImage is used by the create command. debian:bookworm-slim has bash
	// and common utilities while staying small (~30 MB).
	defaultImage   = "debian:bookworm-slim"
	defaultTimeout = int32(3600) // 1 hour
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	cfg := opensandbox.NewConfigFromEnv()

	switch os.Args[1] {
	case "list":
		if err := runList(cfg); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	case "create":
		if err := runCreate(cfg); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	case "rm":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "error: rm requires a sandbox ID")
			usage()
			os.Exit(1)
		}
		if err := runRm(cfg, os.Args[2:]); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	case "attach":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "error: attach requires a sandbox ID")
			usage()
			os.Exit(1)
		}
		if err := runAttach(cfg, os.Args[2]); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %q\n", os.Args[1])
		usage()
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprint(os.Stderr, `Usage: opensandbox-shell <command> [args]

Commands:
  list              List all sandboxes
  create            Create a new sandbox (image: `+defaultImage+`)
  rm <id> [id...]   Delete one or more sandboxes
  attach <id>       Attach an interactive shell to a sandbox

Environment variables:
  OPEN_SANDBOX_SANDBOX_URL         Sandbox API URL (default: http://localhost:8080/v1)
  OPEN_SANDBOX_API_KEY             API key for the sandbox API
  OPEN_SANDBOX_EXECD_ACCESS_TOKEN  Access token for execd (fallback if not in endpoint)
`)
}

// runList prints a table of all sandboxes.
func runList(cfg *opensandbox.Config) error {
	cs, err := opensandbox.NewForConfig(cfg)
	if err != nil {
		return fmt.Errorf("client: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	resp, err := cs.Sandbox().List(ctx, &opensandbox.ListOptions{PageSize: 100})
	if err != nil {
		return fmt.Errorf("list: %w", err)
	}

	if len(resp.Items) == 0 {
		fmt.Println("No sandboxes found.")
		return nil
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tSTATE\tIMAGE\tCREATED\tEXPIRES")
	fmt.Fprintln(w, "--\t-----\t-----\t-------\t-------")
	for _, sb := range resp.Items {
		image := sb.Image.Uri
		created := sb.CreatedAt.Local().Format("2006-01-02 15:04:05")
		expires := sb.ExpiresAt.Local().Format("2006-01-02 15:04:05")
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n",
			sb.Id, sb.Status.State, image, created, expires)
	}
	return w.Flush()
}

// runCreate creates a new sandbox using the fixed bash image and prints its ID.
// It waits until the sandbox reaches Running state before returning.
func runCreate(cfg *opensandbox.Config) error {
	cs, err := opensandbox.NewForConfig(cfg)
	if err != nil {
		return fmt.Errorf("client: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	req := &sandbox.CreateSandboxRequest{
		Image:          sandbox.ImageSpec{Uri: defaultImage},
		Timeout:        defaultTimeout,
		ResourceLimits: map[string]string{"cpu": "500m", "memory": "512Mi"},
		Entrypoint:     []string{"/bin/bash", "-c", "sleep infinity"},
	}

	resp, err := cs.Sandbox().Create(ctx, req)
	if err != nil {
		return fmt.Errorf("create: %w", err)
	}

	fmt.Fprintf(os.Stderr, "Creating sandbox %s (image: %s)...\n", resp.Id, defaultImage)

	// Poll until Running or Failed.
	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("timed out waiting for sandbox %s to start", resp.Id)
		case <-time.After(2 * time.Second):
		}

		sb, err := cs.Sandbox().Get(ctx, resp.Id)
		if err != nil {
			return fmt.Errorf("get: %w", err)
		}
		switch sb.Status.State {
		case "Running":
			fmt.Println(resp.Id)
			return nil
		case "Failed":
			msg := ""
			if sb.Status.Message != nil {
				msg = *sb.Status.Message
			}
			return fmt.Errorf("sandbox %s failed: %s", resp.Id, msg)
		}
	}
}

// runRm deletes one or more sandboxes by ID.
func runRm(cfg *opensandbox.Config, ids []string) error {
	cs, err := opensandbox.NewForConfig(cfg)
	if err != nil {
		return fmt.Errorf("client: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var firstErr error
	for _, id := range ids {
		if err := cs.Sandbox().Delete(ctx, id); err != nil {
			fmt.Fprintf(os.Stderr, "error: rm %s: %v\n", id, err)
			if firstErr == nil {
				firstErr = err
			}
			continue
		}
		fmt.Println(id)
	}
	return firstErr
}

// runAttach attaches an interactive pseudo-shell REPL to a sandbox.
//
// Since the execd API is request/response (no persistent PTY), each command is
// sent as an independent shell invocation with the tracked working directory
// passed as Cwd. Output is streamed via SSE so you see it as it arrives.
func runAttach(cfg *opensandbox.Config, sandboxID string) error {
	cs, err := opensandbox.NewForConfig(cfg)
	if err != nil {
		return fmt.Errorf("client: %w", err)
	}

	ctx := context.Background()

	// Verify sandbox exists and is Running.
	sb, err := cs.Sandbox().Get(ctx, sandboxID)
	if err != nil {
		return fmt.Errorf("get sandbox %q: %w", sandboxID, err)
	}
	if sb.Status.State != "Running" {
		return fmt.Errorf("sandbox %q is not Running (state: %s)", sandboxID, sb.Status.State)
	}

	// Resolve execd URL and access token from the sandbox endpoint.
	execdURL, accessToken, err := resolveExecdEndpoint(ctx, cs, cfg, sandboxID)
	if err != nil {
		return fmt.Errorf("resolve execd endpoint: %w", err)
	}

	// Quick health check so we fail fast if execd is unreachable.
	execdCfg := opensandbox.DefaultConfig()
	execdCfg.SandboxAPIURL = cfg.SandboxAPIURL
	execdCfg.APIKey = cfg.APIKey
	execdCfg.ExecdAPIURL = execdURL
	execdCfg.AccessToken = accessToken
	execdCS, err := opensandbox.NewForConfig(execdCfg)
	if err != nil {
		return fmt.Errorf("execd client: %w", err)
	}
	pingCtx, pingCancel := context.WithTimeout(ctx, 10*time.Second)
	defer pingCancel()
	if err := execdCS.Execd().Health().Ping(pingCtx); err != nil {
		return fmt.Errorf("execd not reachable at %s: %w", execdURL, err)
	}

	fmt.Printf("Attached to sandbox %s (execd: %s)\n", sandboxID, execdURL)
	fmt.Println("Type 'exit' or press Ctrl-D to detach. Note: each command runs in a fresh shell; use 'cd' to navigate.")
	fmt.Println()

	// cancelCurrent cancels the in-flight command when Ctrl-C is pressed.
	// It is replaced atomically before each command and reset after.
	cancelCurrent := func() {}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT)
	defer signal.Stop(sigCh)
	go func() {
		for range sigCh {
			fmt.Println("^C")
			cancelCurrent()
		}
	}()

	cwd := "/"
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("[%s:%s]$ ", shortID(sandboxID), cwd)

		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println()
			break
		}
		if err != nil {
			return fmt.Errorf("stdin: %w", err)
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if line == "exit" || line == "quit" {
			break
		}

		cmdCtx, cancel := context.WithCancel(ctx)
		cancelCurrent = cancel

		// cd is resolved locally: parse the argument and apply path arithmetic.
		// Each command runs in a fresh shell, so we track cwd ourselves and pass
		// it via the Cwd field. No server round-trip needed.
		if isCdCommand(line) {
			cancel()
			cancelCurrent = func() {}
			cwd = resolveLocalCd(cwd, line)
			continue
		}

		// Stream the command output in real time.
		runErr := streamCommand(cmdCtx, execdURL, accessToken, line, cwd)
		canceled := cmdCtx.Err() == context.Canceled
		cancel()
		cancelCurrent = func() {}
		if runErr != nil && !canceled {
			fmt.Fprintf(os.Stderr, "error: %v\n", runErr)
		}
	}

	fmt.Printf("Detached from sandbox %s.\n", sandboxID)
	return nil
}

// resolveExecdEndpoint fetches the execd endpoint (port 44772) for a sandbox
// and returns the base URL and access token to use.
func resolveExecdEndpoint(ctx context.Context, cs *opensandbox.Clientset, cfg *opensandbox.Config, sandboxID string) (string, string, error) {
	ep, err := cs.Sandbox().GetEndpointWithProxy(ctx, sandboxID, 44772, false)
	if err != nil {
		return "", "", fmt.Errorf("GetEndpoint(44772): %w", err)
	}
	if ep == nil || ep.Endpoint == "" {
		return "", "", fmt.Errorf("empty endpoint returned")
	}

	raw := strings.TrimSpace(ep.Endpoint)

	// Strip /proxy/44772 suffix — execd serves at root, not under the proxy path.
	if idx := strings.Index(raw, "/proxy/"); idx >= 0 {
		raw = raw[:idx]
	}

	// Add scheme if missing.
	switch {
	case strings.HasPrefix(raw, "http://") || strings.HasPrefix(raw, "https://"):
		// already absolute
	case strings.HasPrefix(raw, "/"):
		// path-only — prepend scheme+host from the sandbox API URL
		if cfg.SandboxAPIURL != "" {
			if u, err := url.Parse(cfg.SandboxAPIURL); err == nil {
				raw = u.Scheme + "://" + u.Host + raw
				break
			}
		}
		raw = "http://localhost:8080" + raw
	default:
		raw = "http://" + raw
	}

	// Extract access token from endpoint headers, fall back to env var.
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

// streamCommand sends a shell command to execd and prints its output to stdout
// as SSE events arrive. A trailing newline is added if the output lacks one.
func streamCommand(ctx context.Context, execdURL, accessToken, cmd, cwd string) error {
	reqBody := map[string]interface{}{"command": cmd}
	if cwd != "" && cwd != "/" {
		reqBody["cwd"] = cwd
	}

	req, err := buildExecdRequest(ctx, execdURL, accessToken, "/command", reqBody)
	if err != nil {
		return err
	}

	debug := os.Getenv("OPENSANDBOX_DEBUG") != ""
	var lastChar byte
	err = sse.Stream(ctx, http.DefaultClient, req, func(raw json.RawMessage) {
		if debug {
			fmt.Fprintf(os.Stderr, "[SSE] %s\n", string(raw))
		}
		text := extractText(raw)
		if text != "" {
			fmt.Print(text)
			lastChar = text[len(text)-1]
		}
	})

	// Ensure the next prompt starts on a fresh line.
	if lastChar != 0 && lastChar != '\n' {
		fmt.Println()
	}
	return err
}

// extractText pulls printable text out of a raw SSE data payload.
// It accepts multiple server event shapes so attach works against older/newer
// server versions (text, stdout/stderr, output/content, results text/plain).
func extractText(raw json.RawMessage) string {
	// Strategy 1: decode to a generic map first and extract common keys.
	// This avoids losing output when JSON unmarshals successfully into
	// ServerStreamEvent but field names differ (e.g. "stdout", "output").
	var m map[string]interface{}
	if err := json.Unmarshal(raw, &m); err == nil {
		if s := firstString(m, "text", "stdout", "stderr", "output", "content", "message", "data"); s != "" {
			return ensureTrailingNewline(s)
		}
		if results, ok := m["results"].(map[string]interface{}); ok {
			if s := firstString(results, "text/plain", "stdout", "stderr", "output", "text", "content"); s != "" {
				return ensureTrailingNewline(s)
			}
		}
		// Some servers only send error blocks with no text payload.
		if errObj, ok := m["error"].(map[string]interface{}); ok {
			evalue := firstString(errObj, "evalue", "value", "message")
			if evalue == "" {
				evalue = "command failed"
			}
			if tbRaw, ok := errObj["traceback"].([]interface{}); ok && len(tbRaw) > 0 {
				tb := make([]string, 0, len(tbRaw))
				for _, item := range tbRaw {
					tb = append(tb, toString(item))
				}
				return ensureTrailingNewline(evalue + "\n" + strings.Join(tb, "\n"))
			}
			return ensureTrailingNewline(evalue)
		}
		return ""
	}

	// Strategy 2: decode with typed struct for expected modern schema.
	var evt execd.ServerStreamEvent
	if err := json.Unmarshal(raw, &evt); err == nil {
		if evt.Text != nil && *evt.Text != "" {
			if evt.Error != nil && evt.Error.Evalue != nil && *evt.Error.Evalue != "" {
				tb := strings.Join(evt.Error.Traceback, "\n")
				if tb != "" {
					return ensureTrailingNewline(*evt.Text + "\n" + tb)
				}
			}
			return ensureTrailingNewline(*evt.Text)
		}
		if plain, ok := evt.Results["text/plain"]; ok {
			if s, ok := plain.(string); ok && s != "" {
				return ensureTrailingNewline(s)
			}
		}
		return ""
	}

	// Strategy 3: server sent raw (non-JSON) text in the data field.
	s := strings.TrimRight(string(raw), "\r\n")
	if s != "" {
		return s + "\n"
	}
	return ""
}

func firstString(m map[string]interface{}, keys ...string) string {
	for _, k := range keys {
		if v, ok := m[k]; ok {
			if s := toString(v); s != "" {
				return s
			}
		}
	}
	return ""
}

func toString(v interface{}) string {
	switch t := v.(type) {
	case string:
		return t
	case []byte:
		return string(t)
	default:
		return ""
	}
}

func ensureTrailingNewline(s string) string {
	if s == "" || strings.HasSuffix(s, "\n") {
		return s
	}
	return s + "\n"
}

// buildExecdRequest constructs an HTTP POST request targeting an execd endpoint.
func buildExecdRequest(ctx context.Context, execdURL, accessToken, endpoint string, body interface{}) (*http.Request, error) {
	reqURL, err := url.JoinPath(strings.TrimSuffix(execdURL, "/"), endpoint)
	if err != nil {
		return nil, err
	}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/event-stream")
	if accessToken != "" {
		req.Header.Set("X-EXECD-ACCESS-TOKEN", accessToken)
	}
	return req, nil
}

// isCdCommand reports whether the input line is a cd command.
func isCdCommand(line string) bool {
	parts := strings.Fields(line)
	return len(parts) > 0 && parts[0] == "cd"
}

// resolveLocalCd applies a cd command to the current directory without a
// server round-trip. Handles absolute paths, relative paths, `cd` / `cd ~`
// (home), and `cd ..`. Does not support `cd -` (previous directory).
func resolveLocalCd(cwd, line string) string {
	parts := strings.Fields(line)
	if len(parts) < 2 || parts[1] == "~" {
		return "/root"
	}
	target := parts[1]
	if strings.HasPrefix(target, "~/") {
		target = "/root/" + target[2:]
	}
	if !strings.HasPrefix(target, "/") {
		target = path.Join(cwd, target)
	}
	return path.Clean(target)
}

// shortID returns the first 8 characters of a sandbox ID for display.
func shortID(id string) string {
	if len(id) > 8 {
		return id[:8]
	}
	return id
}

