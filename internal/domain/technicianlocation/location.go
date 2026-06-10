package technicianlocation

import "time"

type TechnicianLocation struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint64    `gorm:"not null;index:idx_user_id_created_at,priority:1" json:"user_id"`
	Latitude  float64   `gorm:"not null" json:"latitude"`
	Longitude float64   `gorm:"not null" json:"longitude"`
	CreatedAt time.Time `gorm:"not null;index:idx_user_id_created_at,priority:2;index:idx_created_at,sort:desc" json:"created_at"`
}

func (TechnicianLocation) TableName() string {
	return "technician_locations"
}
