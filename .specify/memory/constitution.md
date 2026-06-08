<!--
  SYNC IMPACT REPORT
  Version change: (template/unversioned) → 1.0.0
  Modified principles: ALL — initialized from template placeholders
    - [PRINCIPLE_1_NAME] → I. Architecture First
    - [PRINCIPLE_2_NAME] → II. Security By Default
    - [PRINCIPLE_3_NAME] → III. Database Ownership
    - [PRINCIPLE_4_NAME] → IV. API Consistency
    - [PRINCIPLE_5_NAME] → V. Testing Requirements
  Added sections (within Core Principles):
    - VI. Observability
    - VII. Documentation
    - VIII. Dependency Management
    - IX. Performance
    - X. Production Readiness
  Removed sections: N/A
  Unused template slots (omitted): SECTION_2_NAME, SECTION_2_CONTENT, SECTION_3_NAME, SECTION_3_CONTENT
  Templates requiring updates:
    - .specify/templates/plan-template.md — ✅ no changes needed (generic Constitution Check gate)
    - .specify/templates/spec-template.md — ✅ no changes needed
    - .specify/templates/tasks-template.md — ✅ no changes needed
    - .specify/templates/checklist-template.md — ✅ no changes needed
  Follow-up TODOs: None
-->

# Service Management API Constitution

## Purpose

Este repositorio define los principios obligatorios para el desarrollo del sistema Service Management API.

Todas las especificaciones, planes de implementación y código generado deben cumplir esta constitución.

---

## Core Principles

### I. Architecture First

**Principle**: La arquitectura es más importante que la velocidad de desarrollo.

**Requirements**:
- Debe utilizarse Clean Architecture.
- La lógica de negocio no puede depender de frameworks.
- Los controladores HTTP no pueden contener lógica de negocio.
- El acceso a datos debe realizarse mediante repositorios.
- La comunicación entre capas debe utilizar DTOs.

**Layers**:

Required:
- Domain
- Repository
- Service
- Handler
- Middleware

Forbidden:
- Acceso directo a la base de datos desde handlers.
- Lógica de negocio en controllers.
- Dependencias circulares.

---

### II. Security By Default

**Principle**: Todo endpoint debe considerarse inseguro hasta ser protegido explícitamente.

**Requirements**:

Authentication:
- JWT Access Token
- Refresh Token

Password Storage:
- bcrypt

Authorization:
- Role Based Access Control (RBAC)

Input Protection:
- Request validation
- Sanitization
- SQL Injection prevention

Infrastructure:
- Security Headers
- CORS
- Rate Limiting
- Panic Recovery

Secrets:
- Nunca almacenar credenciales en código fuente.
- Todo valor sensible debe provenir de variables de entorno.

---

### III. Database Ownership

**Principle**: La aplicación es responsable de la creación y mantenimiento de su esquema.

**Requirements**:
- Todas las tablas deben generarse mediante migraciones.
- No asumir estructuras preexistentes.
- Toda modificación de esquema requiere una nueva migración.
- Las relaciones deben tener claves foráneas explícitas.
- Todas las tablas deben incluir timestamps.

**Required Fields**:
- `created_at`
- `updated_at`

**Soft Delete**: `deleted_at` para entidades principales.

---

### IV. API Consistency

**Principle**: Todas las APIs deben comportarse de forma uniforme.

**Requirements**:
- Versioning: `/api/v1`

**Success Response**:
```json
{
  "success": true,
  "message": "",
  "data": {}
}
```

**Error Response**:
```json
{
  "success": false,
  "message": "",
  "errors": []
}
```

**Pagination**:
```json
{
  "page": 1,
  "per_page": 20,
  "total": 100
}
```
- Todas las listas deben soportar paginación.

---

### V. Testing Requirements

**Principle**: El código no se considera terminado sin pruebas.

**Requirements**:

Minimum Coverage:
- 80% Services
- 80% Business Rules

Required Tests:
- Unit Tests
- Repository Tests

Optional:
- Integration Tests

Forbidden:
- Código productivo sin pruebas de negocio.

---

### VI. Observability

**Principle**: Toda acción importante debe poder ser auditada.

**Requirements**:

Logging:
- Startup events
- Authentication events
- Authorization failures
- Unexpected errors

Audit Trail — registrar:
- User
- Action
- Entity
- Entity ID
- Previous values
- New values
- Timestamp

Audit records nunca pueden ser eliminados.

---

### VII. Documentation

**Principle**: Toda funcionalidad debe estar documentada.

**Requirements**:

Generate:
- Swagger/OpenAPI
- README.md
- Environment Variables Guide

Each endpoint must include:
- Description
- Request example
- Response example
- Error examples

---

### VIII. Dependency Management

**Principle**: Minimizar dependencias externas.

**Requirements**:

Before adding a dependency:
1. Justify its usage.
2. Verify active maintenance.
3. Verify security status.

Forbidden:
- Abandoned libraries.
- Experimental libraries in production code.

---

### IX. Performance

**Principle**: La aplicación debe ser eficiente desde el primer despliegue.

**Requirements**:

Database:
- Índices para búsquedas frecuentes.
- Evitar N+1 queries.

API:
- Timeouts configurados.
- Context propagation.

Maximum Response Time:
- 500 ms promedio para operaciones CRUD.

---

### X. Production Readiness

**Principle**: Todo código generado debe poder desplegarse inmediatamente.

**Requirements**:

Mandatory Deliverables:
- Dockerfile
- docker-compose.yml
- Swagger
- Migrations
- Tests
- Seed Data
- Environment Example

The application must start successfully using `docker compose up -d` without requiring manual database creation.

---

## Governance

Todas las especificaciones futuras deben cumplir esta constitución.
Si existe conflicto entre una especificación y esta constitución, prevalece esta constitución.
Cambios a esta constitución requieren actualización explícita de versión y aprobación documentada.

**Version**: 1.0.0 | **Ratified**: 2026-06-07 | **Last Amended**: 2026-06-07
