package router

import (
	"go-shop-api/internal/handler"
	"go-shop-api/internal/middleware"
	"go-shop-api/internal/repository"
	"go-shop-api/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// r.Use(middleware.Logging())
	r.Use(middleware.ErrorHandler())
	r.Use(middleware.CORS())

	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)
	authRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(authRepo)
	authHandler := handler.NewAuthHandler(authService)

	// hash, _ := authService.HashPassword("password123")
	// user := &model.User{Email: "admin@example.com", Password: hash, Role: "admin"}
	// _ = authRepo.Create(user)



	api := r.Group("/api/v1")
	{
		api.POST("/auth/login", authHandler.Login)
		protected := api.Group("/")
		protected.Use(middleware.AuthRequired())
		{
			protected.GET("/me", func(c *gin.Context) {
				claims, _ := middleware.GetClaims(c)
				c.JSON(200, gin.H{"user_id": claims.UserID, "email": claims.Email, "role": claims.Role})
			})
		
		}

		cat := api.Group("/categories")
		{
			cat.GET("", categoryHandler.GetAll)
			cat.GET("/:id", categoryHandler.GetByID)
			cat.POST("", categoryHandler.Create)
			cat.PUT("/:id", categoryHandler.Update)
			cat.DELETE("/:id", categoryHandler.Delete)
			cat.GET("/search", categoryHandler.Search)
		}
	}

	return r
}
