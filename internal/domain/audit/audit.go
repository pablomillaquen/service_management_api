package audit

import (
	"time"
)

type Action string

const (
	ActionInsert       Action = "INSERT"
	ActionUpdate       Action = "UPDATE"
	ActionDelete       Action = "DELETE"
	ActionStatusChange Action = "STATUS_CHANGE"
	ActionAssignment   Action = "ASSIGNMENT"
)

type AuditLog struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    *uint64   `gorm:"index" json:"user_id,omitempty"`
	Action    Action    `gorm:"size:50;not null" json:"action"`
	Entity    string    `gorm:"size:100;not null;index" json:"entity"`
	EntityID  uint64    `gorm:"not null;index" json:"entity_id"`
	OldValues *string   `gorm:"type:json" json:"old_values,omitempty"`
	NewValues *string   `gorm:"type:json" json:"new_values,omitempty"`
	CreatedAt time.Time `gorm:"index" json:"created_at"`
}

func (AuditLog) TableName() string {
	return "audit_logs"
}
