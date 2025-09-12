package main

import (
	"go-shop-api/config"
	"go-shop-api/internal/model"
	"go-shop-api/internal/router"
)

func main() {
	config.ConnectDB()
	db := config.DB

	db.AutoMigrate(&model.Category{}, &model.Product{}, &model.User{}, &model.RefreshToken{})

	r := router.SetupRouter(db)

	r.Run(":8080")
}
