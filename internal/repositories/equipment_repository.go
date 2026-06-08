package repositories

import (
	"github.com/pablomillaquen/speckit_golang_api/internal/domain/equipment"
	"gorm.io/gorm"
)

type EquipmentTypeRepository struct {
	db *gorm.DB
}

func NewEquipmentTypeRepository(db *gorm.DB) *EquipmentTypeRepository {
	return &EquipmentTypeRepository{db: db}
}

func (r *EquipmentTypeRepository) Create(et *equipment.EquipmentType) error {
	return r.db.Create(et).Error
}

func (r *EquipmentTypeRepository) FindByID(id uint64) (*equipment.EquipmentType, error) {
	var et equipment.EquipmentType
	err := r.db.First(&et, id).Error
	if err != nil {
		return nil, err
	}
	return &et, nil
}

func (r *EquipmentTypeRepository) FindAll() ([]equipment.EquipmentType, error) {
	var types []equipment.EquipmentType
	err := r.db.Find(&types).Error
	return types, err
}

func (r *EquipmentTypeRepository) Update(et *equipment.EquipmentType) error {
	return r.db.Save(et).Error
}

func (r *EquipmentTypeRepository) Delete(id uint64) error {
	return r.db.Delete(&equipment.EquipmentType{}, id).Error
}

type BrandRepository struct {
	db *gorm.DB
}

func NewBrandRepository(db *gorm.DB) *BrandRepository {
	return &BrandRepository{db: db}
}

func (r *BrandRepository) Create(b *equipment.Brand) error {
	return r.db.Create(b).Error
}

func (r *BrandRepository) FindByID(id uint64) (*equipment.Brand, error) {
	var b equipment.Brand
	err := r.db.First(&b, id).Error
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func (r *BrandRepository) FindByTypeID(typeID uint64) ([]equipment.Brand, error) {
	var brands []equipment.Brand
	err := r.db.Where("equipment_type_id = ?", typeID).Find(&brands).Error
	return brands, err
}

func (r *BrandRepository) FindAll() ([]equipment.Brand, error) {
	var brands []equipment.Brand
	err := r.db.Find(&brands).Error
	return brands, err
}

func (r *BrandRepository) Update(b *equipment.Brand) error {
	return r.db.Save(b).Error
}

func (r *BrandRepository) Delete(id uint64) error {
	return r.db.Delete(&equipment.Brand{}, id).Error
}

type EquipmentModelRepository struct {
	db *gorm.DB
}

func NewEquipmentModelRepository(db *gorm.DB) *EquipmentModelRepository {
	return &EquipmentModelRepository{db: db}
}

func (r *EquipmentModelRepository) Create(m *equipment.EquipmentModel) error {
	return r.db.Create(m).Error
}

func (r *EquipmentModelRepository) FindByID(id uint64) (*equipment.EquipmentModel, error) {
	var m equipment.EquipmentModel
	err := r.db.First(&m, id).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *EquipmentModelRepository) FindByBrandID(brandID uint64) ([]equipment.EquipmentModel, error) {
	var models []equipment.EquipmentModel
	err := r.db.Where("brand_id = ?", brandID).Find(&models).Error
	return models, err
}

func (r *EquipmentModelRepository) FindAll() ([]equipment.EquipmentModel, error) {
	var models []equipment.EquipmentModel
	err := r.db.Find(&models).Error
	return models, err
}

func (r *EquipmentModelRepository) Update(m *equipment.EquipmentModel) error {
	return r.db.Save(m).Error
}

func (r *EquipmentModelRepository) Delete(id uint64) error {
	return r.db.Delete(&equipment.EquipmentModel{}, id).Error
}

type EquipmentRepository struct {
	db *gorm.DB
}

func NewEquipmentRepository(db *gorm.DB) *EquipmentRepository {
	return &EquipmentRepository{db: db}
}

func (r *EquipmentRepository) Create(e *equipment.Equipment) error {
	return r.db.Create(e).Error
}

func (r *EquipmentRepository) FindByID(id uint64) (*equipment.Equipment, error) {
	var e equipment.Equipment
	err := r.db.Where("id = ?", id).First(&e).Error
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *EquipmentRepository) FindBySerialNumber(serial string) (*equipment.Equipment, error) {
	var e equipment.Equipment
	err := r.db.Where("serial_number = ?", serial).First(&e).Error
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *EquipmentRepository) FindAll(page, perPage int, search string, clientID uint64) ([]equipment.Equipment, int64, error) {
	var equipments []equipment.Equipment
	var total int64
	query := r.db.Model(&equipment.Equipment{})
	if search != "" {
		query = query.Where("serial_number LIKE ?", "%"+search+"%")
	}
	if clientID > 0 {
		query = query.Where("client_id = ?", clientID)
	}
	query.Count(&total)
	offset := (page - 1) * perPage
	err := query.Offset(offset).Limit(perPage).Find(&equipments).Error
	return equipments, total, err
}

func (r *EquipmentRepository) Update(e *equipment.Equipment) error {
	return r.db.Save(e).Error
}

func (r *EquipmentRepository) Delete(id uint64) error {
	return r.db.Where("id = ?", id).Delete(&equipment.Equipment{}).Error
}

func (r *EquipmentRepository) FindBySerialExcludingID(serial string, excludeID uint64) (*equipment.Equipment, error) {
	var e equipment.Equipment
	err := r.db.Where("serial_number = ? AND id != ?", serial, excludeID).First(&e).Error
	if err != nil {
		return nil, err
	}
	return &e, nil
}
