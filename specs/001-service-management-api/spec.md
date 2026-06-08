# Feature Specification: Service Management API

**Feature Branch**: `001-service-management-api`

**Created**: 2026-06-07

**Status**: Draft

**Input**: User description: "Service Management API"

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Authentication & User Administration (Priority: P1)

As an Administrator, I want to manage user accounts with roles so that the right
people have secure access to the system. As any user, I want to authenticate and
maintain my session securely.

**Why this priority**: Authentication and user administration are foundational —
every other feature depends on users being authenticated and properly authorized.

**Independent Test**: Can be fully tested by creating an administrator account,
logging in with valid credentials, receiving tokens, and verifying that only
authenticated users can access protected endpoints.

**Acceptance Scenarios**:

1. **Given** an administrator user exists, **When** I log in with valid email and
   password, **Then** I receive an access token and a refresh token.
2. **Given** an authenticated session, **When** my access token expires and I
   present a valid refresh token, **Then** I receive a new access token.
3. **Given** an authenticated administrator, **When** I create a new user with
   name, email, password, and role, **Then** the user is stored with status
   "active" and can log in.
4. **Given** an existing active user, **When** an administrator deactivates the
   user, **Then** the user can no longer log in.

---

### User Story 2 - Client & Equipment Management (Priority: P1)

As an Administrator, I want to register clients and their installed equipment so
that work orders can reference accurate asset information.

**Why this priority**: Clients and equipment are the core business entities —
work orders cannot exist without them.

**Independent Test**: Can be fully tested by creating a client, registering
multiple equipment types and models, assigning equipment to a client, and
searching by tax ID or serial number.

**Acceptance Scenarios**:

1. **Given** I am an authenticated administrator, **When** I register a new
   client with business name, tax ID, contact, email, phone, and address,
   **Then** the client is stored and searchable by name or tax ID.
2. **Given** an existing client, **When** I attempt to register a second client
   with the same tax ID, **Then** the system rejects the duplicate.
3. **Given** equipment types (Printer, Scanner, POS, Notebook) and brands are
   defined, **When** I register equipment with model, serial number, client,
   location, and status, **Then** the equipment is stored and searchable by
   serial number or client.
4. **Given** an existing equipment record, **When** I try to register another
   equipment with the same serial number, **Then** the system rejects the
   duplicate.

---

### User Story 3 - Work Order Lifecycle (Priority: P1)

As an Administrator, I want to create work orders, assign them to technicians,
and track their progress. As a Technician, I want to view my assigned orders,
update their status, add observations, and complete them.

**Why this priority**: Work orders are the central operational feature — they
coordinate all field activities and are the primary reason for the system.

**Independent Test**: Can be fully tested by creating a work order for a
client/equipment, assigning a technician, the technician updating status through
the lifecycle, adding observations, and completing the order.

**Acceptance Scenarios**:

1. **Given** an existing client and equipment, **When** an administrator creates
   a work order with description, priority, and scheduled date, **Then** the
   order is created with status "pending".
2. **Given** a pending work order, **When** an administrator assigns a
   technician, **Then** the order status changes to "assigned" and the
   assignment is recorded with assignor, technician, and date.
3. **Given** an assigned work order, **When** the technician starts work,
   **Then** the status changes to "in_progress".
4. **Given** an in-progress work order, **When** the technician marks it as
   needing parts, **Then** the status changes to "waiting_parts".
5. **Given** an in-progress work order, **When** the technician completes it,
   **Then** the status changes to "completed" and a completion date is
   recorded.
6. **Given** an active work order, **When** an administrator cancels it,
   **Then** the status changes to "cancelled".
7. **Given** any work order with status other than the six allowed values,
   **When** a status change is attempted, **Then** the system rejects it.
8. **Given** an active work order, **When** an authorized user adds an
   observation with text, **Then** the observation is stored with author,
   date, and the text, and cannot be deleted.

---

### User Story 4 - Materials & Consumption Tracking (Priority: P2)

As an Administrator, I want to manage a material catalog. As a Technician,
I want to record materials used during a work order so that consumption is
tracked accurately.

**Why this priority**: Material tracking adds significant operational value but
work orders can function without it initially.

**Independent Test**: Can be fully tested by creating materials in the catalog,
then recording consumption against a work order.

**Acceptance Scenarios**:

1. **Given** I am an authenticated administrator, **When** I create a material
   with code, description, and unit cost, **Then** the material is stored and
   the code is unique.
2. **Given** an existing work order and materials, **When** a technician records
   material usage with quantity, **Then** the consumption is stored with
   material, quantity, responsible user, and date.

---

### User Story 5 - Search, Audit & Documentation (Priority: P3)

As any authenticated user, I want to search across records, view the audit
trail of changes, and access API documentation.

**Why this priority**: Search improves usability, audit trail satisfies
compliance, and documentation enables integration — all valuable but not
critical for initial operations.

**Independent Test**: Can be fully tested by performing filtered searches on
work orders and clients, verifying audit records are generated on modification,
and accessing the API documentation endpoint.

**Acceptance Scenarios**:

1. **Given** multiple work orders exist, **When** I search by status, priority,
   client, technician, or date range, **Then** I receive paginated results
   matching the filters.
2. **Given** multiple clients exist, **When** I search by name or tax ID,
   **Then** I receive matching results.
3. **Given** multiple equipment records exist, **When** I search by serial
   number or client, **Then** I receive matching results.
