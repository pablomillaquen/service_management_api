package client

import (
	"time"

	"gorm.io/gorm"
)

type Client struct {
	ID             uint64         `gorm:"primaryKey;autoIncrement" json:"id"`
	BusinessName   string         `gorm:"size:255;not null" json:"business_name"`
	TaxID          string         `gorm:"size:50;not null;uniqueIndex" json:"tax_id"`
	PrimaryContact string         `gorm:"size:255;not null" json:"primary_contact"`
	Email          string         `gorm:"size:255;not null" json:"email"`
	Phone          string         `gorm:"size:50;not null" json:"phone"`
	Address        string         `gorm:"type:text;not null" json:"address"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Client) TableName() string {
	return "clients"
}
