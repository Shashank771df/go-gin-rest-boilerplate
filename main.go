package main

import (
	"go-gin-rest-boilerplate/database"
	"go-gin-rest-boilerplate/routes"
	"log"

	"github.com/gin-gonic/gin"
)

// @title GO Gin REST BPOILERPLATE
// @description This is a REST API built with Gin framework.
// @version 1.0
// @host localhost:8080
// @BasePath /api/v1
func main() {
	database.Init()
	e := gin.Default()
	routes.InitRoutes(e)
	e.Run(":8080")
	log.Default().Println("Server is running on port 8080")
}
