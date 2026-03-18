# SkyControl Gateway API

API Gateway for the SkyControl platform. Proxies HTTP/gRPC requests to internal microservices (Auth, Telemetry, Platform) and exposes health/readiness probes for the entire system.

Built on [Kratos v2](https://github.com/go-kratos/kratos) with dual HTTP (port 8000) and gRPC (port 9000) support.

---

## Architecture

```
cmd/SkyControl_gateway_API/
├── main.go                  # Entry point
├── wire.go                  # Dependency injection declarations
└── wire_gen.go              # Auto-generated DI wiring

internal/
├── server/
│   ├── http.go              # HTTP server setup
│   └── grpc.go              # gRPC server setup
├── service/
│   ├── health_service.go    # Viability (health/ready) handlers
│   └── auth_service.go      # Auth proxy service
├── data/
│   └── data.go              # External gRPC client connections
├── biz/                     # Business logic layer (reserved)
└── conf/
    └── conf.proto           # Configuration schema

api/
└── skycontrol/
    ├── viability/            # Health check proto + generated code
    ├── common/               # Shared error reason enums
    └── generated/            # Auto-generated from remote contracts
        ├── auth/v1/
        ├── telemetry/v1/
        └── platform/v1/
```

---

## API Endpoints

### Viability (Health Checks)

| Method | Path | Description |
|--------|------|-------------|
| GET | `/api/v1/viability/health` | Gateway liveness — returns status and uptime |
| GET | `/api/v1/viability/ready` | System readiness — checks Auth, Telemetry, Platform |

**Health response:**
```json
{
  "gateway_status": "...",
  "gateway_uptime": "..."
}
```

**Ready response:**
```json
{
  "status": "...",
  "auth_status": "...",
  "auth_db_status": "...",
  "auth_uptime": "...",
  "telemetry_status": "...",
  "telemetry_db_status": "...",
  "telemetry_uptime": "..."
}
```

### Auth Service

HTTP and gRPC routes are proxied to the Auth microservice. Routes are generated from the remote `SkyControlAPI_Contracts` repository protobuf definitions.

---

## Configuration

Copy the example config and edit for your environment:

```bash
cp configs_example/cfg_ex.yaml.example configs/config.yaml
```

```yaml
server:
  http:
    addr: 0.0.0.0:8000
    timeout: 1s
  grpc:
    addr: 0.0.0.0:9000
    timeout: 1s
data:
  auth:
    addr: 127.0.0.1:8080   # Address of the Auth microservice
    timeout: 0.2s
```

> `configs/config.yaml` is gitignored. Never commit real credentials.

---

## Getting Started

### Prerequisites

```bash
# Install Go tools
make init
```

This installs: `protoc-gen-go`, `protoc-gen-go-grpc`, `kratos`, `protoc-gen-go-http`, `protoc-gen-openapi`, `wire`.

### Build & Run

```bash
# Build binary
make build

# Run
./bin/SkyControl_gateway_API -conf ./configs
```

### Docker

```bash
# Build image
docker build -t skycontrol-gateway .

# Run container
docker run --rm \
  -p 8000:8000 \
  -p 9000:9000 \
  -v /path/to/your/configs:/data/conf \
  skycontrol-gateway
```

---

## Code Generation

### API & Config Protos (local)

```bash
make api       # Generate Go, HTTP, gRPC, OpenAPI from api/ proto files
make config    # Generate Go from internal/conf/conf.proto
make generate  # Run go generate + go mod tidy
make all       # api + config + generate
```

### Remote Service Contracts

Protobuf contracts are pulled from the [`SkyControlAPI_Contracts`](https://github.com/IdzAnAG1/SkyControlAPI_Contracts) repository and generated into `api/skycontrol/generated/`.

```bash
make generate_contracts   # Auth + Telemetry + Platform
make auth                 # Auth only
make telemetry            # Telemetry only
make platform             # Platform only
```

Requires [buf](https://buf.build/docs/installation) to be installed.

### Dependency Injection (Wire)

```bash
# Install wire
go install github.com/google/wire/cmd/wire@latest

# Regenerate wiring
cd cmd/SkyControl_gateway_API
wire
```

---

## Make Targets

```
make init                  Install protoc code generators
make api                   Generate API files from proto (Go, HTTP, gRPC, OpenAPI)
make config                Generate internal config from proto
make generate              Run go generate and tidy modules
make all                   Run api + config + generate
make generate_contracts    Pull and generate all remote service contracts
make auth                  Generate Auth service contracts
make telemetry             Generate Telemetry service contracts
make platform              Generate Platform service contracts
make build                 Build binary to ./bin/
make help                  Show all available targets
```

---

## CI/CD

GitHub Actions workflows are defined in `.github/workflows/`:

- **CI** — lint, test, Docker build and push on pull requests / pushes
- **CD** — deploy the container after a successful CI run

---

## Project Dependencies

| Package | Purpose |
|---------|---------|
| `github.com/go-kratos/kratos/v2` | Microservices framework (HTTP + gRPC) |
| `github.com/google/wire` | Compile-time dependency injection |
| `google.golang.org/grpc` | gRPC transport |
| `google.golang.org/protobuf` | Protocol Buffers runtime |
| `go.opentelemetry.io/otel` | Distributed tracing & metrics |
| `go.uber.org/automaxprocs` | Auto-set GOMAXPROCS to container CPU quota |
