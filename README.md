# GoStarter

A starter template for Go applications with support for HTTP services, Database, Cache, background Jobs, and Schedulers.

**Note:** 

⚠️ This project is currently under development and not production ready.

⚠️ This project is intended for personal/entertainment purposes only and not meant to be used in production environments.

## Features

- HTTP server using Fiber framework
- Background Jobs using Asynq
- Scheduled tasks with Asynq Scheduler
- Database using MySQL
- Cache using Redis
- Storage using Local and S3
- Mailer using SMTP and AWS SES
- Docker and Docker Compose ready
- Structured logging with Zap

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
4. Run the scheduler:
   ```bash
   go run main.go --svc=schedule
   ```

## Project Structure

```
├── Dockerfile
├── README.md
├── cmd
│   ├── cmd.go
│   ├── http
│   │   └── http.go
│   ├── scheduler
│   │   └── scheduler.go
│   └── worker
│       └── worker.go
├── config
│   ├── config.go
│   └── domain.go
├── docker-compose.yaml
├── go.mod
├── go.sum
├── internal
│   ├── example
│   │   ├── delivery
│   │   │   ├── http.go
│   │   │   ├── schedule.go
│   │   │   └── task.go
│   │   ├── domain
│   │   │   └── domain.go
│   │   ├── module.go
│   │   ├── repository
│   │   │   └── repository.go
│   │   └── usecase
│   │       └── usecase.go
│   └── helloworld
│       ├── delivery
│       │   ├── http.go
│       │   ├── schedule.go
│       │   └── task.go
│       ├── domain
│       │   └── domain.go
│       ├── module.go
│       ├── repository
│       │   └── repository.go
│       ├── templates
│       │   ├── mail
│       │   │   └── example.html
│       │   └── web
│       └── usecase
│           └── usecase.go
├── main.go
├── pkg
│   ├── adapter.go
│   ├── cache
│   │   └── redis.go
│   ├── db
│   │   ├── mysql.go
│   │   └── redis.go
│   ├── delivery.go
│   ├── interfaces
│   │   ├── cache.go
│   │   ├── mailer.go
│   │   ├── module.go
│   │   ├── queue.go
│   │   └── storage.go
│   ├── logging
│   │   └── logger.go
│   ├── mailer
│   │   ├── ses.go
│   │   └── smtp.go
│   ├── queue
│   │   └── asynq.go
│   ├── storage
│   │   └── s3.go
│   └── utils
│       ├── env.go
│       └── validation.go
└── shared
    ├── constant
    │   └── env.go
    └── middleware
        └── logging.go
```

## License

MIT
