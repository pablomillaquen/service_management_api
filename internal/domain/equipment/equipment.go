package equipment

import (
	"time"

	"gorm.io/gorm"
)

type EquipmentType struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"size:100;not null;uniqueIndex" json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Brand struct {
	ID              uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name            string    `gorm:"size:100;not null" json:"name"`
	EquipmentTypeID uint64    `gorm:"not null;index" json:"equipment_type_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type EquipmentModel struct {
	ID              uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name            string    `gorm:"size:100;not null" json:"name"`
	BrandID         uint64    `gorm:"not null;index" json:"brand_id"`
	EquipmentTypeID uint64    `gorm:"not null;index" json:"equipment_type_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type Equipment struct {
	ID           uint64         `gorm:"primaryKey;autoIncrement" json:"id"`
	ClientID     uint64         `gorm:"not null;index" json:"client_id"`
	ModelID      uint64         `gorm:"not null;index" json:"model_id"`
	SerialNumber string         `gorm:"size:100;not null;uniqueIndex" json:"serial_number"`
	Location     string         `gorm:"size:255;not null" json:"location"`
	Status       string         `gorm:"size:50;not null" json:"status"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

func (EquipmentType) TableName() string  { return "equipment_types" }
func (Brand) TableName() string          { return "brands" }
func (EquipmentModel) TableName() string { return "equipment_models" }
func (Equipment) TableName() string      { return "equipment" }
