# Quickstart: Service Management API

## Prerequisites

- Docker & Docker Compose
- Go 1.24+ (for local development)

## Setup & Run

```bash
# 1. Start the application
docker compose up -d

# 2. Verify containers are running
docker compose ps
```

The API will be available at `http://localhost:8080`.

## Validation Scenarios

### 1. Authentication

```bash
# Login as admin
curl -s -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@example.com","password":"Admin123!"}' | jq .

# Expected: HTTP 200, access_token and refresh_token in response

# Extract token (save to variable)
TOKEN=$(curl -s -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@example.com","password":"Admin123!"}' \
  | jq -r '.data.access_token')

# Refresh token
REFRESH=$(curl -s -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@example.com","password":"Admin123!"}' \
  | jq -r '.data.refresh_token')

curl -s -X POST http://localhost:8080/api/v1/auth/refresh \
  -H "Content-Type: application/json" \
  -d "{\"refresh_token\":\"$REFRESH\"}" | jq .

# Expected: HTTP 200, new access_token
```

### 2. Client CRUD

```bash
# Create client
curl -s -X POST http://localhost:8080/api/v1/clients \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "business_name": "Empresa SAC",
    "tax_id": "20123456789",
    "primary_contact": "Carlos López",
    "email": "carlos@empresa.com",
    "phone": "+51999000111",
    "address": "Av. Principal 123, Lima"
  }' | jq .

# Expected: HTTP 201, client data returned

# List clients
curl -s http://localhost:8080/api/v1/clients?page=1\&per_page=20 \
  -H "Authorization: Bearer $TOKEN" | jq .

# Expected: HTTP 200, paginated client list

# Search by tax ID
curl -s "http://localhost:8080/api/v1/clients?search=20123456789" \
  -H "Authorization: Bearer $TOKEN" | jq .

# Expected: HTTP 200, matching client

# Reject duplicate tax ID
curl -s -X POST http://localhost:8080/api/v1/clients \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "business_name": "Otra Empresa",
    "tax_id": "20123456789",
    "primary_contact": "Maria Perez",
    "email": "maria@otra.com",
    "phone": "+51999000222",
    "address": "Av. Secundaria 456"
  }' | jq .

# Expected: HTTP 409, duplicate error
```

### 3. Equipment Catalog & Equipment

```bash
# Create equipment type
curl -s -X POST http://localhost:8080/api/v1/equipment-types \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"name":"Printer"}' | jq .

# Expected: HTTP 201

# Create brand
curl -s -X POST http://localhost:8080/api/v1/brands \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"name":"HP","equipment_type_id":1}' | jq .

# Expected: HTTP 201

# Create model
curl -s -X POST http://localhost:8080/api/v1/models \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"name":"LaserJet Pro M404dn","brand_id":1,"equipment_type_id":1}' | jq .

# Expected: HTTP 201

# Register equipment
curl -s -X POST http://localhost:8080/api/v1/equipment \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "client_id": 1,
    "model_id": 1,
    "serial_number": "SN-2024-001",
    "location": "Oficina 203",
    "status": "active"
  }' | jq .

# Expected: HTTP 201
```

### 4. Work Order Lifecycle

```bash
# Create work order
curl -s -X POST http://localhost:8080/api/v1/work-orders \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "client_id": 1,
    "equipment_id": 1,
    "description": "Impresora no enciende",
    "priority": "high",
    "scheduled_date": "2026-06-10"
  }' | jq .

# Expected: HTTP 201, status "pending"

# Assign technician (use technician user ID, e.g., 2)
curl -s -X POST http://localhost:8080/api/v1/work-orders/1/assign \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"technician_id": 2}' | jq .

# Expected: HTTP 200, status "assigned"

# Change status (as technician)
TECH_TOKEN=$(curl -s -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"tech@example.com","password":"TechPass1"}' \
  | jq -r '.data.access_token')

curl -s -X PATCH http://localhost:8080/api/v1/work-orders/1/status \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TECH_TOKEN" \
  -d '{"status":"in_progress"}' | jq .

# Expected: HTTP 200, status "in_progress"

# Add observation
curl -s -X POST http://localhost:8080/api/v1/work-orders/1/notes \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TECH_TOKEN" \
  -d '{"text":"Se encontró fusible quemado"}' | jq .

# Expected: HTTP 201

# Complete work order
curl -s -X PATCH http://localhost:8080/api/v1/work-orders/1/status \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TECH_TOKEN" \
  -d '{"status":"completed"}' | jq .

# Expected: HTTP 200, status "completed", completed_date set
```

### 5. Material Management & Consumption

```bash
# Create material
curl -s -X POST http://localhost:8080/api/v1/materials \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "code": "TON-001",
    "description": "Toner HP 85A",
    "unit_cost": 45.50
  }' | jq .

# Expected: HTTP 201

# Record material consumption on work order
curl -s -X POST http://localhost:8080/api/v1/work-orders/1/materials \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TECH_TOKEN" \
  -d '{"material_id": 1, "quantity": 1}' | jq .

# Expected: HTTP 201
```

### 6. Audit Trail

```bash
# View audit logs (admin only)
curl -s http://localhost:8080/api/v1/audit-logs?page=1\&per_page=10 \
  -H "Authorization: Bearer $TOKEN" | jq .

# Expected: HTTP 200, audit records for all previous operations
```

### 7. Authorization Enforcement

```bash
# Access without token
curl -s http://localhost:8080/api/v1/clients | jq .

# Expected: HTTP 401

# Viewer access attempt to create
VIEWER_TOKEN=$(curl -s -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"viewer@example.com","password":"ViewerPass1"}' \
  | jq -r '.data.access_token')

curl -s -X POST http://localhost:8080/api/v1/clients \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $VIEWER_TOKEN" \
  -d '{"business_name":"Test","tax_id":"999","primary_contact":"Test","email":"test@test.com","phone":"999","address":"Test"}' | jq .

# Expected: HTTP 403
```

### 8. Swagger Documentation

```bash
# Access Swagger UI
curl -s http://localhost:8080/swagger/index.html

# Expected: HTTP 200, Swagger UI page

# Access OpenAPI spec
curl -s http://localhost:8080/swagger/doc.json | jq .

# Expected: HTTP 200, OpenAPI JSON
```

## Expected Outcomes

| Scenario | Expected Result |
|----------|----------------|
| Login with valid credentials | HTTP 200 + tokens |
| Login with invalid credentials | HTTP 401 |
| Access without token | HTTP 401 |
| Create resource without permission | HTTP 403 |
| Create client with duplicate tax ID | HTTP 409 |
| Work order lifecycle (pending → completed) | All status transitions valid |
| Add observation | HTTP 201, observation is immutable |
| Audit log populated | Logs exist for each tracked action |
| Pagination | All collection endpoints paginated |
