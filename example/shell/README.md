# OpenSandbox Shell Example

CLI for listing, creating, deleting, and attaching to OpenSandbox sandboxes.

## Run

```bash
cd example/shell
go run . list
```

## Environment variables

- `OPEN_SANDBOX_SANDBOX_URL` (default `http://localhost:8080/v1`)
- `OPEN_SANDBOX_API_KEY` (optional)
- `OPEN_SANDBOX_EXECD_ACCESS_TOKEN` (used when endpoint headers do not include token)
