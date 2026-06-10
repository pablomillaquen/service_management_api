# Research: Technician Geolocation

## Unknowns Resolved

No NEEDS CLARIFICATION markers existed in the spec. All requirements were unambiguous.

## Technology Decisions

| Decision | Choice | Rationale | Alternatives Considered |
|----------|--------|-----------|------------------------|
| Storage format | MySQL table `technician_locations` | Proyecto existente usa MySQL 8 con GORM. No se justifica base separada | PostgreSQL, Redis, MongoDB — descartados por sobreingeniería para este alcance |
| Coordinate validation | Application-level (handler) | GORM no soporta CHECK constraints cross-dialect; validación inline es el patrón existente | MySQL CHECK constraints, GORM hooks |
| Index strategy | Composite index (user_id, created_at) | Optimiza las queries principales: filtro por técnico + ordenamiento por fecha | Single index on user_id, separate on created_at |
| RBAC reuse | Middleware existente | `rbacMiddleware("administrator", "technician")` ya soporta múltiples roles | Nuevo middleware específico — innecesario |
| Soft delete | No aplica | Location data es append-only sin delete. No se requiere soft delete | Soft delete — descartado (no hay delete endpoint) |

## Best Practices

- Timestamps deben ser generados por el servidor (no por el cliente) para prevenir manipulación
- Coordenadas GPS: latitud rango -90 a 90, longitud rango -180 a 180
- Orden descendente por defecto (más reciente primero)
- Paginación obligatoria (mismo estándar del proyecto existente)
