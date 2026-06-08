package repositories

import (
	"github.com/pablomillaquen/speckit_golang_api/internal/domain/material"
	"gorm.io/gorm"
)

type MaterialRepository struct {
	db *gorm.DB
}

func NewMaterialRepository(db *gorm.DB) *MaterialRepository {
	return &MaterialRepository{db: db}
}

func (r *MaterialRepository) Create(m *material.Material) error {
	return r.db.Create(m).Error
}

func (r *MaterialRepository) FindByID(id uint64) (*material.Material, error) {
	var m material.Material
	err := r.db.Where("id = ?", id).First(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *MaterialRepository) FindByCode(code string) (*material.Material, error) {
	var m material.Material
	err := r.db.Where("code = ?", code).First(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *MaterialRepository) FindAll(page, perPage int) ([]material.Material, int64, error) {
	var materials []material.Material
	var total int64
	r.db.Model(&material.Material{}).Count(&total)
	offset := (page - 1) * perPage
	err := r.db.Offset(offset).Limit(perPage).Find(&materials).Error
	return materials, total, err
}

func (r *MaterialRepository) Update(m *material.Material) error {
	return r.db.Save(m).Error
}

func (r *MaterialRepository) Delete(id uint64) error {
	return r.db.Where("id = ?", id).Delete(&material.Material{}).Error
}

func (r *MaterialRepository) FindByCodeExcludingID(code string, excludeID uint64) (*material.Material, error) {
	var m material.Material
	err := r.db.Where("code = ? AND id != ?", code, excludeID).First(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}
