# Data Model: Technician Geolocation

## Entity: TechnicianLocation

Represents a GPS location report from a technician during a service visit.

### Fields

| Field | Type | Constraints | Description |
|-------|------|-------------|-------------|
| `id` | uint64 | PK, auto-increment | Unique identifier |
| `user_id` | uint64 | FK → users.id, NOT NULL, indexed | Technician who reported location |
| `latitude` | float64 | NOT NULL, range: -90 to 90 | GPS latitude |
| `longitude` | float64 | NOT NULL, range: -180 to 180 | GPS longitude |
| `created_at` | datetime | NOT NULL, auto-set | Server-generated timestamp |

### Indexes

- `idx_user_id_created_at` on (user_id, created_at DESC) — optimize queries by technician + date ordering
- `idx_created_at` on (created_at DESC) — optimize global date range queries

### Relationships

- **TechnicianLocation belongs_to User**: `user_id` references `users.id`
- No cascade delete on technician soft-delete (location history is preserved)

### Validation Rules

- Latitude: -90.0 to 90.0 (inclusive)
- Longitude: -180.0 to 180.0 (inclusive)
- user_id: must reference an existing active user with role "technician"

### Notes

- No update or delete operations (append-only log)
- No soft delete — records are immutable once created
- Timestamps are server-generated, not client-supplied
