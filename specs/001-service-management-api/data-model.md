# Data Model: Service Management API

## Entities

### User

| Field | Type | Constraints |
|-------|------|-------------|
| ID | uint64 | PK, auto-increment |
| Name | string | NOT NULL |
| Email | string | NOT NULL, UNIQUE, INDEX |
| Password | string | NOT NULL (bcrypt hash) |
| Role | enum | NOT NULL (administrator, technician, viewer) |
| Active | bool | NOT NULL, DEFAULT true |
| CreatedAt | datetime | NOT NULL |
| UpdatedAt | datetime | NOT NULL |
| DeletedAt | datetime | NULLABLE (soft delete) |

**Relationships**:
- Has many WorkOrders (as assigned technician)
- Has many AuditLogs (as actor)
- Has many WorkOrderObservations (as author)
- Has many MaterialConsumptions (as responsible user)

---

### Client

| Field | Type | Constraints |
|-------|------|-------------|
| ID | uint64 | PK, auto-increment |
| BusinessName | string | NOT NULL |
| TaxID | string | NOT NULL, UNIQUE, INDEX |
| PrimaryContact | string | NOT NULL |
| Email | string | NOT NULL |
| Phone | string | NOT NULL |
| Address | text | NOT NULL |
| CreatedAt | datetime | NOT NULL |
| UpdatedAt | datetime | NOT NULL |
| DeletedAt | datetime | NULLABLE (soft delete) |

**Relationships**:
- Has many Equipment
- Has many WorkOrders

---

### EquipmentType

| Field | Type | Constraints |
|-------|------|-------------|
| ID | uint64 | PK, auto-increment |
| Name | string | NOT NULL, UNIQUE |
| CreatedAt | datetime | NOT NULL |
| UpdatedAt | datetime | NOT NULL |

**Relationships**:
- Has many Brands

---

### Brand

| Field | Type | Constraints |
|-------|------|-------------|
| ID | uint64 | PK, auto-increment |
| Name | string | NOT NULL |
| EquipmentTypeID | uint64 | FK → EquipmentType.ID, NOT NULL |
| CreatedAt | datetime | NOT NULL |
| UpdatedAt | datetime | NOT NULL |

**Unique Constraint**: (Name, EquipmentTypeID)

**Relationships**:
- Belongs to EquipmentType
- Has many Models

---

### Model

| Field | Type | Constraints |
|-------|------|-------------|
| ID | uint64 | PK, auto-increment |
| Name | string | NOT NULL |
| BrandID | uint64 | FK → Brand.ID, NOT NULL |
| EquipmentTypeID | uint64 | FK → EquipmentType.ID, NOT NULL |
| CreatedAt | datetime | NOT NULL |
| UpdatedAt | datetime | NOT NULL |

**Unique Constraint**: (Name, BrandID)

**Relationships**:
- Belongs to Brand
- Belongs to EquipmentType
- Has many Equipment

---

### Equipment

| Field | Type | Constraints |
|-------|------|-------------|
| ID | uint64 | PK, auto-increment |
| ClientID | uint64 | FK → Client.ID, NOT NULL, INDEX |
| ModelID | uint64 | FK → Model.ID, NOT NULL |
| SerialNumber | string | NOT NULL, UNIQUE, INDEX |
| Location | string | NOT NULL |
| Status | string | NOT NULL |
| CreatedAt | datetime | NOT NULL |
| UpdatedAt | datetime | NOT NULL |
| DeletedAt | datetime | NULLABLE (soft delete) |

**Relationships**:
- Belongs to Client
- Belongs to Model
- Has many WorkOrders

---

### WorkOrder

| Field | Type | Constraints |
|-------|------|-------------|
| ID | uint64 | PK, auto-increment |
| ClientID | uint64 | FK → Client.ID, NOT NULL, INDEX |
| EquipmentID | uint64 | FK → Equipment.ID, NOT NULL |
| Description | text | NOT NULL |
| Priority | enum | NOT NULL (low, medium, high, critical) |
| Status | enum | NOT NULL, DEFAULT pending |
| ScheduledDate | date | NOT NULL |
| CompletedDate | datetime | NULLABLE |
| TechnicianID | uint64 | FK → User.ID, NULLABLE, INDEX |
| AssignedByID | uint64 | FK → User.ID, NULLABLE |
| AssignedAt | datetime | NULLABLE |
| CreatedAt | datetime | NOT NULL |
| UpdatedAt | datetime | NOT NULL |
| DeletedAt | datetime | NULLABLE (soft delete) |

