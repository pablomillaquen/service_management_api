package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pablomillaquen/speckit_golang_api/internal/handlers"
	"github.com/pablomillaquen/speckit_golang_api/internal/middleware"
)

func setupRoutes(
	r *gin.Engine,
	authMw gin.HandlerFunc,
	rbacMiddleware func(...string) gin.HandlerFunc,
	healthHandler *handlers.HealthHandler,
	authHandler *handlers.AuthHandler,
	userHandler *handlers.UserHandler,
	clientHandler *handlers.ClientHandler,
	catalogHandler *handlers.CatalogHandler,
	equipmentHandler *handlers.EquipmentHandler,
	materialHandler *handlers.MaterialHandler,
	workOrderHandler *handlers.WorkOrderHandler,
	auditHandler *handlers.AuditHandler,
	technicianLocationHandler *handlers.TechnicianLocationHandler,
	rateLimiter *middleware.RateLimiter,
) {
	api := r.Group("/api/v1")
	{
		api.GET("/health", healthHandler.Check)

		auth := api.Group("/auth")
		{
			auth.POST("/login", rateLimiter.Middleware(), authHandler.Login)
			auth.POST("/refresh", authHandler.RefreshToken)
		}

		secured := api.Group("")
		secured.Use(authMw)
		{
			secured.POST("/auth/change-password", authHandler.ChangePassword)

			users := secured.Group("/users")
			users.Use(rbacMiddleware("administrator"))
			{
				users.POST("", userHandler.Create)
				users.GET("", userHandler.FindAll)
				users.GET("/:id", userHandler.FindByID)
				users.PUT("/:id", userHandler.Update)
				users.DELETE("/:id", userHandler.Delete)
			}

			clients := secured.Group("/clients")
			clients.Use(rbacMiddleware("administrator", "technician"))
			{
				clients.POST("", clientHandler.Create)
				clients.GET("", clientHandler.FindAll)
				clients.GET("/:id", clientHandler.FindByID)
				clients.PUT("/:id", clientHandler.Update)
				clients.DELETE("/:id", clientHandler.Delete)
			}

			catalog := secured.Group("/catalog")
			catalog.Use(rbacMiddleware("administrator"))
			{
				catalog.POST("/types", catalogHandler.CreateType)
				catalog.GET("/types", catalogHandler.ListTypes)
				catalog.GET("/types/:id", catalogHandler.GetType)
				catalog.PUT("/types/:id", catalogHandler.UpdateType)
				catalog.DELETE("/types/:id", catalogHandler.DeleteType)

				catalog.POST("/brands", catalogHandler.CreateBrand)
				catalog.GET("/brands/by-type/:typeId", catalogHandler.ListBrandsByType)
				catalog.GET("/brands/:id", catalogHandler.GetBrand)
				catalog.PUT("/brands/:id", catalogHandler.UpdateBrand)
				catalog.DELETE("/brands/:id", catalogHandler.DeleteBrand)

				catalog.POST("/models", catalogHandler.CreateModel)
				catalog.GET("/models/by-brand/:brandId", catalogHandler.ListModelsByBrand)
				catalog.GET("/models/:id", catalogHandler.GetModel)
				catalog.PUT("/models/:id", catalogHandler.UpdateModel)
				catalog.DELETE("/models/:id", catalogHandler.DeleteModel)
			}

			equipment := secured.Group("/equipment")
			equipment.Use(rbacMiddleware("administrator", "technician"))
			{
				equipment.POST("", equipmentHandler.Create)
				equipment.GET("", equipmentHandler.FindAll)
				equipment.GET("/:id", equipmentHandler.FindByID)
				equipment.PUT("/:id", equipmentHandler.Update)
				equipment.DELETE("/:id", equipmentHandler.Delete)
			}

			materials := secured.Group("/materials")
			materials.Use(rbacMiddleware("administrator"))
			{
				materials.POST("", materialHandler.Create)
				materials.GET("", materialHandler.FindAll)
				materials.GET("/:id", materialHandler.FindByID)
				materials.PUT("/:id", materialHandler.Update)
				materials.DELETE("/:id", materialHandler.Delete)
			}

			workOrders := secured.Group("/work-orders")
			workOrders.Use(rbacMiddleware("administrator", "technician"))
			{
				workOrders.POST("", workOrderHandler.Create)
				workOrders.GET("", workOrderHandler.FindAll)
				workOrders.GET("/:id", workOrderHandler.FindByID)
				workOrders.PUT("/:id", workOrderHandler.Update)
				workOrders.DELETE("/:id", workOrderHandler.Delete)

				workOrders.POST("/:id/assign", workOrderHandler.AssignTechnician)
				workOrders.POST("/:id/status", workOrderHandler.ChangeStatus)
				workOrders.POST("/:id/notes", workOrderHandler.AddNote)
				workOrders.POST("/:id/materials", workOrderHandler.AddMaterial)
			}

			technicianLocations := secured.Group("/technician-locations")
		{
			technicianLocations.POST("", rbacMiddleware("technician"), technicianLocationHandler.Create)
			technicianLocations.GET("", rbacMiddleware("administrator", "technician"), technicianLocationHandler.FindAll)
		}

		audit := secured.Group("/audit")
			audit.Use(rbacMiddleware("administrator"))
			{
				audit.GET("", auditHandler.FindAll)
			}
		}
	}
}
