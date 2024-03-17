package routes

import (
	"go-gin-rest-boilerplate/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterPublicRoutes(engine *gin.RouterGroup) {
	//register public routes
	authController := new(controllers.AuthController)
	engine.POST("/login", authController.HandleLogin)
	engine.GET("/getmessage", controllers.GetPublicText)
}
