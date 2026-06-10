// @title Service Management API
// @version 1.0.0
// @description API for managing technical services, clients, equipment, work orders, and materials.
// @termsOfService https://swagger.io/terms/
// @contact.name API Support
// @contact.email support@speckit.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @Security BearerAuth
package main

import (
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/pablomillaquen/speckit_golang_api/configs"
	_ "github.com/pablomillaquen/speckit_golang_api/docs"
	"github.com/pablomillaquen/speckit_golang_api/internal/auth"
	"github.com/pablomillaquen/speckit_golang_api/internal/domain/audit"
	"github.com/pablomillaquen/speckit_golang_api/internal/domain/client"
	"github.com/pablomillaquen/speckit_golang_api/internal/domain/technicianlocation"
	"github.com/pablomillaquen/speckit_golang_api/internal/domain/equipment"
	"github.com/pablomillaquen/speckit_golang_api/internal/domain/material"
	"github.com/pablomillaquen/speckit_golang_api/internal/domain/user"
	"github.com/pablomillaquen/speckit_golang_api/internal/domain/workorder"
	"github.com/pablomillaquen/speckit_golang_api/internal/handlers"
	"github.com/pablomillaquen/speckit_golang_api/internal/middleware"
	"github.com/pablomillaquen/speckit_golang_api/internal/repositories"
	"github.com/pablomillaquen/speckit_golang_api/internal/services"
	"github.com/pablomillaquen/speckit_golang_api/pkg/database"
	"github.com/pablomillaquen/speckit_golang_api/pkg/logger"
)

func main() {
	cfg := configs.Load()

	logger.Init()

	db, err := database.Connect(cfg.Database)
	if err != nil {
		logger.Error("Failed to connect to database: %v", err)
		os.Exit(1)
	}

	if err := database.AutoMigrate(db,
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
		logger.Error("Failed to run migrations: %v", err)
		os.Exit(1)
	}

	database.SeedAdmin(db)

	jwtService := auth.NewJWTService(cfg.JWT)

	userRepo := repositories.NewUserRepository(db)
	clientRepo := repositories.NewClientRepository(db)
	equipmentTypeRepo := repositories.NewEquipmentTypeRepository(db)
	brandRepo := repositories.NewBrandRepository(db)
	equipmentModelRepo := repositories.NewEquipmentModelRepository(db)
	equipmentRepo := repositories.NewEquipmentRepository(db)
	materialRepo := repositories.NewMaterialRepository(db)
	workOrderRepo := repositories.NewWorkOrderRepository(db)
	workOrderNoteRepo := repositories.NewWorkOrderNoteRepository(db)
	workOrderMaterialRepo := repositories.NewWorkOrderMaterialRepository(db)
	auditRepo := repositories.NewAuditRepository(db)
	technicianLocationRepo := repositories.NewTechnicianLocationRepository(db)

	userService := services.NewUserService(userRepo, jwtService)
	clientService := services.NewClientService(clientRepo)
	equipmentTypeService := services.NewEquipmentTypeService(equipmentTypeRepo)
	brandService := services.NewBrandService(brandRepo)
	equipmentModelService := services.NewEquipmentModelService(equipmentModelRepo)
	equipmentService := services.NewEquipmentService(equipmentRepo)
	materialService := services.NewMaterialService(materialRepo)
	workOrderService := services.NewWorkOrderService(workOrderRepo, workOrderNoteRepo, workOrderMaterialRepo, auditRepo)
	auditService := services.NewAuditService(auditRepo)
	technicianLocationService := services.NewTechnicianLocationService(technicianLocationRepo)

	healthHandler := handlers.NewHealthHandler(db)
	authHandler := handlers.NewAuthHandler(userService)
	userHandler := handlers.NewUserHandler(userService)
	clientHandler := handlers.NewClientHandler(clientService)
	catalogHandler := handlers.NewCatalogHandler(equipmentTypeService, brandService, equipmentModelService)
	equipmentHandler := handlers.NewEquipmentHandler(equipmentService)
	materialHandler := handlers.NewMaterialHandler(materialService)
	workOrderHandler := handlers.NewWorkOrderHandler(workOrderService)
	auditHandler := handlers.NewAuditHandler(auditService)
	technicianLocationHandler := handlers.NewTechnicianLocationHandler(technicianLocationService)

	authMw := middleware.Auth(jwtService)
	rateLimiter := middleware.NewRateLimiter(cfg.RateLimit.RequestsPerMinute)

	r := gin.Default()
	r.Use(middleware.SecurityHeaders())
	r.Use(middleware.CORS(cfg.CORS.AllowedOrigins))
	r.Use(middleware.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	setupRoutes(r, authMw, middleware.RequiredRole,
		healthHandler, authHandler, userHandler,
		clientHandler, catalogHandler, equipmentHandler,
		materialHandler, workOrderHandler, auditHandler,
		technicianLocationHandler,
		rateLimiter,
	)

	logger.Info("Server starting on port %s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		logger.Error("Failed to start server: %v", err)
		os.Exit(1)
	}
}
