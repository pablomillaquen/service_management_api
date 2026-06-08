# API Contracts: Service Management API

Base URL: `/api/v1`

All endpoints require `Authorization: Bearer <access_token>` header except
`POST /auth/login` and `POST /auth/refresh`.

## Response Format

**Success**:
```json
{
  "success": true,
  "message": "Operation successful",
  "data": {}
}
```

**Error**:
```json
{
  "success": false,
  "message": "Error description",
  "errors": []
}
```

**Paginated**:
```json
{
  "success": true,
  "message": "",
  "data": [],
  "pagination": {
    "page": 1,
    "per_page": 20,
    "total": 100
  }
}
```

## Authentication

### POST /auth/login
Authenticate user with email and password.

**Request**:
```json
{
  "email": "admin@example.com",
  "password": "Admin123!"
}
```

**Response 200**:
```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIs...",
    "refresh_token": "dGhpcyBpcyBhIHJlZnJl...",
    "expires_in": 900,
    "user": {
      "id": 1,
      "name": "Admin",
      "email": "admin@example.com",
      "role": "administrator"
    }
  }
}
```

**Response 401**:
```json
{
  "success": false,
  "message": "Invalid credentials",
  "errors": ["Email or password is incorrect"]
}
```

### POST /auth/refresh
Renew access token using refresh token.

**Request**:
```json
{
  "refresh_token": "dGhpcyBpcyBhIHJlZnJl..."
}
```

**Response 200**:
```json
{
  "success": true,
  "message": "Token refreshed",
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIs...",
    "refresh_token": "bmV3IHJlZnJlc2ggdG9r...",
    "expires_in": 900
  }
}
```

### POST /auth/logout
Invalidate current session.

**Headers**: Authorization: Bearer \<access_token\>

**Response 200**:
```json
{
  "success": true,
  "message": "Logout successful",
  "data": null
}
```

### POST /auth/change-password
Change authenticated user's password.

**Headers**: Authorization: Bearer \<access_token\>

**Request**:
```json
{
  "current_password": "OldPass1",
  "new_password": "NewPass1"
}
```

**Response 200**:
```json
{
  "success": true,
  "message": "Password changed successfully",
  "data": null
}
```

## Users

### GET /api/v1/users
List users (paginated). | Administrator

**Query**: ?page=1&per_page=20&search=&role=

### GET /api/v1/users/:id
Get user by ID. | Administrator

### POST /api/v1/users
Create user. | Administrator

