---

description: "Task list for Technician Geolocation implementation"

---

# Tasks: Technician Geolocation

**Input**: Design documents from `specs/002-technician-geolocation/`

**Prerequisites**: plan.md (required), spec.md (required for user stories), research.md, data-model.md, contracts/

**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story.

**Note**: This feature extends the existing Go API project. All Phase 1 (Setup) and Phase 2 (Foundational) infrastructure is already in place. Only new module tasks are listed.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (US1, US2)
- Include exact file paths in descriptions

## Path Conventions

- **Project root**: repository root
- **Go module root**: `cmd/api/`
- **Source structure**: `cmd/`, `internal/`, `pkg/`, `tests/`

---

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Project initialization — already complete from existing project. No tasks required.

✅ All infrastructure already exists: Go module, Gin, GORM, MySQL, Docker, middleware.

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Domain entity that MUST be complete before ANY user story can begin.

**CRITICAL**: No user story work can begin until this phase is complete.

### Domain Entity

- [x] T001 [P] Create TechnicianLocation entity in `internal/domain/technicianlocation/location.go`

### DTO

- [x] T002 [P] Create TechnicianLocation DTOs (create request, response) in `internal/dto/technicianlocation.go`

**Checkpoint**: Foundation ready — user story implementation can now begin.

---

## Phase 3: User Story 1 - Technician Reports Location (Priority: P1) 🎯 MVP

**Goal**: Technicians can report GPS coordinates. Only the authenticated technician's own records are accepted.

**Independent Test**: Authenticate as a technician, send valid coordinates, verify the record is stored with the correct user_id and timestamp. Verify invalid coordinates are rejected with 400. Verify viewer role is rejected with 403.

### Implementation for User Story 1

- [x] T003 [P] [US1] Create TechnicianLocation repository in `internal/repositories/technician_location_repository.go`
- [x] T004 [P] [US1] Create TechnicianLocation service in `internal/services/technician_location_service.go`
- [x] T005 [P] [US1] Create TechnicianLocation handler in `internal/handlers/technician_location_handler.go`
- [x] T006 [US1] Register report location route (POST /api/v1/technician-locations) in `cmd/api/routes.go`
- [x] T007 [US1] Wire dependencies in `cmd/api/main.go`

**Checkpoint**: User Story 1 complete — technicians can report GPS locations. Endpoint returns 201, 400, or 403.

---

## Phase 4: User Story 2 - Administrator Queries Technician Locations (Priority: P1)

**Goal**: Administrators can query all technician locations with optional filters. Technicians can query only their own records.

**Independent Test**: Authenticate as an admin, query all locations, verify paginated results. Filter by technician ID and date range. Authenticate as a technician, verify only own locations returned.

### Implementation for User Story 2

- [x] T008 [P] [US2] Add query parameters logic to TechnicianLocation handler (user_id, start_date, end_date, page, per_page) in `internal/handlers/technician_location_handler.go`
- [x] T009 [P] [US2] Add query method to TechnicianLocation repository in `internal/repositories/technician_location_repository.go`
- [x] T010 [US2] Add query method to TechnicianLocation service in `internal/services/technician_location_service.go`
- [x] T011 [US2] Register list location route (GET /api/v1/technician-locations) in `cmd/api/routes.go`

**Checkpoint**: User Story 2 complete — administrators can query all locations with filters. Technicians see only their own.

---

## Phase 5: Polish & Cross-Cutting Concerns

**Purpose**: Validation, tests, and documentation

- [x] T012 [P] Add coordinate validation (latitude -90 to 90, longitude -180 to 180) in handler/request level
- [x] T013 [P] Add Swagger/OpenAPI annotations to both handler endpoints in `internal/handlers/technician_location_handler.go`
- [x] T014 [P] Regenerate Swagger docs with `swag init`
- [x] T015 [P] Write integration tests in `tests/services/technician_location_service_test.go`
- [x] T016 Verify all tests pass with `go test ./... -count=1`
- [x] T017 Verify `docker compose up -d` builds and starts without errors
- [x] T018 Verify Swagger UI shows new endpoints at `/swagger/index.html`

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: Already complete
- **Foundational (Phase 2)**: T001-T002 — BLOCKS all user stories
- **User Stories (Phase 3-4)**: Both depend on Foundational
  - US2 depends on US1 (need location data to query)
- **Polish (Phase 5)**: Depends on US1 + US2 being complete

### User Story Dependencies

- **US1 Technician Reports Location**: Can start after Foundational
- **US2 Administrator Queries**: Depends on US1 (needs data to query)

### Within Each User Story

- Repository before service
- Service before handler
- Routes after handler
- Dependencies after services

### Parallel Opportunities

- T001 (entity) and T002 (DTO) can run in parallel
- T003 (repository), T004 (service), T005 (handler) can run in parallel within US1
- T008 (handler logic) and T009 (repository query) can run in parallel within US2

---

## Implementation Strategy

### MVP First (User Story 1 Only)

1. Complete Phase 2: Foundational (T001-T002)
2. Complete Phase 3: User Story 1 (T003-T007)
3. **STOP and VALIDATE**: Technician can report location

### Full Delivery

1. Phase 2: Foundational
2. Phase 3: US1 (MVP — technician reports location)
3. Phase 4: US2 (admin queries)
4. Phase 5: Polish & validation

---

## Notes

- All new code follows existing Clean Architecture patterns
- Reuse existing auth middleware and RBAC
- No new external dependencies required
- No database migrations (GORM AutoMigrate handles new table)
- Coordinate validation at handler level to prevent invalid data
