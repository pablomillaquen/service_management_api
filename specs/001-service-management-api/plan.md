# Implementation Plan: Service Management API

**Branch**: `main` | **Date**: 2026-06-07 | **Spec**: specs/001-service-management-api/spec.md

**Input**: Feature specification from `specs/001-service-management-api/spec.md`

**Note**: This template is filled in by the `/speckit.plan` command. See `.specify/templates/plan-template.md` for the execution workflow.

## Summary

Implementar una API REST para gestión de servicios técnicos (clientes, equipos,
órdenes de trabajo, materiales) utilizando Golang, MySQL y JWT Authentication,
desplegable mediante Docker Compose y cumpliendo los principios de Clean
Architecture definidos en la Constitución del proyecto.

## Technical Context

**Language/Version**: Golang 1.24+

**Primary Dependencies**: Gin (HTTP), GORM (ORM), golang-jwt, bcrypt, testify

**Storage**: MySQL 8+

**Testing**: Go Testing + Testify

**Target Platform**: Linux container (Docker), MySQL 8+ container

**Project Type**: web-service (REST API)

**Performance Goals**: < 500 ms promedio para operaciones CRUD

**Constraints**: Clean Architecture, JWT Auth + Refresh, RBAC, Soft Delete
en entidades principales, paginación obligatoria, auditoría obligatoria

**Scale/Scope**: ~3 roles (administrator, technician, viewer), 5 módulos
principales, API versionada (/api/v1)

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

| Principle | Status | Notes |
|-----------|--------|-------|
| I. Architecture First | ✅ Pass | Clean Architecture, capas separadas, DTOs |
| II. Security By Default | ✅ Pass | JWT, bcrypt, RBAC, validación, rate limiting |
| III. Database Ownership | ✅ Pass | Migraciones automáticas, FK, timestamps, soft delete |
| IV. API Consistency | ✅ Pass | /api/v1, formato unificado, paginación |
| V. Testing Requirements | ✅ Pass | Unit + Repository tests, 80% coverage |
| VI. Observability | ✅ Pass | Logging eventos + Audit Trail |
| VII. Documentation | ✅ Pass | Swagger/OpenAPI, README, .env.example |
| VIII. Dependency Management | ✅ Pass | Solo dependencias activas y justificadas |
| IX. Performance | ✅ Pass | < 500ms, índices, N+1 prevention, context |
| X. Production Readiness | ✅ Pass | Dockerfile, compose, migrations, seed |

**GATE RESULT**: ✅ ALL PRINCIPLES COMPLIANT. No violations to justify.

## Project Structure

### Documentation (this feature)

```text
specs/001-service-management-api/
├── plan.md              # This file (/speckit.plan command output)
├── research.md          # Phase 0 output (/speckit.plan command)
├── data-model.md        # Phase 1 output (/speckit.plan command)
├── quickstart.md        # Phase 1 output (/speckit.plan command)
├── contracts/           # Phase 1 output (/speckit.plan command)
└── tasks.md             # Phase 2 output (/speckit.tasks command)
```

### Source Code (repository root)

```text
cmd/
  api/                   # Application entry point

configs/                 # Configuration files

docs/                    # Swagger / OpenAPI specs

internal/
  domain/                # Domain entities & business rules
    user/
    client/
    equipment/
    workorder/
    material/
    audit/

  dto/                   # Data Transfer Objects

  handlers/              # HTTP handlers (controllers)

  middleware/            # Auth, RBAC, rate limiting, CORS

  repositories/          # Data access layer (GORM)

  services/              # Business logic layer

  validators/            # Custom validation rules

  auth/                  # JWT token management

pkg/
  database/              # Database connection & migration
  logger/                # Structured logging
  response/              # Unified API response format

migrations/              # SQL migration files

tests/                   # Integration & repository tests
```

**Structure Decision**: Se adopta la estructura Clean Architecture orientada a
dominio con separación explícita de capas (handlers → services → repositories).
Cada módulo de negocio (user, client, equipment, workorder, material, audit)
reside en `internal/domain/` con sus propios archivos de entidad. Las capas
transversales (dto, handlers, repositories, services) están separadas por
carpeta, no por módulo, para mantener la coherencia.

## Complexity Tracking

> No se requieren justificaciones — todas las decisiones de arquitectura
> están alineadas con la Constitución.
