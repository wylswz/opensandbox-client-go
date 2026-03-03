# OpenSandbox Python MCP Example

This example exposes one MCP tool over stdio:

- `execute_python`: run Python code in a single reusable OpenSandbox sandbox

The server lazily creates one sandbox on first tool call and reuses it for subsequent calls.
When the process exits, it deletes that sandbox.

## Run

```bash
cd example/mcp-python
go run .
```

## Environment variables

- `OPEN_SANDBOX_SANDBOX_URL` (default `http://localhost:8080/v1`)
- `OPEN_SANDBOX_API_KEY` (optional)
- `OPEN_SANDBOX_EXECD_ACCESS_TOKEN` (fallback token if endpoint headers do not include one)
