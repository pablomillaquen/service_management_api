---

description: "Task list for Service Management API implementation"
---

# Tasks: Service Management API

**Input**: Design documents from `specs/001-service-management-api/`

**Prerequisites**: plan.md (required), spec.md (required for user stories),
research.md, data-model.md, contracts/

## Organization

Tasks are grouped by user story to enable independent implementation and
testing of each story.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (US1, US2, US3, US4, US5)
- Include exact file paths in descriptions

## Path Conventions

- **Project root**: repository root
- **Go module root**: `cmd/api/`
- **Source structure**: `cmd/`, `internal/`, `pkg/`, `migrations/`, `tests/`

---

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Project initialization and runtime environment

- [ ] T001 Initialize Go module at repository root
- [ ] T002 [P] Create project folder structure per Clean Architecture (cmd/,
      internal/, pkg/, migrations/, tests/, configs/, docs/)
- [ ] T003 Configure Gin framework with middleware stack in cmd/api/main.go
- [ ] T004 [P] Configure environment variable loading in configs/config.go
- [ ] T005 [P] Configure structured logging in pkg/logger/logger.go
- [ ] T006 Create application startup entrypoint in cmd/api/main.go
- [ ] T007 [P] Configure Dockerfile at repository root
- [ ] T008 [P] Configure docker-compose.yml at repository root
- [ ] T009 Create .env.example at repository root

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Database, domain entities, authentication, and authorization —
MUST be complete before ANY user story can begin

**CRITICAL**: No user story work can begin until this phase is complete

### Database Foundation

- [ ] T010 [P] Configure MySQL connection with pooling in pkg/database/mysql.go
- [ ] T011 [P] Create migration framework in pkg/database/migration.go
- [ ] T012 [P] Create seed framework in pkg/database/seed.go
- [ ] T013 Configure health check endpoint in internal/handlers/health.go
- [ ] T014 Create base repository with common CRUD in internal/repositories/base.go

### Domain Entities

- [ ] T015 [P] Create User entity in internal/domain/user/user.go
- [ ] T016 [P] Create Client entity in internal/domain/client/client.go
- [ ] T017 [P] Create EquipmentType entity in internal/domain/equipment/type.go
- [ ] T018 [P] Create Brand entity in internal/domain/equipment/brand.go
- [ ] T019 [P] Create EquipmentModel entity in internal/domain/equipment/model.go
- [ ] T020 [P] Create Equipment entity in internal/domain/equipment/equipment.go
- [ ] T021 [P] Create Material entity in internal/domain/material/material.go
- [ ] T022 [P] Create WorkOrder entity in internal/domain/workorder/workorder.go
- [ ] T023 [P] Create WorkOrderNote entity in internal/domain/workorder/note.go
- [ ] T024 [P] Create WorkOrderMaterial entity in internal/domain/workorder/material.go
- [ ] T025 [P] Create AuditLog entity in internal/domain/audit/audit.go
- [ ] T026 [P] Create RefreshToken entity in internal/domain/user/refresh_token.go

### Authentication & Authorization

- [ ] T027 [P] Create unified API response format in pkg/response/response.go
- [ ] T028 [P] Implement JWT token generation in internal/auth/jwt.go
- [ ] T029 [P] Implement JWT token validation in internal/auth/jwt.go
- [ ] T030 [P] Implement bcrypt password hashing in internal/auth/password.go
- [ ] T031 Create User repository in internal/repositories/user_repository.go
- [ ] T032 Create User service with auth logic in internal/services/user_service.go
- [ ] T033 Create auth DTOs (login, refresh, change password) in internal/dto/auth.go
- [ ] T034 [P] Create authentication middleware in internal/middleware/auth.go
- [ ] T035 [P] Create RBAC middleware in internal/middleware/rbac.go
- [ ] T036 Create auth handlers (login, refresh, logout, change password) in
      internal/handlers/auth_handler.go
- [ ] T037 Register auth routes and middleware in cmd/api/routes.go

**Checkpoint**: Foundation ready — user story implementation can now begin in
parallel. Application starts, migrates DB, seeds initial schema, and
authenticates users.

---

## Phase 3: User Story 1 - Authentication & User Administration (Priority: P1) MVP

