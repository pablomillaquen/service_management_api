# Quickstart: Technician Geolocation

## Prerequisites

- Docker y Docker Compose instalados
- API corriendo en `http://localhost:8080` (ver `docker compose up -d` en raíz del proyecto)
- Token JWT de administrador: `POST /api/v1/auth/login` con `admin@speckit.com` / `Admin123!`

## Endpoints a validar

### 1. Crear un técnico (admin)

```bash
curl -s -X POST http://localhost:8080/api/v1/users \
  -H 'Authorization: Bearer <admin_token>' \
  -H 'Content-Type: application/json' \
  -d '{"name":"Técnico Uno","email":"tecnico1@test.com","password":"Pass123!","role":"technician"}'
```

**Esperado**: `201 Created`, usuario creado con id.

### 2. Reportar ubicación como técnico

```bash
TECH_TOKEN=$(curl -s -X POST http://localhost:8080/api/v1/auth/login \
  -H 'Content-Type: application/json' \
  -d '{"email":"tecnico1@test.com","password":"Pass123!"}' | python3 -c "import sys,json; print(json.load(sys.stdin)['data']['access_token'])")

curl -s -X POST http://localhost:8080/api/v1/technician-locations \
  -H "Authorization: Bearer $TECH_TOKEN" \
  -H 'Content-Type: application/json' \
  -d '{"latitude":-33.4569,"longitude":-70.6483}'
```

**Esperado**: `201 Created`, location registrada con user_id del técnico y timestamp.

### 3. Validar que técnico no puede reportar para otro

```bash
curl -s -X POST http://localhost:8080/api/v1/technician-locations \
  -H "Authorization: Bearer $TECH_TOKEN" \
  -H 'Content-Type: application/json' \
  -d '{"user_id":999,"latitude":-33.4569,"longitude":-70.6483}'
```

**Esperado**: `403 Forbidden`, el user_id se ignora y siempre se usa el del token.

### 4. Validar coordenadas inválidas

```bash
curl -s -X POST http://localhost:8080/api/v1/technician-locations \
  -H "Authorization: Bearer $TECH_TOKEN" \
  -H 'Content-Type: application/json' \
  -d '{"latitude":-100,"longitude":-70.6483}'
```

**Esperado**: `400 Bad Request`, error de validación.

### 5. Consultar ubicaciones como administrador

```bash
ADMIN_TOKEN=$(curl -s -X POST http://localhost:8080/api/v1/auth/login \
  -H 'Content-Type: application/json' \
  -d '{"email":"admin@speckit.com","password":"Admin123!"}' | python3 -c "import sys,json; print(json.load(sys.stdin)['data']['access_token'])")

curl -s http://localhost:8080/api/v1/technician-locations \
  -H "Authorization: Bearer $ADMIN_TOKEN"
```

**Esperado**: `200 OK`, lista paginada con las ubicaciones del técnico.

### 6. Filtrar por técnico y rango de fechas

```bash
curl -s "http://localhost:8080/api/v1/technician-locations?user_id=2&start_date=2026-06-01T00:00:00Z&end_date=2026-06-30T23:59:59Z" \
  -H "Authorization: Bearer $ADMIN_TOKEN"
```

**Esperado**: `200 OK`, solo ubicaciones del técnico 2 en junio.

### 7. Técnico consulta solo sus propias ubicaciones

```bash
curl -s http://localhost:8080/api/v1/technician-locations \
  -H "Authorization: Bearer $TECH_TOKEN"
```

**Esperado**: `200 OK`, solo ubicaciones del técnico autenticado.

### 8. Viewer no puede acceder

```bash
# Crear usuario viewer si no existe
curl -s -X POST http://localhost:8080/api/v1/users \
  -H "Authorization: Bearer $ADMIN_TOKEN" \
  -H 'Content-Type: application/json' \
  -d '{"name":"Visor","email":"visor@test.com","password":"Pass123!","role":"viewer"}'

VIEWER_TOKEN=$(curl -s -X POST http://localhost:8080/api/v1/auth/login \
  -H 'Content-Type: application/json' \
  -d '{"email":"visor@test.com","password":"Pass123!"}' | python3 -c "import sys,json; print(json.load(sys.stdin)['data']['access_token'])")

# Viewer intenta crear ubicación
curl -s -X POST http://localhost:8080/api/v1/technician-locations \
  -H "Authorization: Bearer $VIEWER_TOKEN" \
  -H 'Content-Type: application/json' \
  -d '{"latitude":-33.4569,"longitude":-70.6483}'

# Viewer intenta consultar
curl -s http://localhost:8080/api/v1/technician-locations \
  -H "Authorization: Bearer $VIEWER_TOKEN"
```

**Esperado**: Ambos endpoints retornan `403 Forbidden`.

## Resumen de validación

| # | Escenario | Método | Ruta | Código esperado |
|---|-----------|--------|------|-----------------|
| 1 | Crear técnico | POST | /api/v1/users | 201 |
| 2 | Reportar ubicación | POST | /api/v1/technician-locations | 201 |
| 3 | Reportar para otro | POST | /api/v1/technician-locations | 403 |
| 4 | Coordenadas inválidas | POST | /api/v1/technician-locations | 400 |
| 5 | Admin consulta todo | GET | /api/v1/technician-locations | 200 |
| 6 | Admin filtra | GET | /api/v1/technician-locations?user_id=&start_date=&end_date= | 200 |
| 7 | Técnico consulta propio | GET | /api/v1/technician-locations | 200 |
| 8 | Viewer sin acceso | POST+GET | /api/v1/technician-locations | 403 |
