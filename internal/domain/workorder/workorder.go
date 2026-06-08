package workorder

import (
	"time"

	"gorm.io/gorm"
)

type Status string

const (
	StatusPending      Status = "pending"
	StatusAssigned     Status = "assigned"
	StatusInProgress   Status = "in_progress"
	StatusWaitingParts Status = "waiting_parts"
	StatusCompleted    Status = "completed"
	StatusCancelled    Status = "cancelled"
)

var ValidStatuses = map[Status]bool{
	StatusPending:      true,
	StatusAssigned:     true,
	StatusInProgress:   true,
	StatusWaitingParts: true,
	StatusCompleted:    true,
	StatusCancelled:    true,
}

func (s Status) IsValid() bool {
	return ValidStatuses[s]
}

var AllowedTransitions = map[Status][]Status{
	StatusPending:      {StatusAssigned, StatusCancelled},
	StatusAssigned:     {StatusInProgress, StatusCancelled},
	StatusInProgress:   {StatusWaitingParts, StatusCompleted, StatusCancelled},
	StatusWaitingParts: {StatusInProgress, StatusCancelled},
	StatusCompleted:    {},
	StatusCancelled:    {},
}

func IsValidTransition(from, to Status) bool {
	allowed, ok := AllowedTransitions[from]
	if !ok {
		return false
	}
	for _, s := range allowed {
		if s == to {
			return true
		}
	}
	return false
}

type Priority string

const (
	PriorityLow     Priority = "low"
	PriorityMedium  Priority = "medium"
	PriorityHigh    Priority = "high"
	PriorityCritical Priority = "critical"
)

func (p Priority) IsValid() bool {
	switch p {
	case PriorityLow, PriorityMedium, PriorityHigh, PriorityCritical:
		return true
	default:
		return false
	}
}

type WorkOrder struct {
	ID            uint64         `gorm:"primaryKey;autoIncrement" json:"id"`
	ClientID      uint64         `gorm:"not null;index" json:"client_id"`
	EquipmentID   uint64         `gorm:"not null;index" json:"equipment_id"`
	Description   string         `gorm:"type:text;not null" json:"description"`
	Priority      Priority       `gorm:"size:20;not null" json:"priority"`
	Status        Status         `gorm:"size:20;not null;default:pending;index" json:"status"`
	ScheduledDate string         `gorm:"size:10;not null" json:"scheduled_date"`
	CompletedDate *time.Time     `json:"completed_date,omitempty"`
	TechnicianID  *uint64        `gorm:"index" json:"technician_id,omitempty"`
	AssignedByID  *uint64        `json:"assigned_by_id,omitempty"`
	AssignedAt    *time.Time     `json:"assigned_at,omitempty"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

type WorkOrderNote struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	WorkOrderID uint64    `gorm:"not null;index" json:"work_order_id"`
	AuthorID    uint64    `gorm:"not null;index" json:"author_id"`
	Text        string    `gorm:"type:text;not null" json:"text"`
	CreatedAt   time.Time `json:"created_at"`
}

type WorkOrderMaterial struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	WorkOrderID uint64    `gorm:"not null;index" json:"work_order_id"`
	MaterialID  uint64    `gorm:"not null;index" json:"material_id"`
	Quantity    float64   `gorm:"type:decimal(10,2);not null" json:"quantity"`
	UserID      uint64    `gorm:"not null;index" json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
}

func (WorkOrder) TableName() string         { return "work_orders" }
func (WorkOrderNote) TableName() string     { return "work_order_notes" }
func (WorkOrderMaterial) TableName() string { return "work_order_materials" }
