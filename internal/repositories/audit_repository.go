package repositories

import (
	"github.com/pablomillaquen/speckit_golang_api/internal/domain/audit"
	"gorm.io/gorm"
)

type AuditRepository struct {
	db *gorm.DB
}

func NewAuditRepository(db *gorm.DB) *AuditRepository {
	return &AuditRepository{db: db}
}

func (r *AuditRepository) Create(log *audit.AuditLog) error {
	return r.db.Create(log).Error
}

func (r *AuditRepository) FindAll(page, perPage int, filters map[string]interface{}) ([]audit.AuditLog, int64, error) {
	var logs []audit.AuditLog
	var total int64
	query := r.db.Model(&audit.AuditLog{})
	if entity, ok := filters["entity"]; ok && entity != "" {
		query = query.Where("entity = ?", entity)
	}
	if entityID, ok := filters["entity_id"]; ok && entityID != "" {
		query = query.Where("entity_id = ?", entityID)
	}
	if userID, ok := filters["user_id"]; ok && userID != "" {
		query = query.Where("user_id = ?", userID)
	}
	if action, ok := filters["action"]; ok && action != "" {
		query = query.Where("action = ?", action)
	}
	query.Count(&total)
	offset := (page - 1) * perPage
	err := query.Order("created_at DESC").Offset(offset).Limit(perPage).Find(&logs).Error
	return logs, total, err
}
