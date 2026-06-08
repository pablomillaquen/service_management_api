package material

import (
	"time"

	"gorm.io/gorm"
)

type Material struct {
	ID          uint64         `gorm:"primaryKey;autoIncrement" json:"id"`
	Code        string         `gorm:"size:50;not null;uniqueIndex" json:"code"`
	Description string         `gorm:"size:255;not null" json:"description"`
	UnitCost    float64        `gorm:"type:decimal(10,2);not null" json:"unit_cost"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Material) TableName() string {
	return "materials"
}
