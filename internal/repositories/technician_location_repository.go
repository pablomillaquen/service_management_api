package repositories

import (
	"time"

	"github.com/pablomillaquen/speckit_golang_api/internal/domain/technicianlocation"
	"gorm.io/gorm"
)

type TechnicianLocationRepository struct {
	db *gorm.DB
}

func NewTechnicianLocationRepository(db *gorm.DB) *TechnicianLocationRepository {
	return &TechnicianLocationRepository{db: db}
}

func (r *TechnicianLocationRepository) Create(loc *technicianlocation.TechnicianLocation) error {
	return r.db.Create(loc).Error
}

func (r *TechnicianLocationRepository) FindAll(page, perPage int, userID *uint64, startDate, endDate *time.Time) ([]technicianlocation.TechnicianLocation, int64, error) {
	var locations []technicianlocation.TechnicianLocation
	var total int64
	query := r.db.Model(&technicianlocation.TechnicianLocation{})
	if userID != nil {
		query = query.Where("user_id = ?", *userID)
	}
	if startDate != nil {
		query = query.Where("created_at >= ?", *startDate)
	}
	if endDate != nil {
		query = query.Where("created_at <= ?", *endDate)
	}
	query.Count(&total)
	offset := (page - 1) * perPage
	err := query.Order("created_at DESC").Offset(offset).Limit(perPage).Find(&locations).Error
	return locations, total, err
}