**Request**:
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "SecurePass1",
  "role": "technician"
}
```

### PUT /api/v1/users/:id
Update user. | Administrator

### DELETE /api/v1/users/:id
Deactivate user (soft delete). | Administrator

## Clients

### GET /api/v1/clients
List clients (paginated, searchable by name or tax ID). | Administrator, Viewer

### GET /api/v1/clients/:id
Get client by ID. | Administrator, Viewer

### POST /api/v1/clients
Create client. | Administrator

**Request**:
```json
{
  "business_name": "Empresa SAC",
  "tax_id": "20123456789",
  "primary_contact": "Carlos López",
  "email": "carlos@empresa.com",
  "phone": "+51999000111",
  "address": "Av. Principal 123, Lima"
}
```

### PUT /api/v1/clients/:id
Update client. | Administrator

### DELETE /api/v1/clients/:id
Soft delete client. | Administrator

## Equipment Catalog

### GET /api/v1/equipment-types
List equipment types. | Administrator, Technician, Viewer

### POST /api/v1/equipment-types
Create equipment type. | Administrator

### PUT /api/v1/equipment-types/:id
Update equipment type. | Administrator

### DELETE /api/v1/equipment-types/:id
Delete equipment type. | Administrator

### GET /api/v1/equipment-types/:id/brands
List brands for a type. | Administrator, Technician, Viewer

### POST /api/v1/brands
Create brand. | Administrator

### PUT /api/v1/brands/:id
Update brand. | Administrator

### DELETE /api/v1/brands/:id
Delete brand. | Administrator

### GET /api/v1/brands/:id/models
List models for a brand. | Administrator, Technician, Viewer

### POST /api/v1/models
Create model. | Administrator

### PUT /api/v1/models/:id
Update model. | Administrator

### DELETE /api/v1/models/:id
Delete model. | Administrator

## Equipment

### GET /api/v1/equipment
List equipment (paginated, searchable by serial number or client). | Administrator, Technician, Viewer

### GET /api/v1/equipment/:id
Get equipment by ID. | Administrator, Technician, Viewer

### POST /api/v1/equipment
Register equipment. | Administrator

**Request**:
```json
{
  "client_id": 1,
  "model_id": 5,
  "serial_number": "SN-2024-001",
  "location": "Oficina 203, Piso 2",
  "status": "active"
}
```

### PUT /api/v1/equipment/:id
Update equipment. | Administrator

### DELETE /api/v1/equipment/:id
Soft delete equipment. | Administrator

## Materials

### GET /api/v1/materials
List materials (paginated). | Administrator, Technician, Viewer

### GET /api/v1/materials/:id
Get material by ID. | Administrator, Technician, Viewer

### POST /api/v1/materials
Create material. | Administrator

**Request**:
```json
{
  "code": "TON-001",
  "description": "Toner HP 85A",
  "unit_cost": 45.50
}
```

### PUT /api/v1/materials/:id
Update material. | Administrator

### DELETE /api/v1/materials/:id
Soft delete material. | Administrator

## Work Orders

### GET /api/v1/work-orders
List work orders (paginated, filterable by status, priority, client, technician, date range). | Administrator, Technician, Viewer

### GET /api/v1/work-orders/:id
Get work order by ID. | Administrator, Technician, Viewer

### POST /api/v1/work-orders
Create work order. | Administrator

**Request**:
```json
{
  "client_id": 1,
  "equipment_id": 3,
  "description": "Impresora no enciende, revisar fuente de poder",
  "priority": "high",
  "scheduled_date": "2026-06-10"
}
```

### PUT /api/v1/work-orders/:id
Update work order. | Administrator

### DELETE /api/v1/work-orders/:id
Soft delete work order. | Administrator

### POST /api/v1/work-orders/:id/assign
Assign technician. | Administrator

**Request**:
```json
{
  "technician_id": 5
}
```

**Response 200**:
```json
{
  "success": true,
  "message": "Work order assigned",
  "data": {
    "id": 10,
    "status": "assigned",
    "technician": {
      "id": 5,
      "name": "John Doe"
    },
    "assigned_by": {
      "id": 1,
      "name": "Admin"
    },
    "assigned_at": "2026-06-07T14:30:00Z"
  }
}
```

### PATCH /api/v1/work-orders/:id/status
Change work order status. | Administrator (all statuses), Technician (limited)

**Request**:
```json
{
  "status": "in_progress"
}
```

### POST /api/v1/work-orders/:id/notes
Add observation. | Administrator, Technician

**Request**:
```json
{
  "text": "Se encontró fusible quemado, se requiere repuesto"
}
```

### POST /api/v1/work-orders/:id/materials
Record material consumption. | Technician

**Request**:
```json
{
  "material_id": 1,
  "quantity": 2
}
```

## Audit

### GET /api/v1/audit-logs
List audit logs (paginated, filterable by entity, entity_id, user, action, date range). | Administrator

## Error Codes

| HTTP Status | Code | Description |
|-------------|------|-------------|
| 400 | VALIDATION_ERROR | Invalid input data |
| 401 | AUTHENTICATION_ERROR | Missing or invalid token |
| 403 | AUTHORIZATION_ERROR | Insufficient permissions |
| 404 | NOT_FOUND | Resource not found |
| 409 | CONFLICT | Duplicate unique field |
| 500 | INTERNAL_ERROR | Unexpected server error |
