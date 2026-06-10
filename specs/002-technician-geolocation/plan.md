# Implementation Plan: Technician Geolocation

**Branch**: `002-technician-geolocation` | **Date**: 2026-06-08 | **Spec**: specs/002-technician-geolocation/spec.md

**Input**: Feature specification from `specs/002-technician-geolocation/spec.md`

## Summary

Añadir capacidades de geolocalización de técnicos: endpoint para que técnicos reporten coordenadas GPS (latitud, longitud) y endpoint para que administradores consulten ubicaciones. Solo el propio técnico puede crear sus registros de ubicación. Los administradores pueden consultar todas. La implementación sigue los patrones de Clean Architecture del proyecto existente.

## Technical Context

**Language/Version**: Golang 1.26+

**Primary Dependencies**: Gin (HTTP), GORM (ORM), golang-jwt, bcrypt, testify

**Storage**: MySQL 8+

**Testing**: Go Testing + Testify

**Target Platform**: Linux container (Docker), MySQL 8+ container

**Project Type**: web-service (REST API)

**Performance Goals**: < 500 ms para consultas de ubicaciones

**Constraints**: Clean Architecture, JWT Auth + RBAC, mismo patrón del feature existente (domain → repository → service → handler)

**Scale/Scope**: ~50 técnicos, reportes cada 15-30 min por visita

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

| Principle | Status | Notes |
|-----------|--------|-------|
| I. Architecture First | ✅ Pass | Clean Architecture, capas separadas, mismo patrón existente |
| II. Security By Default | ✅ Pass | JWT + RBAC (technician solo crea sus registros, admin consulta todos) |
| III. Database Ownership | ✅ Pass | GORM AutoMigrate, nueva tabla `technician_locations` con timestamps |
| IV. API Consistency | ✅ Pass | `/api/v1` + formato unificado + paginación |
| V. Testing Requirements | ✅ Pass | Unit + integration tests con Testify |
| VI. Observability | ✅ Pass | Logging de reportes de ubicación + audit trail |
| VII. Documentation | ✅ Pass | Swagger/OpenAPI para nuevos endpoints |
| VIII. Dependency Management | ✅ Pass | Sin nuevas dependencias externas |
| IX. Performance | ✅ Pass | Índice por user_id + created_at, < 500ms |
| X. Production Readiness | ✅ Pass | Migración automática, seed data existente |

**GATE RESULT**: ✅ ALL PRINCIPLES COMPLIANT. No violations to justify.

## Project Structure

### Documentation (this feature)

```text
specs/002-technician-geolocation/
├── plan.md              # This file
├── research.md          # Phase 0 output
├── data-model.md        # Phase 1 output
├── quickstart.md        # Phase 1 output
├── contracts/           # Phase 1 output
├── checklists/
│   └── requirements.md  # Spec quality checklist
└── spec.md              # Feature specification
```

### Source Code (repository root)

```text
internal/
├── domain/
│   └── technicianlocation/   # New domain entity
│       └── location.go
├── dto/
│   └── technicianlocation.go  # New DTOs
├── handlers/
│   └── technician_location_handler.go  # New handler
├── repositories/
│   └── technician_location_repository.go  # New repository
├── services/
│   └── technician_location_service.go  # New service
├── middleware/
│   └── ...  # Reuse existing auth + RBAC
├── auth/
│   └── ...  # Reuse existing JWT
└── domain/
    └── ...  # Reuse existing entities

cmd/api/
├── main.go    # Wire new dependencies
└── routes.go  # Register new routes

tests/
├── services/
│   └── technician_location_service_test.go
└── test_helper.go  # Reuse existing test helper
```

**Structure Decision**: Se adopta la misma estructura Clean Architecture del feature existente. El nuevo módulo `technicianlocation` sigue el mismo patrón: domain entity → repository → service → handler. No se requieren cambios en infraestructura existente.

## Complexity Tracking

> Sin violaciones constitucionales que justificar.