4. **Given** a record is created, updated, or logically deleted, **Then** an
   audit entry is stored with user, action, entity, entity ID, date, previous
   values, and new values.
5. **Given** a work order status or assignment changes, **Then** an audit entry
   is stored for that change.
6. **Given** I access the API documentation endpoint, **Then** I receive the
   complete OpenAPI/Swagger documentation for all endpoints.

### Edge Cases

- What happens when a user attempts to log in with incorrect credentials more
  than 5 times?
- How does the system handle assigning a work order to a deactivated
  technician?
- What happens when a material with a duplicate code is submitted?
- How does the system handle concurrent updates to the same work order?
- What happens when a work order in "completed" or "cancelled" status receives
  a status change request?
- How are pagination parameters handled when page or per_page exceed reasonable
  limits?

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: System MUST authenticate users via email and password, issuing
  JWT access tokens and refresh tokens.
- **FR-002**: System MUST allow token renewal using a valid refresh token.
- **FR-003**: System MUST allow users to log out, invalidating their session.
- **FR-004**: System MUST allow authenticated users to change their password.
- **FR-005**: Administrators MUST be able to create, modify, deactivate, and
  list users with name, email, password, role, and active status.
- **FR-006**: Administrators MUST be able to register clients with business
  name, tax ID, primary contact, email, phone, and address.
- **FR-007**: System MUST reject duplicate clients by tax ID.
- **FR-008**: System MUST support an equipment catalog with types (Printer,
  Scanner, POS, Notebook), brands (per type), and models (per brand and type).
- **FR-009**: Administrators MUST be able to register equipment with owner
  client, model, serial number, physical location, and status.
- **FR-010**: System MUST enforce unique serial numbers for equipment.
- **FR-011**: Administrators MUST be able to create work orders with client,
  equipment, description, priority, status, scheduled date, and assigned
  technician.
- **FR-012**: Administrators MUST be able to assign work orders to technicians,
  recording the assignor, technician, and assignment date.
- **FR-013**: Work orders MUST support exactly six lifecycle states: pending,
  assigned, in_progress, waiting_parts, completed, cancelled.
- **FR-014**: Authorized users MUST be able to add non-removable observations
  to work orders, recording author, date, and text.
- **FR-015**: Administrators MUST be able to create materials with unique code,
  description, and unit cost.
- **FR-016**: Technicians MUST be able to record material consumption on work
  orders, logging material, quantity, responsible user, and date.
- **FR-017**: System MUST record audit trail for creation, update, logical
  deletion, status changes, and technician assignment.
- **FR-018**: Each audit record MUST include user, action, entity, entity ID,
  date, previous values, and new values.
- **FR-019**: Work orders MUST be searchable by status, priority, client,
  technician, and date range.
- **FR-020**: Clients MUST be searchable by name and tax ID.
- **FR-021**: Equipment MUST be searchable by serial number and client.
- **FR-022**: All collection endpoints MUST support pagination with page and
  per_page parameters.
- **FR-023**: All exposed functionality MUST be documented via OpenAPI/Swagger.

### Key Entities

- **User**: System actor with name, email, password (hashed), role
  (administrator, technician, viewer), and active status.
- **Client**: Business entity with business name, tax ID (unique), primary
  contact, email, phone, and address.
- **EquipmentType**: Classification category (e.g., Printer, Scanner, POS,
  Notebook).
- **Brand**: Manufacturer associated with an equipment type.
- **Model**: Product model associated with a brand and equipment type.
- **Equipment**: Physical asset installed at a client, with serial number
  (unique), model, location, and status.
- **WorkOrder**: Service request with client, equipment, description, priority,
  status, scheduled date, assigned technician, and completion date.
- **WorkOrderObservation**: Free-text note attached to a work order, with
  author, date, and text (immutable).
- **Material**: Supply item with unique code, description, and unit cost.
- **MaterialConsumption**: Record of material used in a work order, with
  material, quantity, user, and date.
- **AuditLog**: Immutable record of changes with user, action, entity, entity
  ID, date, previous values, and new values.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: A user can complete the login flow and receive valid tokens in
  under 2 seconds.
- **SC-002**: An administrator can register a new client with all required
  fields in under 3 minutes.
- **SC-003**: A technician can view, update status, and complete a work order
  in under 5 interactions.
- **SC-004**: All collection endpoints return paginated results within 500 ms
  under normal load.
- **SC-005**: Audit trail is generated automatically for every tracked action
  with zero manual intervention.
- **SC-006**: Duplicate tax IDs and serial numbers are consistently rejected
  with clear error messages.
- **SC-007**: Unauthorized access to any protected endpoint is consistently
  rejected with a 401 response.
- **SC-008**: A new developer can set up and run the application using only
  environment variables and `docker compose up` without manual database setup.

## Assumptions

- The system will be deployed as a REST API consumed by future integrations
  (mobile app, web interface) — no UI is included in this scope.
- Email and password are the sole authentication method; social login or SSO
  are not required in this version.
- The initial admin user (admin@example.com / Admin123!) will be created via
  seed data, not through a registration endpoint.
- Rate limiting and security headers are configured at the infrastructure level
  or as middleware, not per-endpoint.
- The audit log is append-only and never purged; no data retention policy is
  specified.
- The six work order states are exhaustive and closed — no custom or
  user-defined states are supported.
- Soft delete is required for main entities (users, clients, equipment, work
  orders, materials) per the project constitution.
- All sensitive values (database credentials, JWT secrets, etc.) are provided
  via environment variables.
