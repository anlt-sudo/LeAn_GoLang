package router

import (
	"fmt"
	"go-shop-api/internal/handler"
	"go-shop-api/internal/middleware"
	"go-shop-api/internal/model"
	"go-shop-api/internal/repository"
	"go-shop-api/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.Use(middleware.ErrorHandler())
	r.Use(middleware.CORS())

	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	authRepo := repository.NewUserRepository(db)
	refreshTokenRepo := repository.NewRefreshTokenRepository(db)
	authService := service.NewAuthService(authRepo, refreshTokenRepo)
	authHandler := handler.NewAuthHandler(authService)

	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	seedAdmin(db, authService)

	api := r.Group("/api/v1")
	{
		api.POST("/auth/login", authHandler.Login)
		api.POST("/auth/register", authHandler.Register)
		api.POST("/auth/refresh", authHandler.Refresh)
		api.POST("/auth/logout", authHandler.Logout)

		protected := api.Group("/")
		protected.Use(middleware.AuthRequired())
		{
			protected.GET("/me", func(c *gin.Context) {
				claims, _ := middleware.GetClaims(c)
				c.JSON(200, gin.H{
					"user_id": claims.UserID,
					"email":   claims.Email,
					"role":    claims.Role,
				})
			})
		}

		cat := api.Group("/categories")
		{
			cat.GET("", categoryHandler.GetAll)
			cat.GET("/:id", categoryHandler.GetByID)
			cat.GET("/search", categoryHandler.Search)

			cat.Use(middleware.AuthRequired(), middleware.RoleRequired("admin"))
			{
				cat.POST("", categoryHandler.Create)
				cat.PUT("/:id", categoryHandler.Update)
				cat.DELETE("/:id", categoryHandler.Delete)
			}
		}

		products := api.Group("/products")
		{
			products.GET("", productHandler.GetAll)
			products.GET("/:id", productHandler.GetByID)
			products.GET("/search", productHandler.Search)

			products.Use(middleware.AuthRequired(), middleware.RoleRequired("admin"))
			{
				products.POST("", productHandler.Create)
				products.PUT("/:id", productHandler.Update)
				products.DELETE("/:id", productHandler.Delete)
			}
		}
	}

	return r
}

func seedAdmin(db *gorm.DB, authService *service.AuthService) {
	var count int64
	db.Model(&model.User{}).Where("role = ?", "admin").Count(&count)
	if count == 0 {
		hash, _ := authService.HashPassword("password123")
		admin := &model.User{
			Email:    "admin@example.com",
			Password: hash,
			Role:     "admin",
		}
		if err := db.Create(admin).Error; err != nil {
			panic(err)
		}
		access, refresh, _, err := authService.Authenticate(admin.Email, "password123")
		if err != nil {
			panic(err)
		}
		fmt.Println("Default admin created:")
		fmt.Println("Email: admin@example.com")
		fmt.Println("Password: password123")
		fmt.Println("Access Token:", access)
		fmt.Println("Refresh Token:", refresh)
	}
}
