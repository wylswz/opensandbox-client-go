package opensandbox

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/wylswz/opensandbox-client-go/internal/generated/execd"
	"github.com/wylswz/opensandbox-client-go/internal/sse"
)

type execdClient struct {
	client      *execd.APIClient
	accessToken string
	execdAPIURL string
	httpClient  *http.Client
	userAgent   string
}

func (c *execdClient) authContext(ctx context.Context) context.Context {
	if c.accessToken == "" {
		return ctx
	}
	return context.WithValue(ctx, execd.ContextAPIKeys, map[string]execd.APIKey{
		"AccessToken": {Key: c.accessToken},
	})
}

func (c *execdClient) Code() CodeInterface {
	return &codeClient{execd: c}
}

func (c *execdClient) Command() CommandInterface {
	return &commandClient{execd: c}
}

func (c *execdClient) Filesystem() FilesystemInterface {
	return &filesystemClient{execd: c}
}

func (c *execdClient) Metrics() MetricsInterface {
	return &metricsClient{execd: c}
}

func (c *execdClient) Health() HealthInterface {
	return &healthClient{execd: c}
}

// buildSSERequest builds an HTTP request for SSE endpoints (RunCode, RunCommand).
// All headers (auth, content-type, etc.) are set here; the sse package is pure.
func (c *execdClient) buildSSERequest(ctx context.Context, path string, body interface{}) (*http.Request, error) {
	baseURL := c.execdAPIURL
	if baseURL == "" && len(c.client.GetConfig().Servers) > 0 {
		baseURL = c.client.GetConfig().Servers[0].URL
	}
	baseURL = strings.TrimSuffix(baseURL, "/")

	reqURL, err := url.JoinPath(baseURL, path)
	if err != nil {
		return nil, err
	}

	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set(headerContentType, contentTypeJSON)
	req.Header.Set(headerAccept, acceptSSE)
	if c.userAgent != "" {
		req.Header.Set(headerUserAgent, c.userAgent)
	}
	if c.accessToken != "" {
		req.Header.Set(headerExecdAccessToken, c.accessToken)
	}

	return req, nil
}

func parseLastServerStreamEvent(events []json.RawMessage) (*execd.ServerStreamEvent, error) {
	if len(events) == 0 {
		return &execd.ServerStreamEvent{}, nil
	}
	var evt execd.ServerStreamEvent
	if err := json.Unmarshal(events[len(events)-1], &evt); err != nil {
		return nil, err
	}
	return &evt, nil
}

// parseAggregatedServerStreamEvent merges output from all SSE events into a single
// ServerStreamEvent. The execd API streams multiple events (stdout, result,
// execution_complete); the last event alone often has no output. Aggregating
// ensures print() output and result display_data are captured.
func parseAggregatedServerStreamEvent(events []json.RawMessage) (*execd.ServerStreamEvent, error) {
	merged := &execd.ServerStreamEvent{}
	var textParts []string
	var lastResults map[string]interface{}
	var lastError *execd.ServerStreamEventError
	var lastExecutionCount *int32
	var lastExecutionTime *int64

	for _, raw := range events {
		var evt execd.ServerStreamEvent
		if err := json.Unmarshal(raw, &evt); err != nil {
			continue
		}
		t := evt.GetType()
		// Capture text from stdout/stderr/stream; skip init/status/ping
		if (t == "stdout" || t == "stderr" || t == "stream" || t == "") && evt.GetText() != "" {
			textParts = append(textParts, evt.GetText())
		}
		// result or execute_result (Jupyter) carry display_data
		if (t == "result" || t == "execute_result") && evt.Results != nil {
			lastResults = evt.Results
		}
		if evt.Error != nil && lastError == nil {
			lastError = evt.Error
		}
		if evt.ExecutionCount != nil {
			lastExecutionCount = evt.ExecutionCount
		}
		if evt.ExecutionTime != nil {
			lastExecutionTime = evt.ExecutionTime
		}
	}

	if len(textParts) > 0 {
		merged.Text = stringPtr(strings.Join(textParts, ""))
	}
	if lastResults != nil {
		merged.Results = lastResults
	}
	if lastError != nil {
		merged.Error = lastError
	}
	if lastExecutionCount != nil {
		merged.ExecutionCount = lastExecutionCount
	}
	if lastExecutionTime != nil {
		merged.ExecutionTime = lastExecutionTime
	}
	return merged, nil
}

func stringPtr(s string) *string { return &s }

// codeClient implements CodeInterface
type codeClient struct {
	execd *execdClient
}

func (c *codeClient) CreateContext(ctx context.Context, req execd.CodeContextRequest) (*execd.CodeContext, error) {
	resp, _, err := c.execd.client.CodeInterpretingAPI.CreateCodeContext(c.execd.authContext(ctx)).
		CodeContextRequest(req).
		Execute()
	return resp, err
}

func (c *codeClient) ListContexts(ctx context.Context, language string) ([]execd.CodeContext, error) {
	resp, _, err := c.execd.client.CodeInterpretingAPI.ListContexts(c.execd.authContext(ctx)).
		Language(language).
		Execute()
	return resp, err
}

