![coverage](docs/coverage.svg) ![coverage](docs/time.svg)

# Calendarium

Calendarium is a backend system designed to manage meeting room reservations within an organization. It enables users to register, add and view meeting rooms, check room availability, and make or cancel bookings. The system ensures conflict-free scheduling, handles concurrent access safely, and supports bookings that span across multiple days.

## Project Structure

```
.
├── cmd/                    # Application entry points
│   ├── httpserver/        # HTTP server executable
│   └── migrate/           # Database migration tool
├── docs/                  # Documentation and OpenAPI specs
│   └── architecture.png   # Architecture diagram
├── internal/              # Private application code
│   └── book/             # Book domain module
│       ├── rest/         # HTTP handlers
│       └── store/        # Database implementations
├── pkg/                   # Public shared packages
│   ├── config/           # Configuration handling
│   ├── migration/        # Database migration utilities
│   ├── postgres/         # PostgreSQL client
│   └── testutil/         # Testing utilities
└── tools/                # Scripts and tools
    └── compose/          # Docker compose files
```

## Architecture

This project follows the Clean Architecture pattern (also known as Onion Architecture). See the architecture diagram in [docs/architecture.png](docs/architecture.png).

Key principles:
- Dependencies flow inward
- Inner layers contain business logic
- Outer layers contain implementation details
- Domain entities are at the core
- Each layer is isolated and testable

## Prerequisites

- Go 1.21+
- Docker and Docker Compose
- PostgreSQL 15

### Development Tools

- [air](https://github.com/air-verse/air) - Live reload for Go applications
- [golangci-lint](https://golangci-lint.run/) - Go linters aggregator
- [gotestsum](https://github.com/gotestyourself/gotestsum) - Better test output formatter
- [sql-migrate](https://github.com/rubenv/sql-migrate) - Database migration tool

## Getting Started

1. Clone the repository

    ```bash
    git clone https://github.com/ngoctrng/calendarium.git
    ```

2. Copy environment file and configure

    ```bash
    cp .env.example .env
    ```

3. Start dependencies

    ```bash
    make local-dev
    ```

4. Run database migrations

    ```bash
    make db/migrate
    ```

5. Start the server
```bash
go run cmd/httpserver/main.go
```

## Development

### Project Layout

- `cmd/` - Entry points for executables
- `internal/` - Private application code
- `pkg/` - Public shared packages
- `tools/` - Development and deployment tools
- `docs/` - Documentation and OpenAPI specs
- `documentation/` - Technical documentation (TBD)

### Testing

Run all tests:
```bash
make test
```

### Database Migrations

Create a new migration:
```bash
sql-migrate new -env="development" your-new-migration
```

### Development Tools

Hot reload during development:
```bash
make run
```

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request

## License

MIT