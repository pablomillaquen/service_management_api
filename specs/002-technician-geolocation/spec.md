# Feature Specification: Technician Geolocation

**Feature Branch**: `002-technician-geolocation`

**Created**: 2026-06-08

**Status**: Draft

**Input**: User description: "Technician Geolocation — Allow technicians to report GPS coordinates. Store latitude, longitude, and timestamp. Administrators can query technician locations. Technicians can only create their own location records."

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Technician Reports Location (Priority: P1)

As a Technician, I want to report my current GPS coordinates so that administrators can track my location during service visits.

**Why this priority**: Location reporting is the core value of this feature — without it, there is no data to query.

**Independent Test**: Can be fully tested by authenticating as a technician, sending GPS coordinates via the API, and verifying the record is stored with the correct technician, latitude, longitude, and timestamp.

**Acceptance Scenarios**:

1. **Given** I am an authenticated technician with an active session, **When** I send my current GPS coordinates (latitude, longitude), **Then** the system stores the location with my user ID and the current timestamp.
2. **Given** an authenticated administrator, **When** I attempt to report a location for a technician other than myself, **Then** the system rejects the request with a 403 Forbidden.
3. **Given** an authenticated viewer, **When** I attempt to report a location, **Then** the system rejects the request with a 403 Forbidden.

---

### User Story 2 - Administrator Queries Technician Locations (Priority: P1)

As an Administrator, I want to query the reported locations of technicians so that I can monitor service visit progress and respond to issues.

**Why this priority**: Querying is equally critical — storing locations without the ability to retrieve them provides no value.

**Independent Test**: Can be fully tested by having multiple technicians report locations, then querying as an administrator with optional date range and technician filters.

**Acceptance Scenarios**:

1. **Given** multiple technicians have reported locations, **When** I query all locations, **Then** I receive a paginated list ordered by most recent first.
2. **Given** locations exist for a specific technician, **When** I filter by that technician's ID, **Then** I receive only that technician's locations.
3. **Given** locations exist across multiple days, **When** I filter by a date range, **Then** I receive only locations within that range.
4. **Given** I am a technician, **When** I attempt to query another technician's locations, **Then** the system rejects the request with a 403 Forbidden.
5. **Given** I am a technician, **When** I query my own locations, **Then** I receive only my location records.

---

### Edge Cases

- What happens when latitude or longitude values are out of valid range (-90 to 90 for latitude, -180 to 180 for longitude)?
- How does the system handle concurrent location reports from the same technician?
- What happens when a deactivated technician attempts to report a location?
- How are location records handled when a technician user is soft-deleted?
- What happens when a query returns an empty result set?

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: System MUST allow authenticated technicians to report their GPS location with latitude, longitude, and automatic timestamp.
- **FR-002**: System MUST validate latitude (-90 to 90) and longitude (-180 to 180) on every create request.
- **FR-003**: Technicians MUST only be able to create location records for themselves (user_id matches their JWT identity).
- **FR-004**: Viewers MUST NOT be able to create location records.
- **FR-005**: Administrators MUST be able to query all technician locations with optional filters: technician ID, date range.
- **FR-006**: Technicians MUST be able to query only their own location records.
- **FR-007**: All location collection endpoints MUST support pagination with page and per_page parameters.
- **FR-008**: Location records MUST include user_id (the technician), latitude, longitude, and created_at timestamp.
- **FR-009**: Location records MUST be ordered by most recent first by default.

### Key Entities

- **TechnicianLocation**: GPS location report with user_id (FK to users), latitude, longitude, and created_at timestamp.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: A technician can report GPS coordinates in under 1 second (API round-trip).
- **SC-002**: An administrator can retrieve technician locations with filters and see results within 500 ms under normal load.
- **SC-003**: Invalid coordinates (latitude out of range, longitude out of range) are consistently rejected with clear error messages.
- **SC-004**: Unauthorized access (viewer creating, technician querying others) is consistently rejected with 403.

## Assumptions

- Location data is write-once, read-many — no update or delete endpoints are required.
- No real-time streaming or WebSocket support is needed for v1.
- Coordinates are reported by the technician manually or via client app — server does not poll or request location.
- Timestamps are server-generated, not client-supplied, to prevent tampering.
- The soft-delete of a technician user does not cascade-delete their location history.
- No geofencing, reverse geocoding, or map visualization is included in this scope.
