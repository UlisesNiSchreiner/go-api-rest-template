# Go REST API Template (Layered Architecture)

[![CI](https://github.com/your-org/go-rest-layered-template/actions/workflows/ci.yml/badge.svg)](https://github.com/your-org/go-rest-layered-template/actions/workflows/ci.yml)
[![Go Version](https://img.shields.io/badge/go-1.22+-blue.svg)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-brightgreen.svg)](LICENSE)

A market-ready starting point for REST APIs in Go using a pragmatic layered architecture:

- **Handlers (HTTP)**: transport concerns (routing, parsing, rendering)
- **Services**: business logic
- **Repositories**: data access (MySQL example)
- **Platform**: infrastructure wiring (DB, config, logging)

Sample endpoints:

- `GET /v1/health` – health check
- `GET /v1/users/{id}` – fetch a user from MySQL
- `GET /swagger` – Swagger UI
- `GET /swagger/openapi.yaml` – OpenAPI spec

---

## Requirements

- Go **1.22+**
- MySQL **8+** (or compatible)

---

## Quick start (local, without Docker)

1) Export environment variables:

```bash
export MYSQL_DSN="user:password@tcp(127.0.0.1:3306)/app?parseTime=true"
export HTTP_HOST="0.0.0.0"
export HTTP_PORT="8080"
```

2) Run the API:

```bash
go run ./cmd/api
```

3) Try it:

```bash
curl -s http://localhost:8080/v1/health
curl -s http://localhost:8080/v1/users/1
```

Open docs:

```bash
open http://localhost:8080/swagger
```

### MySQL schema

The `users` endpoint expects:

```sql
CREATE TABLE users (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  email VARCHAR(255) NOT NULL,
  name VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
```

---

## Development mode (watch changes)

This template uses **Air** for rebuild/restart on file changes:

```bash
make dev
```

`make dev` installs `air` into `./bin/` automatically.

---

## Docker

### Build and run

```bash
docker build -t go-rest-template .
docker run --rm -p 8080:8080 \
  -e APP_ENV=prod \
  -e HTTP_HOST=0.0.0.0 \
  -e HTTP_PORT=8080 \
  -e MYSQL_DSN="user:password@tcp(host.docker.internal:3306)/app?parseTime=true" \
  go-rest-template
```

### Optional: local MySQL via docker-compose

```bash
docker compose up -d mysql
```

---

## Tests and coverage

Run tests:

```bash
go test ./...
```

Coverage report:

```bash
make test
make cover
```

CI enforces a **minimum 80% coverage** on pull requests.

---

## Project structure

```
cmd/api/                         # main entrypoint
internal/config/                 # env-based configuration
internal/domain/                 # domain models
internal/handlers/               # HTTP handlers + middleware
internal/services/               # business logic
internal/repositories/           # repository interfaces
internal/repositories/mysqlrepo/ # MySQL repositories (database/sql)
internal/platform/db/            # DB bootstrap (sql.DB)
internal/logger/                 # zap logger wrapper
```

---

## CI

GitHub Actions runs:

- `go test` with coverage + coverage gate
- `golangci-lint` static analysis

See `.github/workflows/ci.yml`.

---

## License

MIT — see [LICENSE](LICENSE).
