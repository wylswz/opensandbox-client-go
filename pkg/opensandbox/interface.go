package opensandbox

import (
	"context"
	"io"

	"github.com/wylswz/opensandbox-client-go/internal/generated/execd"
	"github.com/wylswz/opensandbox-client-go/internal/generated/sandbox"
)

// Interface is the main client interface for OpenSandbox, similar to kubernetes.Interface.
type Interface interface {
	// Sandbox returns the sandbox lifecycle API client.
	Sandbox() SandboxInterface
	// Execd returns the code execution API client.
	Execd() ExecdInterface
}

// SandboxInterface provides access to sandbox lifecycle operations.
type SandboxInterface interface {
	// Create creates a new sandbox from a container image.
	Create(ctx context.Context, req *sandbox.CreateSandboxRequest) (*sandbox.CreateSandboxResponse, error)
	// List returns sandboxes with optional filters and pagination.
	List(ctx context.Context, opts *ListOptions) (*sandbox.ListSandboxesResponse, error)
	// Get retrieves a sandbox by ID.
	Get(ctx context.Context, sandboxID string) (*sandbox.Sandbox, error)
	// Delete removes a sandbox.
	Delete(ctx context.Context, sandboxID string) error
	// Pause pauses a running sandbox.
	Pause(ctx context.Context, sandboxID string) error
	// Resume resumes a paused sandbox.
	Resume(ctx context.Context, sandboxID string) error
	// RenewExpiration extends the sandbox TTL.
	RenewExpiration(ctx context.Context, sandboxID string, req *sandbox.RenewSandboxExpirationRequest) (*sandbox.RenewSandboxExpirationResponse, error)
	// GetEndpoint returns the access endpoint for a service port.
	GetEndpoint(ctx context.Context, sandboxID string, port int32) (*sandbox.Endpoint, error)
	// GetEndpointWithProxy returns the access endpoint, optionally via server proxy (for reachability from host).
	GetEndpointWithProxy(ctx context.Context, sandboxID string, port int32, useServerProxy bool) (*sandbox.Endpoint, error)
}

// SandboxListOptions contains options for listing sandboxes.
type ListOptions struct {
	// State filters by lifecycle state (OR logic for multiple).
	State []string
	// Metadata filters by key-value pairs (URL encoded).
	Metadata string
	// Page is the page number (1-based).
	Page int32
	// PageSize is the number of items per page.
	PageSize int32
}

// ExecdInterface provides access to code execution, commands, filesystem, and metrics.
type ExecdInterface interface {
	// Code returns the code execution interface.
	Code() CodeInterface
	// Command returns the command execution interface.
	Command() CommandInterface
	// Filesystem returns the file operations interface.
	Filesystem() FilesystemInterface
	// Metrics returns the system metrics interface.
	Metrics() MetricsInterface
	// Health returns the health check interface.
	Health() HealthInterface
}

// CodeInterface provides code execution operations.
type CodeInterface interface {
	// CreateContext creates a new code execution context.
	CreateContext(ctx context.Context, req execd.CodeContextRequest) (*execd.CodeContext, error)
	// ListContexts lists active contexts, optionally filtered by language.
	ListContexts(ctx context.Context, language string) ([]execd.CodeContext, error)
	// GetContext retrieves a context by ID.
	GetContext(ctx context.Context, contextID string) (*execd.CodeContext, error)
	// DeleteContext deletes a specific context.
	DeleteContext(ctx context.Context, contextID string) error
	// DeleteContextsByLanguage deletes all contexts for a language.
	DeleteContextsByLanguage(ctx context.Context, language string) error
	// RunCode executes code in a context.
	RunCode(ctx context.Context, req execd.RunCodeRequest) (*execd.ServerStreamEvent, error)
	// InterruptCode interrupts running code execution.
	InterruptCode(ctx context.Context) error
}

// CommandInterface provides shell command execution.
type CommandInterface interface {
	// Run executes a shell command.
	Run(ctx context.Context, req execd.RunCommandRequest) (*execd.ServerStreamEvent, error)
	// GetStatus returns the status of a command session.
	GetStatus(ctx context.Context, sessionID string) (*execd.CommandStatusResponse, error)
	// GetLogs returns stdout/stderr for a background command.
	GetLogs(ctx context.Context, sessionID string, cursor *int64) (string, error)
	// Interrupt interrupts command execution.
	Interrupt(ctx context.Context) error
}

// FilesystemInterface provides file and directory operations.
type FilesystemInterface interface {
	// GetInfo returns metadata for files.
	GetInfo(ctx context.Context, paths []string) (*map[string]execd.FileInfo, error)
	// Upload uploads a file (path is the destination path in sandbox).
	Upload(ctx context.Context, path string, file io.Reader) error
	// Download downloads a file, returns the response body to read.
	Download(ctx context.Context, path string) (io.ReadCloser, error)
	// Delete removes files.
	Delete(ctx context.Context, paths []string) error
	// CreateDirectory creates directories (mkdir -p semantics).
	CreateDirectory(ctx context.Context, path string, mode *int32) error
	// DeleteDirectory recursively deletes a directory.
	DeleteDirectory(ctx context.Context, path string) error
}

// MetricsInterface provides system resource metrics.
type MetricsInterface interface {
	// Get returns current CPU and memory metrics.
	Get(ctx context.Context) (*execd.Metrics, error)
	// Watch returns real-time metrics via SSE stream.
	Watch(ctx context.Context) (*execd.Metrics, error)
}

// HealthInterface provides health check operations.
type HealthInterface interface {
	// Ping performs a health check.
	Ping(ctx context.Context) error
}
