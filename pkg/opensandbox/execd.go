package opensandbox

import (
	"context"
	"encoding/json"
	"io"
	"os"

	"github.com/wylswz/opensandbox-client-go/internal/generated/execd"
)

type execdClient struct {
	client      *execd.APIClient
	accessToken string
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
	resp, _, err := c.execd.client.CodeInterpretingAPI.RunCode(c.execd.authContext(ctx)).
		RunCodeRequest(req).
		Execute()
	return resp, err
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
	resp, _, err := c.execd.client.CommandAPI.RunCommand(c.execd.authContext(ctx)).
		RunCommandRequest(req).
		Execute()
	return resp, err
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
	// Create temp file for multipart upload (generated API expects *os.File)
	tmp, err := os.CreateTemp("", "opensandbox-upload-*")
	if err != nil {
		return err
	}
	defer os.Remove(tmp.Name())
	defer tmp.Close()

	if _, err := io.Copy(tmp, content); err != nil {
		return err
	}
	if _, err := tmp.Seek(0, 0); err != nil {
		return err
	}

	metadata, _ := json.Marshal(execd.FileMetadata{Path: &path})
	_, err = c.execd.client.FilesystemAPI.UploadFile(c.execd.authContext(ctx)).
		Metadata(string(metadata)).
		File(tmp).
		Execute()
	return err
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
	modeVal := int32(0755)
	if mode != nil {
		modeVal = *mode
	}
	perm := execd.Permission{Mode: modeVal}
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
