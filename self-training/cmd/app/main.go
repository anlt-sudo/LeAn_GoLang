package main

import (
	"self-training/internal/config"
	"self-training/internal/handler"
	"self-training/internal/model"
	"self-training/internal/repository"
	"self-training/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	db := config.DB

	// migrate tables
	db.AutoMigrate(&model.Category{}, &model.Product{})

	// init repos, services, handlers
	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	r := gin.Default()

	api := r.Group("/api/v1")
	{
		cat := api.Group("/categories")
		{
			cat.GET("", categoryHandler.GetAll)
			cat.GET("/:id", categoryHandler.GetByID)
			cat.POST("", categoryHandler.Create)
			cat.PUT("/:id", categoryHandler.Update)
			cat.DELETE("/:id", categoryHandler.Delete)
		}
		// TODO: thêm routes cho products
	}

	r.Run(":8080")
}
