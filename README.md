# Go REST API Template

[![Go Version](https://img.shields.io/badge/go-1.22+-blue.svg)](https://go.dev/)
[![CI](https://github.com/UlisesNiSchreiner/go-api-rest-template/actions/workflows/ci.yml/badge.svg)](https://github.com/UlisesNiSchreiner/go-api-rest-template/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/UlisesNiSchreiner/go-api-rest-template/branch/master/graph/badge.svg)](https://codecov.io/gh/UlisesNiSchreiner/go-api-rest-template)
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

## Template initialization

This repository is meant to be used as a **project template**.

Before starting development, you should initialize it with your own module name.  
This will update the Go module (`go.mod`) and replace template references in the README.

### Initialize the template

From the root of the repository, run:

```bash
go run scripts/init-template.go github.com/your-org/your-project
```

Example:

```bash
go run scripts/init-template.go github.com/acme/users-api
```

This will:

- Update the `module` name in `go.mod`
- Replace template references in `README.md`
- Prepare the project for first use

### After initialization

```bash
go mod tidy
git init
git commit -m "init project"
```

After that, the project is ready for development.

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

```bash
make dev
```

---

## Docker

```bash
docker build -t go-rest-template .
```

---

## Tests and coverage

```bash
go test ./...
```

CI enforces a **minimum 80% coverage** on pull requests.

---

## Project structure

```
cmd/api/
internal/
```

---

## CI

GitHub Actions runs tests, coverage and lint.

---

## License

MIT — see [LICENSE](LICENSE).
