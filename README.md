# GoStarter

A starter template for Go applications with support for HTTP services and background workers.

## Features

- HTTP server using Fiber framework
- Background worker using Asynq
- MySQL and Redis support
- Docker and Docker Compose ready

> **Note:** This project is currently under development and not production ready.

## Getting Started

### Prerequisites

- Go 1.24+
- Docker and Docker Compose
- MySQL
- Redis

### Running the Application

#### Using Docker

```bash
# Build and start all services
docker-compose up --build

# Stop all services
docker-compose down
```

#### Manually

1. Copy `.env.example` to `.env` and configure your environment variables
2. Run the HTTP server:
   ```bash
   go run main.go --svc=http
   ```
3. Run the worker:
   ```bash
   go run main.go --svc=worker
   ```

## Project Structure

```
├── cmd/                # Application entry points
│   ├── http/           # HTTP server
│   └── worker/         # Background worker
├── config/             # Configuration
├── internal/           # Application code
│   └── example/        # Example module
│       ├── delivery/   # HTTP handlers, tasks, schedules
│       ├── domain/     # Domain models and interfaces
│       ├── repository/ # Data access layer
│       └── usecase/    # Business logic
├── pkg/                # Shared packages
│   └── db/             # Database connections
├── shared/             # Shared utilities
└── main.go             # Main application entry point
```

## License

MIT
