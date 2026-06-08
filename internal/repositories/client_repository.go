package repositories

import (
	"github.com/pablomillaquen/speckit_golang_api/internal/domain/client"
	"gorm.io/gorm"
)

type ClientRepository struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) *ClientRepository {
	return &ClientRepository{db: db}
}

func (r *ClientRepository) Create(c *client.Client) error {
	return r.db.Create(c).Error
}

func (r *ClientRepository) FindByID(id uint64) (*client.Client, error) {
	var c client.Client
	err := r.db.Where("id = ?", id).First(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *ClientRepository) FindByTaxID(taxID string) (*client.Client, error) {
	var c client.Client
	err := r.db.Where("tax_id = ?", taxID).First(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *ClientRepository) FindAll(page, perPage int, search string) ([]client.Client, int64, error) {
	var clients []client.Client
	var total int64
	query := r.db.Model(&client.Client{})
	if search != "" {
		query = query.Where("business_name LIKE ? OR tax_id LIKE ?", "%"+search+"%", "%"+search+"%")
	}
	query.Count(&total)
	offset := (page - 1) * perPage
	err := query.Offset(offset).Limit(perPage).Find(&clients).Error
	return clients, total, err
}

func (r *ClientRepository) Update(c *client.Client) error {
	return r.db.Save(c).Error
}

func (r *ClientRepository) Delete(id uint64) error {
	return r.db.Where("id = ?", id).Delete(&client.Client{}).Error
}

func (r *ClientRepository) FindByTaxIDExcludingID(taxID string, excludeID uint64) (*client.Client, error) {
	var c client.Client
	err := r.db.Where("tax_id = ? AND id != ?", taxID, excludeID).First(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}