**Goal**: Administrators can manage user accounts and roles. Users can
authenticate and maintain sessions securely.

**Independent Test**: Create an admin user, log in, receive tokens, verify
protected endpoint access, refresh token, change password, and log out.

### Implementation for User Story 1

- [ ] T038 [P] [US1] Create user DTOs (create, update, list, response) in
      internal/dto/user.go
- [ ] T039 [P] [US1] Create User handler in internal/handlers/user_handler.go
- [ ] T040 [US1] Register user management routes (GET, GET/:id, POST, PUT,
      DELETE /api/v1/users) in cmd/api/routes.go
- [ ] T041 [US1] Implement user validations (unique email, valid role) in
      internal/validators/user_validator.go

**Checkpoint**: User Story 1 complete — admin can create/manage users, users
can authenticate and be authorized by role.

---

## Phase 4: User Story 2 - Client & Equipment Management (Priority: P1) MVP

**Goal**: Administrators can register clients, define equipment catalog
(types, brands, models), and track equipment inventory per client.

**Independent Test**: Create equipment types, brands, models; register a
client; register equipment for that client; verify duplicate detection.

### Implementation for User Story 2

- [ ] T042 [P] [US2] Create catalog DTOs (type, brand, model) in
      internal/dto/catalog.go
- [ ] T043 [P] [US2] Create client DTOs in internal/dto/client.go
- [ ] T044 [P] [US2] Create EquipmentType repository in
      internal/repositories/equipment_type_repository.go
- [ ] T045 [P] [US2] Create Brand repository in
      internal/repositories/brand_repository.go
- [ ] T046 [P] [US2] Create EquipmentModel repository in
      internal/repositories/equipment_model_repository.go
- [ ] T047 [P] [US2] Create Client repository in
      internal/repositories/client_repository.go
- [ ] T048 [P] [US2] Create Equipment repository in
      internal/repositories/equipment_repository.go
- [ ] T049 [P] [US2] Create EquipmentType service in
      internal/services/equipment_type_service.go
- [ ] T050 [P] [US2] Create Brand service in
      internal/services/brand_service.go
- [ ] T051 [P] [US2] Create EquipmentModel service in
      internal/services/equipment_model_service.go
- [ ] T052 [P] [US2] Create Client service in
      internal/services/client_service.go
- [ ] T053 [P] [US2] Create Equipment service in
      internal/services/equipment_service.go
- [ ] T054 [P] [US2] Create EquipmentType handler in
      internal/handlers/equipment_type_handler.go
- [ ] T055 [P] [US2] Create Brand handler in
      internal/handlers/brand_handler.go
- [ ] T056 [P] [US2] Create EquipmentModel handler in
      internal/handlers/equipment_model_handler.go
- [ ] T057 [P] [US2] Create Client handler in
      internal/handlers/client_handler.go
- [ ] T058 [P] [US2] Create Equipment handler in
      internal/handlers/equipment_handler.go
- [ ] T059 [US2] Register catalog + client + equipment routes in
      cmd/api/routes.go
- [ ] T060 [US2] Implement catalog validations (unique per type/brand) in
      internal/validators/catalog_validator.go
- [ ] T061 [US2] Implement client validations (unique tax ID) in
      internal/validators/client_validator.go
- [ ] T062 [US2] Implement equipment validations (unique serial number) in
      internal/validators/equipment_validator.go

**Checkpoint**: User Story 2 complete — clients, catalog, and equipment fully
manageable with validation and pagination.

---

## Phase 5: User Story 3 - Work Order Lifecycle (Priority: P1) MVP

**Goal**: Administrators create and assign work orders; technicians update
status through the lifecycle, add observations, and complete orders.

**Independent Test**: Create work order, assign technician, change status
through all six states, add notes, record material usage, and complete.

### Implementation for User Story 3

- [ ] T063 [P] [US3] Create work order DTOs in internal/dto/workorder.go
- [ ] T064 [P] [US3] Create WorkOrder repository in
      internal/repositories/workorder_repository.go
- [ ] T065 [P] [US3] Create WorkOrderNote repository in
      internal/repositories/workorder_note_repository.go
- [ ] T066 [P] [US3] Create WorkOrderMaterial repository in
      internal/repositories/workorder_material_repository.go
