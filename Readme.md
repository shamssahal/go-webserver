# Go HTTP Server

A production-ready HTTP server boilerplate built with Go, featuring structured logging, request tracing, health checks, and containerization.

## Features

- ✅ Structured JSON logging with `log/slog`
- ✅ Request tracing with `X-Request-ID`
- ✅ Health and readiness endpoints
- ✅ Graceful shutdown
- ✅ CORS support
- ✅ Panic recovery with stack traces
- ✅ Request timeout middleware
- ✅ Environment-based configuration
- ✅ Docker support with multi-stage builds

## Quick Start

### Prerequisites

- Go 1.22+
- Docker & Docker Compose (optional)

### Run Locally

```bash
# Install dependencies
go mod download

# Run the server
make run

# Server starts at http://localhost:3000
```

### Run with Docker

```bash
# Build and run (always fresh build)
make docker-run

# View logs
docker compose logs -f

# Stop
make docker-stop
```

## API Endpoints

| Endpoint  | Method | Description                            |
| --------- | ------ | -------------------------------------- |
| `/health` | GET    | Liveness probe (no middleware)         |
| `/ready`  | GET    | Readiness probe with dependency checks |
| `/do`     | GET    | Example application endpoint           |

## Configuration

Copy `.env.example` to `.env` and configure:

```bash
# Server
SERVER_HOST=localhost
SERVER_PORT=3000

# Timeouts (seconds)
SERVER_READ_TIMEOUT=10
SERVER_WRITE_TIMEOUT=15
SERVER_IDLE_TIMEOUT=60

# Environment
APP_ENV=development  # development or production
```

## Request Tracing

### How It Works

Every request gets a unique correlation ID for end-to-end tracing:

1. **Client sends optional ID:**

   ```bash
   curl -H "X-Request-ID: my-custom-id" http://localhost:3000/do
   ```

2. **Middleware (`RequestID`) extracts or generates ID:**

   - If `X-Request-ID` header present → use it
   - If missing → generate new UUID
   - Adds ID to request context
   - Returns ID in response header

3. **All logs include the request ID:**

   ```json
   {
     "time": "2025-12-10T12:00:00Z",
     "level": "INFO",
     "msg": "request completed",
     "request_id": "my-custom-id",
     "method": "GET",
     "path": "/do",
     "status": 200
   }
   ```

4. **Downstream services can propagate:**
   ```go
   rid := utils.RequestIDFromContext(r.Context())
   // Pass to external API, database queries, etc.
   ```

### Implementation

**Middleware** (`internal/transport/http/middleware/requestIDmw.go`):

```go
func RequestID(next http.Handler) http.Handler {
    rid := r.Header.Get(config.HeaderRequestID)
    if rid == "" {
        rid = utils.NewRequestID()  // Generate UUID
    }
    ctx := utils.WithRequestID(r.Context(), rid)
    w.Header().Set(config.HeaderRequestID, rid)
    next.ServeHTTP(w, r.WithContext(ctx))
}
```

**Extract in handlers:**

```go
rid := utils.RequestIDFromContext(r.Context())
slog.Info("processing request", "request_id", rid)
```

### Benefits

- **Distributed tracing** - Track requests across services
- **Debugging** - Correlate logs for a single request
- **Client correlation** - Clients can trace their requests
- **Monitoring** - Group metrics by request ID

## Architecture

```
cmd/api/              - Application entry point
config/               - Configuration management
internal/
  transport/http/
    handlers/         - HTTP handlers
    middleware/       - Middleware (CORS, logging, recovery, etc.)
pkg/
  errors/            - Error handling
  logger/            - Logger factory
  utils/             - Utilities (RequestID, WriteJson)
```

## Middleware Stack

Requests flow through middleware in this order:

```
Health/Ready → No middleware (fast response)

App routes → CORS → RequestTimeout → RequestID → Recover → RequestLog → Handler
```

## Development

```bash
# Build
make build

# Run
make run

# Clean
make clean

# Help
make help
```

## Docker Commands

```bash
# Build image
make docker-build

# Run container
make docker-run

# Stop container
make docker-stop

# Remove image
make docker-clean
```

## Production Deployment

1. Set `APP_ENV=production` in environment
2. Configure appropriate timeouts
3. Deploy with Docker:
   ```bash
   docker run -d \
     -p 3000:3000 \
     -e APP_ENV=production \
     -e SERVER_HOST=0.0.0.0 \
     go-server:latest
   ```

## Health Checks

- **Liveness** (`/health`): Always returns 200 if process is running
- **Readiness** (`/ready`): Returns 200 if ready to serve traffic (checks dependencies)

Load balancers should use `/ready` for routing decisions.