**Status Values**: pending, assigned, in_progress, waiting_parts, completed, cancelled

**Relationships**:
- Belongs to Client
- Belongs to Equipment
- Belongs to User (technician)
- Has many WorkOrderObservations
- Has many MaterialConsumptions
- Has many AuditLogs

---

### WorkOrderObservation

| Field | Type | Constraints |
|-------|------|-------------|
| ID | uint64 | PK, auto-increment |
| WorkOrderID | uint64 | FK → WorkOrder.ID, NOT NULL, INDEX |
| AuthorID | uint64 | FK → User.ID, NOT NULL |
| Text | text | NOT NULL |
| CreatedAt | datetime | NOT NULL (immutable) |

**Note**: No UpdatedAt or DeletedAt — observations are immutable per FR-014.

**Relationships**:
- Belongs to WorkOrder
- Belongs to User (author)

---

### Material

| Field | Type | Constraints |
|-------|------|-------------|
| ID | uint64 | PK, auto-increment |
| Code | string | NOT NULL, UNIQUE, INDEX |
| Description | string | NOT NULL |
| UnitCost | decimal(10,2) | NOT NULL |
| CreatedAt | datetime | NOT NULL |
| UpdatedAt | datetime | NOT NULL |
| DeletedAt | datetime | NULLABLE (soft delete) |

**Relationships**:
- Has many MaterialConsumptions

---

### MaterialConsumption

| Field | Type | Constraints |
|-------|------|-------------|
| ID | uint64 | PK, auto-increment |
| WorkOrderID | uint64 | FK → WorkOrder.ID, NOT NULL, INDEX |
| MaterialID | uint64 | FK → Material.ID, NOT NULL |
| Quantity | decimal(10,2) | NOT NULL |
| UserID | uint64 | FK → User.ID, NOT NULL |
| CreatedAt | datetime | NOT NULL |

**Relationships**:
- Belongs to WorkOrder
- Belongs to Material
- Belongs to User (responsible)

---

### AuditLog

| Field | Type | Constraints |
|-------|------|-------------|
| ID | uint64 | PK, auto-increment |
| UserID | uint64 | FK → User.ID, NULLABLE |
| Action | string | NOT NULL (INSERT, UPDATE, DELETE, STATUS_CHANGE, ASSIGNMENT) |
| Entity | string | NOT NULL |
| EntityID | uint64 | NOT NULL |
| OldValues | json | NULLABLE |
| NewValues | json | NULLABLE |
| CreatedAt | datetime | NOT NULL (immutable) |

**Note**: Immutable, never deleted. No UpdatedAt or DeletedAt.

**Indexes**: (Entity, EntityID), (UserID), (CreatedAt)

**Relationships**:
- Belongs to User (actor)

---

## Entity Relationship Diagram (text)

```
User ──┬── WorkOrder (technician)
       ├── AuditLog (actor)
       ├── WorkOrderObservation (author)
       └── MaterialConsumption (user)

Client ──┬── Equipment
         └── WorkOrder

EquipmentType ──┬── Brand ──┬── Model ──┬── Equipment ──┬── WorkOrder
                │           │           │               │
                └───────────┘           │               └── WorkOrderObservation
                                        │               └── MaterialConsumption
                                        └── Material ───┘

WorkOrder ──┬── WorkOrderObservation
            ├── MaterialConsumption
            └── AuditLog
```

## Validation Rules

| Entity | Field | Rule |
|--------|-------|------|
| User | Email | Unique |
| User | Role | Must be administrator, technician, or viewer |
| User | Password | Min 8 chars, uppercase, lowercase, number |
| Client | TaxID | Unique |
| Equipment | SerialNumber | Unique |
| Material | Code | Unique |
| WorkOrder | Status | Must be one of 6 allowed values |
| WorkOrder | Priority | Must be low, medium, high, or critical |

## Soft Delete Policy

| Entity | Soft Delete |
|--------|-------------|
| User | ✅ Sí |
| Client | ✅ Sí |
| Equipment | ✅ Sí |
| Material | ✅ Sí |
| WorkOrder | ✅ Sí |
| EquipmentType | ❌ No |
| Brand | ❌ No |
| Model | ❌ No |
| WorkOrderObservation | ❌ No (immutable) |
| MaterialConsumption | ❌ No |
| AuditLog | ❌ No (immutable, never deleted) |