- [ ] T067 [P] [US3] Create WorkOrder service in
      internal/services/workorder_service.go
- [ ] T068 [P] [US3] Create WorkOrder handler in
      internal/handlers/workorder_handler.go
- [ ] T069 [US3] Implement status transition validation (only 6 allowed
      statuses) in internal/validators/workorder_validator.go
- [ ] T070 [US3] Implement work order assign endpoint (POST
      /api/v1/work-orders/:id/assign) in internal/handlers/workorder_handler.go
- [ ] T071 [US3] Implement work order status change (PATCH
      /api/v1/work-orders/:id/status) in internal/handlers/workorder_handler.go
- [ ] T072 [US3] Implement add notes endpoint (POST
      /api/v1/work-orders/:id/notes) in internal/handlers/workorder_handler.go
- [ ] T073 [US3] Implement register materials endpoint (POST
      /api/v1/work-orders/:id/materials) in internal/handlers/workorder_handler.go
- [ ] T074 [US3] Register all work order routes in cmd/api/routes.go

**Checkpoint**: User Story 3 complete — full work order lifecycle operational
with assignment, status transitions, notes, and material tracking.

---

## Phase 6: User Story 4 - Materials & Consumption Tracking (Priority: P2)

**Goal**: Administrators manage material catalog; technicians record material
usage against work orders.

**Independent Test**: Create materials in catalog, then record consumption
against a work order.

### Implementation for User Story 4

- [ ] T075 [P] [US4] Create material DTOs in internal/dto/material.go
- [ ] T076 [P] [US4] Create Material repository in
      internal/repositories/material_repository.go
- [ ] T077 [P] [US4] Create Material service in
      internal/services/material_service.go
- [ ] T078 [P] [US4] Create Material handler in
      internal/handlers/material_handler.go
- [ ] T079 [US4] Register material routes in cmd/api/routes.go
- [ ] T080 [US4] Implement material validations (unique code) in
      internal/validators/material_validator.go

**Checkpoint**: User Story 4 complete — materials catalog and consumption
tracking operational.

---

## Phase 7: User Story 5 - Search, Audit & Documentation (Priority: P3)

**Goal**: Users search across records, view audit trail, and access API
documentation.

**Independent Test**: Perform filtered searches, verify audit records
generated on modification, access Swagger UI.

### Implementation for User Story 5

- [ ] T081 [P] [US5] Create audit DTOs in internal/dto/audit.go
- [ ] T082 [P] [US5] Create AuditLog repository in
      internal/repositories/audit_repository.go
- [ ] T083 [P] [US5] Create Audit service in internal/services/audit_service.go
- [ ] T084 [P] [US5] Create audit middleware in
      internal/middleware/audit.go
- [ ] T085 [P] [US5] Integrate audit middleware with all business handlers
- [ ] T086 [P] [US5] Create audit handler (list/search) in
      internal/handlers/audit_handler.go
- [ ] T087 [US5] Register audit routes in cmd/api/routes.go
- [ ] T088 [US5] Add search/filter parameters to all list endpoints (status,
      priority, client, technician, date range, serial number, tax ID)

**Checkpoint**: User Story 5 complete — search, audit trail, and documentation
operational.

---

## Phase 8: Polish & Cross-Cutting Concerns

**Purpose**: Security hardening, API documentation, testing, seed data, and
production validation

### Security Hardening

- [ ] T089 [P] Implement security headers middleware in
      internal/middleware/security.go
- [ ] T090 [P] Implement CORS middleware in internal/middleware/cors.go
- [ ] T091 [P] Implement rate limiting middleware in
      internal/middleware/ratelimit.go
- [ ] T092 [P] Implement panic recovery middleware in
      internal/middleware/recovery.go
- [ ] T093 [P] Implement request validation and sanitization in
      internal/middleware/validation.go
- [ ] T094 [P] Register all security middleware in cmd/api/routes.go

### API Documentation

- [ ] T095 [P] Configure Swagger/OpenAPI generation in cmd/api/main.go
- [ ] T096 [P] Annotate all handler endpoints with Swagger comments
- [ ] T097 [P] Generate and verify Swagger UI at /swagger/index.html

### Seed Data

- [ ] T098 Create seed script for initial admin user (admin@example.com /
      Admin123!) in pkg/database/seed.go