func (c *codeClient) GetContext(ctx context.Context, contextID string) (*execd.CodeContext, error) {
	resp, _, err := c.execd.client.CodeInterpretingAPI.GetContext(c.execd.authContext(ctx), contextID).Execute()
	return resp, err
}

func (c *codeClient) DeleteContext(ctx context.Context, contextID string) error {
	_, err := c.execd.client.CodeInterpretingAPI.DeleteContext(c.execd.authContext(ctx), contextID).Execute()
	return err
}

func (c *codeClient) DeleteContextsByLanguage(ctx context.Context, language string) error {
	_, err := c.execd.client.CodeInterpretingAPI.DeleteContextsByLanguage(c.execd.authContext(ctx)).
		Language(language).
		Execute()
	return err
}

func (c *codeClient) RunCode(ctx context.Context, req execd.RunCodeRequest) (*execd.ServerStreamEvent, error) {
	httpReq, err := c.execd.buildSSERequest(ctx, execdPathCode, req)
	if err != nil {
		return nil, err
	}
	events, err := sse.Do(ctx, c.execd.httpClient, httpReq)
	if err != nil {
		return nil, err
	}
	if os.Getenv("OPENSANDBOX_DEBUG") != "" {
		for i, raw := range events {
			log.Printf("[execd] event %d: %s", i, string(raw))
		}
	}
	merged, err := parseAggregatedServerStreamEvent(events)
	if err == nil && os.Getenv("OPENSANDBOX_DEBUG") != "" {
		log.Printf("[execd] merged: text=%q results=%v", merged.GetText(), merged.GetResults())
	}
	return merged, err
}

func (c *codeClient) InterruptCode(ctx context.Context) error {
	_, err := c.execd.client.CodeInterpretingAPI.InterruptCode(c.execd.authContext(ctx)).Execute()
	return err
}

// commandClient implements CommandInterface
type commandClient struct {
	execd *execdClient
}

func (c *commandClient) Run(ctx context.Context, req execd.RunCommandRequest) (*execd.ServerStreamEvent, error) {
	httpReq, err := c.execd.buildSSERequest(ctx, execdPathCommand, req)
	if err != nil {
		return nil, err
	}
	events, err := sse.Do(ctx, c.execd.httpClient, httpReq)
	if err != nil {
		return nil, err
	}
	return parseLastServerStreamEvent(events)
}

func (c *commandClient) Stream(ctx context.Context, req execd.RunCommandRequest, onEvent func(CommandStreamEvent) error) error {
	if onEvent == nil {
		return fmt.Errorf(errOnEventCallbackRequired)
	}
	streamCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	httpReq, err := c.execd.buildSSERequest(streamCtx, execdPathCommand, req)
	if err != nil {
		return err
	}

	var callbackErr error
	err = sse.Stream(streamCtx, c.execd.httpClient, httpReq, func(raw json.RawMessage) {
		if callbackErr != nil {
			return
		}
		evt, parseErr := parseCommandStreamEvent(raw)
		if parseErr != nil {
			callbackErr = parseErr
			cancel()
			return
		}
		if cbErr := onEvent(evt); cbErr != nil {
			callbackErr = cbErr
			cancel()
		}
	})
	if callbackErr != nil {
		return callbackErr
	}
	return err
}

func parseCommandStreamEvent(raw json.RawMessage) (CommandStreamEvent, error) {
	var evt execd.ServerStreamEvent
	if err := json.Unmarshal(raw, &evt); err != nil {
		return CommandStreamEvent{}, err
	}

	out := CommandStreamEvent{
		Type:           CommandStreamEventType(evt.GetType()),
		ExecutionCount: evt.ExecutionCount,
		ExecutionTime:  evt.ExecutionTime,
		Timestamp:      evt.Timestamp,
		Results:        evt.Results,
		Raw:            append(json.RawMessage(nil), raw...),
	}
	if evt.Text != nil {
		out.Text = *evt.Text
	}
	if evt.Error != nil {
		streamErr := &CommandStreamError{
			Traceback: evt.Error.Traceback,
		}
		if evt.Error.Ename != nil {
			streamErr.Name = *evt.Error.Ename
		}
		if evt.Error.Evalue != nil {
			streamErr.Value = *evt.Error.Evalue
		}
		out.Error = streamErr
	}
	return out, nil
}

func (c *commandClient) GetStatus(ctx context.Context, sessionID string) (*execd.CommandStatusResponse, error) {
	resp, _, err := c.execd.client.CommandAPI.GetCommandStatus(c.execd.authContext(ctx), sessionID).Execute()
	return resp, err
}

func (c *commandClient) GetLogs(ctx context.Context, sessionID string, cursor *int64) (string, error) {
	r := c.execd.client.CommandAPI.GetBackgroundCommandLogs(c.execd.authContext(ctx), sessionID)
	if cursor != nil {
		r = r.Cursor(*cursor)
	}
	resp, _, err := r.Execute()
	return resp, err
}

func (c *commandClient) Interrupt(ctx context.Context) error {
	_, err := c.execd.client.CommandAPI.InterruptCommand(c.execd.authContext(ctx)).Execute()
	return err
}

