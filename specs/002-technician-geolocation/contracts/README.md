# API Contracts: Technician Geolocation

## Base URL

```
http://localhost:8080/api/v1
```

## Authentication

All endpoints require `Authorization: Bearer <access_token>` header.

## Endpoints

### POST /technician-locations

Report current GPS location.

**Request**:
```json
{
  "latitude": -33.4569,
  "longitude": -70.6483
}
```

**Response (201)**:
```json
{
  "success": true,
  "message": "Location reported successfully",
  "data": {
    "id": 1,
    "user_id": 2,
    "latitude": -33.4569,
    "longitude": -70.6483,
    "created_at": "2026-06-08T14:30:00Z"
  }
}
```

**Errors**:
- `400` — Invalid coordinates (out of range)
- `401` — Missing or invalid token
- `403` — Viewer role not allowed

---

### GET /technician-locations

Query technician locations (paginated).

**Query Parameters**:
| Param | Type | Required | Description |
|-------|------|----------|-------------|
| `user_id` | int | No | Filter by technician ID |
| `start_date` | string | No | ISO 8601 datetime filter (inclusive) |
| `end_date` | string | No | ISO 8601 datetime filter (inclusive) |
| `page` | int | No | Page number (default: 1) |
| `per_page` | int | No | Items per page (default: 20) |

**Response (200)**:
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "user_id": 2,
      "latitude": -33.4569,
      "longitude": -70.6483,
      "created_at": "2026-06-08T14:30:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "per_page": 20,
    "total": 1
  }
}
```

**Errors**:
- `401` — Missing or invalid token
- `403` — Viewer role not allowed

### RBAC Rules

| Role | Create Location | Query Own | Query All |
|------|-----------------|-----------|-----------|
| Administrator | ❌ (cannot report) | ❌ | ✅ |
| Technician | ✅ (own only) | ✅ | ❌ |
| Viewer | ❌ | ❌ | ❌ |
