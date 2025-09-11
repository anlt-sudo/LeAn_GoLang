package router

import (
	"go-shop-api/internal/handler"
	"go-shop-api/internal/repository"
	"go-shop-api/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)


	api := r.Group("/api/v1")
	{
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
