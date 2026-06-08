# Research: Service Management API

**Phase**: 0 — Outline & Research
**Date**: 2026-06-07

## Overview

All technology and architecture decisions were explicitly provided in the user's
implementation plan. No NEEDS CLARIFICATION items were identified.

## Decisions

### Language & Framework

- **Decision**: Golang 1.24+ with Gin framework
- **Rationale**: Especificado por el usuario. Gin es el framework HTTP más
  adoptado en el ecosistema Go, con buen rendimiento y madurez.
- **Alternatives considered**: N/A — especificado por el usuario.

### ORM & Database

- **Decision**: GORM + MySQL 8+
- **Rationale**: GORM es el ORM más utilizado en Go, con soporte para
  migraciones automáticas, soft delete, y relaciones. MySQL 8+ proporciona
  integridad referencial con claves foráneas y buen rendimiento.
- **Alternatives considered**: N/A — especificado por el usuario.

### Authentication

- **Decision**: JWT Access Token (15 min) + Refresh Token (7 días) + bcrypt
- **Rationale**: JWT permite stateless authentication, escalable y estándar.
  Refresh tokens habilitan renovación segura de sesiones. bcrypt cumple el
  requisito constitucional de almacenamiento seguro de contraseñas.
- **Alternatives considered**: N/A — especificado por el usuario.

### Authorization

- **Decision**: RBAC con 3 roles (administrator, technician, viewer)
- **Rationale**: Modelo simple y efectivo para los 3 perfiles definidos.
- **Alternatives considered**: N/A — especificado por el usuario.

### Architecture

- **Decision**: Clean Architecture con capas Handler → Service → Repository
- **Rationale**: Obligatorio por la Constitución (Principio I). Separación
  clara de responsabilidades, testabilidad, y desacoplamiento de frameworks.
- **Alternatives considered**: N/A — requerido por constitución.

### Testing

- **Decision**: Go Testing + Testify
- **Rationale**: Testify provee assertions y mocks sin dependencias externas
  pesadas. Go Testing es el estándar del lenguaje.
- **Alternatives considered**: N/A — especificado por el usuario.

### Deployment

- **Decision**: Docker Compose (API + MySQL)
- **Rationale**: Entorno reproducible, startup automático con migrations y
  seed data. Requerido por Constitución (Principio X).
- **Alternatives considered**: N/A — especificado por el usuario.

## Risks & Mitigations

| Risk | Mitigation |
|------|------------|
| GORM auto-migrations pueden perder datos en producción | Usar migrations versionadas en lugar de AutoMigrate para cambios destructivos |
| JWT sin revocación inmediata | Refresh tokens + short-lived access tokens (15 min) minimizan ventana de exposición |
| Crecimiento de audit log sin límite | El log es append-only por constitución; considerar rotación futura |
