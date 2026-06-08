# Service Management API

RESTful API for managing technical services, clients, equipment, work orders, and materials. Built with Go (Gin), GORM, MySQL 8+, and JWT authentication.

## Quick Start

```bash
docker compose up -d
```

The API will be available at `http://localhost:8080`. Swagger UI at `http://localhost:8080/swagger/index.html`.

### Default Admin Credentials

- **Email**: `admin@speckit.com`
- **Password**: `Admin123!`
- **Role**: `administrator`

## Architecture

Clean Architecture with dependency injection:

```
Handler → Service → Repository → Domain (GORM)
```

### Layers

| Layer | Responsibility |
|-------|---------------|
| **Handler** | HTTP request/response, validation |
| **Service** | Business logic, audit logging |
| **Repository** | Data access (GORM queries) |
| **Domain** | Entities, status transitions, business rules |
| **Middleware** | Auth (JWT), RBAC, security headers, CORS, rate limiting |

### Project Structure

```
cmd/api/          # Application entrypoint & routes
configs/          # Environment config
internal/
  domain/         # Entities (user, client, equipment, workorder, material, audit)
  dto/            # Request/response types
  handlers/       # HTTP handlers
  middleware/     # Auth, RBAC, security
  repositories/   # Database access
  services/       # Business logic
  auth/           # JWT + bcrypt
pkg/
  database/       # Connection, migration, seed
  logger/         # Structured logging
  response/       # Unified JSON response format
docs/             # Swagger/OpenAPI specs
tests/            # Integration tests
```

## Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `SERVER_PORT` | `8080` | API port |
| `DB_HOST` | `localhost` | MySQL host |
| `DB_PORT` | `3306` | MySQL port |
| `DB_USER` | `root` | Database user |
| `DB_PASSWORD` | (empty) | Database password |
| `DB_NAME` | `service_management` | Database name |
| `DB_MAX_OPEN_CONNS` | `25` | Max open connections |
| `DB_MAX_IDLE_CONNS` | `10` | Max idle connections |
| `DB_MAX_LIFETIME` | `5` | Connection max lifetime (minutes) |
| `JWT_SECRET` | `change-me-in-production` | JWT signing key |
| `RATE_LIMIT_PER_MINUTE` | `100` | Requests per minute per IP |

Copy `.env.example` to `.env` and customize for local development.

## API Overview

Base URL: `/api/v1`

### Auth Endpoints

| Method | Path | Auth | Role |
|--------|------|------|------|
| POST | `/auth/login` | No | — |
| POST | `/auth/refresh` | No | — |
| POST | `/auth/change-password` | Yes | Any |
| GET | `/health` | No | — |

### User Management

| Method | Path | Role |
|--------|------|------|
| GET | `/users` | administrator |
| GET | `/users/:id` | administrator |
| POST | `/users` | administrator |
| PUT | `/users/:id` | administrator |
| DELETE | `/users/:id` | administrator |

### Clients

| Method | Path | Role |
|--------|------|------|
| GET | `/clients` | administrator, technician |
| GET | `/clients/:id` | administrator, technician |
| POST | `/clients` | administrator |
| PUT | `/clients/:id` | administrator |
| DELETE | `/clients/:id` | administrator |

### Catalog (Equipment Types, Brands, Models)

| Method | Path | Role |
|--------|------|------|
| POST/GET | `/catalog/types` | administrator |
| GET/PUT/DELETE | `/catalog/types/:id` | administrator |
| POST/GET | `/catalog/brands` | administrator |
| GET | `/catalog/brands/by-type/:typeId` | administrator |
| GET/PUT/DELETE | `/catalog/brands/:id` | administrator |
| POST/GET | `/catalog/models` | administrator |
| GET | `/catalog/models/by-brand/:brandId` | administrator |
| GET/PUT/DELETE | `/catalog/models/:id` | administrator |

### Equipment

| Method | Path | Role |
|--------|------|------|
| GET | `/equipment` | administrator, technician |
| GET | `/equipment/:id` | administrator, technician |
| POST | `/equipment` | administrator |
| PUT | `/equipment/:id` | administrator |
| DELETE | `/equipment/:id` | administrator |

### Materials

| Method | Path | Role |
|--------|------|------|
| GET | `/materials` | administrator |
| GET | `/materials/:id` | administrator |
| POST | `/materials` | administrator |
| PUT | `/materials/:id` | administrator |
| DELETE | `/materials/:id` | administrator |

### Work Orders

| Method | Path | Role |
|--------|------|------|
| GET | `/work-orders` | administrator, technician |
| GET | `/work-orders/:id` | administrator, technician |
| POST | `/work-orders` | administrator, technician |
| PUT | `/work-orders/:id` | administrator, technician |
| DELETE | `/work-orders/:id` | administrator |
| POST | `/work-orders/:id/assign` | administrator |
| POST | `/work-orders/:id/status` | administrator, technician |
| POST | `/work-orders/:id/notes` | administrator, technician |
| POST | `/work-orders/:id/materials` | administrator, technician |

### Audit

| Method | Path | Role |
|--------|------|------|
| GET | `/audit` | administrator |

## Work Order Status Lifecycle

```
pending → assigned → in_progress → waiting_parts
                                  → completed
any → cancelled
```

## Response Format

Success:
```json
{ "success": true, "message": "...", "data": {} }
```

Paginated:
```json
{ "success": true, "message": "...", "data": [], "pagination": { "page": 1, "per_page": 10, "total": 100 } }
```

Error:
```json
{ "success": false, "message": "...", "errors": [] }
```

## Testing

```bash
# Unit tests only
CGO_ENABLED=1 go test -short ./...

# All tests (including integration)
CGO_ENABLED=1 go test ./...

# With coverage
CGO_ENABLED=1 go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Development

```bash
# Run locally (requires MySQL)
cp .env.example .env
# Edit .env with your MySQL credentials
go run ./cmd/api

# Build binary
go build -o server ./cmd/api

# Docker
docker compose up -d
```
