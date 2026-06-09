# Current Architecture

Version: 1.0.0

## Modules

- **Auth** — Login, refresh token, change password
- **Users** — CRUD for administrators only
- **Clients** — CRUD for administrators and technicians
- **Equipment** — CRUD with unique serial number validation
- **Materials** — CRUD for administrators
- **Catalog** — Equipment types, brands, and models (administrators only)
- **WorkOrders** — CRUD, 6-state status machine, assign technician, add notes, register materials
- **Audit** — Immutable audit log for all operations

## Database

- **MySQL 8** via Docker Compose (`mysql:8.0`)
- GORM AutoMigrate for schema creation on startup
- SQLite in-memory for integration tests

## Authentication

- **JWT** — Access token (15 min) + Refresh token (7 days)
- Password hashing with **bcrypt**
- Password policy: minimum 8 chars, uppercase, lowercase, digit, special char

## Authorization

- **RBAC** with 3 roles:
  - `administrator` — Full access
  - `technician` — Assigned work orders + notes + materials
  - `viewer` — Read-only

## Patterns

- **Clean Architecture** — Domain → Repository → Service → Handler
- **Repository Pattern** — Data access abstraction
- **Service Layer** — Business logic encapsulation
- **Dependency Injection** — No global state, all dependencies wired at startup

## Infrastructure

- **Docker** — Multi-stage build (`golang:1.26-alpine`)
- **Docker Compose** — API + MySQL 8 containers
- Rate limiting in-memory per IP

## API

- **REST** — All endpoints under `/api/v1`
- Unified response format: `{success, message, data, errors, pagination}`
- **Swagger/OpenAPI** — Interactive docs at `/swagger/index.html`

## Testing

- **33 test functions** across 10 test files
- JWT, password, RBAC, status transition unit tests
- Service integration tests with isolated SQLite in-memory databases
- Docker image excludes tests (production build)

## Future Planned Modules

- Geolocation
- Route Optimization
- Notifications
- Kafka
