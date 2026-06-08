package repositories

import (
	"github.com/pablomillaquen/speckit_golang_api/internal/domain/workorder"
	"gorm.io/gorm"
)

type WorkOrderRepository struct {
	db *gorm.DB
}

func NewWorkOrderRepository(db *gorm.DB) *WorkOrderRepository {
	return &WorkOrderRepository{db: db}
}

func (r *WorkOrderRepository) Create(wo *workorder.WorkOrder) error {
	return r.db.Create(wo).Error
}

func (r *WorkOrderRepository) FindByID(id uint64) (*workorder.WorkOrder, error) {
	var wo workorder.WorkOrder
	err := r.db.Where("id = ?", id).First(&wo).Error
	if err != nil {
		return nil, err
	}
	return &wo, nil
}

func (r *WorkOrderRepository) FindAll(page, perPage int, filters map[string]interface{}) ([]workorder.WorkOrder, int64, error) {
	var orders []workorder.WorkOrder
	var total int64
	query := r.db.Model(&workorder.WorkOrder{})
	if status, ok := filters["status"]; ok && status != "" {
		query = query.Where("status = ?", status)
	}
	if priority, ok := filters["priority"]; ok && priority != "" {
		query = query.Where("priority = ?", priority)
	}
	if clientID, ok := filters["client_id"]; ok && clientID != "" {
		query = query.Where("client_id = ?", clientID)
	}
	if technicianID, ok := filters["technician_id"]; ok && technicianID != "" {
		query = query.Where("technician_id = ?", technicianID)
	}
	if dateFrom, ok := filters["date_from"]; ok && dateFrom != "" {
		query = query.Where("scheduled_date >= ?", dateFrom)
	}
	if dateTo, ok := filters["date_to"]; ok && dateTo != "" {
		query = query.Where("scheduled_date <= ?", dateTo)
	}
	query.Count(&total)
	offset := (page - 1) * perPage
	err := query.Order("created_at DESC").Offset(offset).Limit(perPage).Find(&orders).Error
	return orders, total, err
}

func (r *WorkOrderRepository) Update(wo *workorder.WorkOrder) error {
	return r.db.Save(wo).Error
}

func (r *WorkOrderRepository) Delete(id uint64) error {
	return r.db.Where("id = ?", id).Delete(&workorder.WorkOrder{}).Error
}

type WorkOrderNoteRepository struct {
	db *gorm.DB
}

func NewWorkOrderNoteRepository(db *gorm.DB) *WorkOrderNoteRepository {
	return &WorkOrderNoteRepository{db: db}
}

func (r *WorkOrderNoteRepository) Create(note *workorder.WorkOrderNote) error {
	return r.db.Create(note).Error
}

func (r *WorkOrderNoteRepository) FindByWorkOrderID(woID uint64) ([]workorder.WorkOrderNote, error) {
	var notes []workorder.WorkOrderNote
	err := r.db.Where("work_order_id = ?", woID).Order("created_at ASC").Find(&notes).Error
	return notes, err
}

type WorkOrderMaterialRepository struct {
	db *gorm.DB
}

func NewWorkOrderMaterialRepository(db *gorm.DB) *WorkOrderMaterialRepository {
	return &WorkOrderMaterialRepository{db: db}
}

func (r *WorkOrderMaterialRepository) Create(wm *workorder.WorkOrderMaterial) error {
	return r.db.Create(wm).Error
}

func (r *WorkOrderMaterialRepository) FindByWorkOrderID(woID uint64) ([]workorder.WorkOrderMaterial, error) {
	var materials []workorder.WorkOrderMaterial
	err := r.db.Where("work_order_id = ?", woID).Order("created_at ASC").Find(&materials).Error
	return materials, err
}