### Testing

- [ ] T099 [P] Write authentication unit tests in tests/auth_test.go
- [ ] T100 [P] Write RBAC middleware tests in tests/rbac_test.go
- [ ] T101 [P] Write User service unit tests in tests/services/user_test.go
- [ ] T102 [P] Write Client service + repository tests in
      tests/services/client_test.go
- [ ] T103 [P] Write Equipment service + repository tests in
      tests/services/equipment_test.go
- [ ] T104 [P] Write WorkOrder service + lifecycle tests in
      tests/services/workorder_test.go
- [ ] T105 [P] Write Material service tests in tests/services/material_test.go
- [ ] T106 [P] Write Audit service tests in tests/services/audit_test.go
- [ ] T107 [P] Write repository tests for all CRUD operations in
      tests/repositories/

### Production Validation

- [ ] T108 Verify `docker compose up` completes without errors
- [ ] T109 Verify database is created and migrations execute automatically
- [ ] T110 Verify seed admin user can log in
- [ ] T111 Verify Swagger UI is accessible at /swagger/index.html
- [ ] T112 Verify all unit tests pass with `go test ./...`
- [ ] T113 Verify coverage meets 80% threshold

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies — can start immediately
- **Foundational (Phase 2)**: Depends on Setup — BLOCKS all user stories
- **User Stories (Phases 3-7)**: All depend on Foundational completion
  - User stories can proceed sequentially in priority order (P1 → P2 → P3)
- **Polish (Phase 8)**: Depends on all desired user stories being complete

### User Story Dependencies

- **US1 Auth & User Admin (Phase 3)**: Can start after Foundational
- **US2 Client & Equipment (Phase 4)**: Can start after Foundational
- **US3 Work Orders (Phase 5)**: Can start after Foundational; depends on
  US2 for client/equipment references
- **US4 Materials (Phase 6)**: Can start after Foundational (independent)
- **US5 Audit & Search (Phase 7)**: Depends on US3 for audit events

### Within Each User Story

- DTOs before handlers
- Repositories before services
- Services before handlers
- Routes after handlers

### Parallel Opportunities

- All Phase 1 [P] tasks can run in parallel
- All Phase 2 [P] tasks (domain entities) can run in parallel
- Once Foundational completes, US1, US2, US4 can start in parallel (US3
  depends on US2; US5 depends on US3)
- All [P] tasks within a phase can run in parallel
- All tests marked [P] can run in parallel

---

## Implementation Strategy

### MVP First (User Stories 1-3, Priorities P1)

1. Complete Phase 1: Setup
2. Complete Phase 2: Foundational (CRITICAL — blocks all stories)
3. Complete Phase 3: User Story 1 (Auth & User Admin)
4. Complete Phase 4: User Story 2 (Client & Equipment)
5. Complete Phase 5: User Story 3 (Work Order Lifecycle)
6. **STOP and VALIDATE**: Full work order lifecycle working end-to-end
7. Deploy/demo if ready

### Incremental Delivery

1. Setup + Foundational → Foundation ready
2. Add US1 (Auth & User Admin) → Test independently → Deploy
3. Add US2 (Client & Equipment) → Test independently → Deploy
4. Add US3 (Work Order Lifecycle) → Test independently → Deploy/Demo (MVP!)
5. Add US4 (Materials) → Test independently → Deploy
6. Add US5 (Audit & Search) → Test independently → Deploy
7. Add Polish (Security, Docs, Tests) → Final validation

### Parallel Team Strategy

With multiple developers:

1. Complete Setup + Foundational together
2. Once Foundational is done:
   - Developer A: US1 (Auth & User Admin) + US5 (Audit, depends on US3)
   - Developer B: US2 (Client & Equipment)
   - Developer C: US4 (Materials, independent)
3. After US2: Developer B continues to US3 (Work Orders)
4. After US3 + US1: Developer A adds US5

---

## Notes

- [P] tasks = different files, no dependencies
- [Story] label maps task to specific user story for traceability
- Each user story is independently completable and testable
- Verify tests fail before implementing (if TDD approach)
- Commit after each task or logical group
- Stop at any checkpoint to validate story independently
- Avoid: vague tasks, same file conflicts, cross-story dependencies that
  break independence
