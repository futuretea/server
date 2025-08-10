# server
Simple HTTP server for logging all GET and POST requests

## Features
- Logs all incoming GET and POST requests with path, method, headers, and body
- Configurable listening paths via environment variable
- Supports multiple paths with comma-separated configuration
- Default path: `/`

## Configuration

### Listening Paths
Configure which paths the server should listen on using the `LISTEN_PATHS` environment variable:

```bash
# Single path (default behavior)
LISTEN_PATHS=/

# Multiple paths
LISTEN_PATHS=/,/api,/health,/webhook

# No spaces around commas
LISTEN_PATHS=/api/v1,/api/v2,/health
```

If `LISTEN_PATHS` is not set or empty, the server defaults to listening on `/`.

## Usage

### Docker
```bash
# Run with default path (/)
docker run -d --name fake-server -p 10000:8080 futuretea/server

# Run with custom paths
docker run -d --name fake-server -p 10000:8080 \
  -e LISTEN_PATHS=/,/api,/health \
  futuretea/server
```

### Local Development
```bash
# Run with default path
go run main.go

# Run with custom paths
LISTEN_PATHS=/,/api,/health go run main.go
```

### Testing
```bash
# Test default path
curl http://127.0.0.1:10000

# Test custom paths (if configured)
curl http://127.0.0.1:10000/api
curl http://127.0.0.1:10000/health
```
