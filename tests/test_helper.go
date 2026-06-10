package tests

import (
	"fmt"
	"sync/atomic"

	"github.com/pablomillaquen/speckit_golang_api/configs"
	"github.com/pablomillaquen/speckit_golang_api/internal/auth"
	"github.com/pablomillaquen/speckit_golang_api/internal/domain/audit"
	"github.com/pablomillaquen/speckit_golang_api/internal/domain/client"
	"github.com/pablomillaquen/speckit_golang_api/internal/domain/technicianlocation"
	"github.com/pablomillaquen/speckit_golang_api/internal/domain/equipment"
	"github.com/pablomillaquen/speckit_golang_api/internal/domain/material"
	"github.com/pablomillaquen/speckit_golang_api/internal/domain/user"
	"github.com/pablomillaquen/speckit_golang_api/internal/domain/workorder"
	"github.com/pablomillaquen/speckit_golang_api/internal/repositories"
	"github.com/pablomillaquen/speckit_golang_api/internal/services"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbCounter int64

func NewTestDB() *gorm.DB {
	id := atomic.AddInt64(&dbCounter, 1)
	dsn := fmt.Sprintf("file:test_%d?mode=memory&cache=private", id)
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic("failed to connect to test database: " + err.Error())
	}
	if err := db.AutoMigrate(
		&user.User{},
		&user.RefreshToken{},
		&client.Client{},
		&equipment.EquipmentType{},
		&equipment.Brand{},
		&equipment.EquipmentModel{},
		&equipment.Equipment{},
		&material.Material{},
		&workorder.WorkOrder{},
		&workorder.WorkOrderNote{},
		&workorder.WorkOrderMaterial{},
		&audit.AuditLog{},
		&technicianlocation.TechnicianLocation{},
	); err != nil {
		panic("failed to migrate test database: " + err.Error())
	}
	return db
}

func NewTestJWTService() *auth.JWTService {
	cfg := configs.JWTConfig{
		Secret: "test-secret",
	}
	return auth.NewJWTService(cfg)
}

func NewTestUserService(db *gorm.DB) *services.UserService {
	return services.NewUserService(
		repositories.NewUserRepository(db),
		NewTestJWTService(),
	)
}

func NewTestClientService(db *gorm.DB) *services.ClientService {
	return services.NewClientService(
		repositories.NewClientRepository(db),
	)
}

func NewTestEquipmentService(db *gorm.DB) *services.EquipmentService {
	return services.NewEquipmentService(
		repositories.NewEquipmentRepository(db),
	)
}

func NewTestWorkOrderService(db *gorm.DB) *services.WorkOrderService {
	return services.NewWorkOrderService(
		repositories.NewWorkOrderRepository(db),
		repositories.NewWorkOrderNoteRepository(db),
		repositories.NewWorkOrderMaterialRepository(db),
		repositories.NewAuditRepository(db),
	)
}

func NewTestTechnicianLocationService(db *gorm.DB) *services.TechnicianLocationService {
	return services.NewTechnicianLocationService(
		repositories.NewTechnicianLocationRepository(db),
	)
}

func NewTestMaterialService(db *gorm.DB) *services.MaterialService {
	return services.NewMaterialService(
		repositories.NewMaterialRepository(db),
	)
}
