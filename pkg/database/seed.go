package database

import (
	"github.com/pablomillaquen/speckit_golang_api/internal/auth"
	"github.com/pablomillaquen/speckit_golang_api/internal/domain/user"
	"github.com/pablomillaquen/speckit_golang_api/pkg/logger"
	"gorm.io/gorm"
)

func SeedAdmin(db *gorm.DB) {
	var count int64
	db.Model(&user.User{}).Where("email = ?", "admin@speckit.com").Count(&count)
	if count > 0 {
		logger.Info("Admin user already exists, skipping seed")
		return
	}
	hashed, err := auth.HashPassword("Admin123!")
	if err != nil {
		logger.Error("Failed to hash admin password: %v", err)
		return
	}
	admin := user.User{
		Name:     "Administrator",
		Email:    "admin@speckit.com",
		Password: hashed,
		Role:     user.RoleAdmin,
		Active:   true,
	}
	if err := db.Create(&admin).Error; err != nil {
		logger.Error("Failed to seed admin user: %v", err)
		return
	}
	logger.Info("Admin user seeded successfully (admin@speckit.com / Admin123!)")
}