// filesystemClient implements FilesystemInterface
type filesystemClient struct {
	execd *execdClient
}

func (c *filesystemClient) GetInfo(ctx context.Context, paths []string) (*map[string]execd.FileInfo, error) {
	resp, _, err := c.execd.client.FilesystemAPI.GetFilesInfo(c.execd.authContext(ctx)).
		Path(paths).
		Execute()
	return resp, err
}

func (c *filesystemClient) Upload(ctx context.Context, path string, content io.Reader) error {
	// Build multipart manually: execd expects metadata as a file part (INVALID_FILE_METADATA / "metadata file is missing")
	// when sent as a form field. Send metadata first (JSON), then file.
	baseURL := c.execd.execdAPIURL
	if baseURL == "" && len(c.execd.client.GetConfig().Servers) > 0 {
		baseURL = c.execd.client.GetConfig().Servers[0].URL
	}
	baseURL = strings.TrimSuffix(baseURL, "/")
	reqURL, err := url.JoinPath(baseURL, execdPathFilesUpload)
	if err != nil {
		return err
	}

	body := &bytes.Buffer{}
	w := multipart.NewWriter(body)

	// 1. metadata part - execd expects it as a file part with Content-Type application/json
	metadata := execd.FileMetadata{Path: &path}
	metadataJSON, _ := json.Marshal(metadata)
	metadataPart, err := w.CreatePart(map[string][]string{
		"Content-Disposition": {`form-data; name="metadata"; filename="metadata.json"`},
		"Content-Type":        {contentTypeJSON},
	})
	if err != nil {
		return err
	}
	if _, err := metadataPart.Write(metadataJSON); err != nil {
		return err
	}

	// 2. file part
	fileBytes, err := io.ReadAll(content)
	if err != nil {
		return err
	}
	filePart, err := w.CreateFormFile("file", "upload")
	if err != nil {
		return err
	}
	if _, err := filePart.Write(fileBytes); err != nil {
		return err
	}

	contentType := w.FormDataContentType()
	if err := w.Close(); err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, body)
	if err != nil {
		return err
	}
	req.Header.Set(headerContentType, contentType)
	if c.execd.userAgent != "" {
		req.Header.Set(headerUserAgent, c.execd.userAgent)
	}
	if c.execd.accessToken != "" {
		req.Header.Set(headerExecdAccessToken, c.execd.accessToken)
	}

	hc := c.execd.httpClient
	if hc == nil {
		hc = c.execd.client.GetConfig().HTTPClient
	}
	resp, err := hc.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("%s (response: %s)", resp.Status, string(respBody))
	}
	return nil
}

func (c *filesystemClient) Download(ctx context.Context, path string) (io.ReadCloser, error) {
	file, _, err := c.execd.client.FilesystemAPI.DownloadFile(c.execd.authContext(ctx)).
		Path(path).
		Execute()
	if err != nil {
		return nil, err
	}
	if _, err := file.Seek(0, 0); err != nil {
		file.Close()
		return nil, err
	}
	// Generated client returns *os.File (temp file) which implements io.ReadCloser
	return file, nil
}

func (c *filesystemClient) Delete(ctx context.Context, paths []string) error {
	_, err := c.execd.client.FilesystemAPI.RemoveFiles(c.execd.authContext(ctx)).
		Path(paths).
		Execute()
	return err
}

func (c *filesystemClient) CreateDirectory(ctx context.Context, path string, mode *int32) error {
	modeVal := int32(755) // API expects decimal 755 (octal notation), not 0755 (493)
	if mode != nil {
		modeVal = *mode
	}
	perm := execd.Permission{Mode: modeVal}
	owner, group := defaultFileOwner, defaultFileGroup // Match OpenAPI example; some servers require these
	perm.Owner = &owner
	perm.Group = &group
	_, err := c.execd.client.FilesystemAPI.MakeDirs(c.execd.authContext(ctx)).
		RequestBody(map[string]execd.Permission{path: perm}).
		Execute()
	return err
}

func (c *filesystemClient) DeleteDirectory(ctx context.Context, path string) error {
	_, err := c.execd.client.FilesystemAPI.RemoveDirs(c.execd.authContext(ctx)).
		Path([]string{path}).
		Execute()
	return err
}

// metricsClient implements MetricsInterface
type metricsClient struct {
	execd *execdClient
}

func (c *metricsClient) Get(ctx context.Context) (*execd.Metrics, error) {
	resp, _, err := c.execd.client.MetricAPI.GetMetrics(c.execd.authContext(ctx)).Execute()
	return resp, err
}

func (c *metricsClient) Watch(ctx context.Context) (*execd.Metrics, error) {
	resp, _, err := c.execd.client.MetricAPI.WatchMetrics(c.execd.authContext(ctx)).Execute()
	return resp, err
}

// healthClient implements HealthInterface
type healthClient struct {
	execd *execdClient
}

func (c *healthClient) Ping(ctx context.Context) error {
	_, err := c.execd.client.HealthAPI.Ping(c.execd.authContext(ctx)).Execute()
	return err
}
