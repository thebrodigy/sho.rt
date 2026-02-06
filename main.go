package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/thebrodigy/sho.rt/db"
	"github.com/thebrodigy/sho.rt/model"
	"github.com/thebrodigy/sho.rt/routes"
)

func main() {
	godotenv.Load()
	db.Connect()
	db.DB.AutoMigrate(&model.ShortUrl{})

	router := gin.Default()
	routes.ApiRoutes(router)
	router.Run(":8080")
}
