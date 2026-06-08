package repositories

import (
	"github.com/pablomillaquen/speckit_golang_api/internal/domain/user"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(u *user.User) error {
	return r.db.Create(u).Error
}

func (r *UserRepository) FindByID(id uint64) (*user.User, error) {
	var u user.User
	err := r.db.Where("id = ?", id).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) FindByEmail(email string) (*user.User, error) {
	var u user.User
	err := r.db.Where("email = ?", email).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) FindAll(page, perPage int) ([]user.User, int64, error) {
	var users []user.User
	var total int64
	r.db.Model(&user.User{}).Count(&total)
	offset := (page - 1) * perPage
	err := r.db.Offset(offset).Limit(perPage).Find(&users).Error
	return users, total, err
}

func (r *UserRepository) Update(u *user.User) error {
	return r.db.Save(u).Error
}

func (r *UserRepository) Delete(id uint64) error {
	return r.db.Where("id = ?", id).Delete(&user.User{}).Error
}

func (r *UserRepository) FindByEmailExcludingID(email string, excludeID uint64) (*user.User, error) {
	var u user.User
	err := r.db.Where("email = ? AND id != ?", email, excludeID).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}
