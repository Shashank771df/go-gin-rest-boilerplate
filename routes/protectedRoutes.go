package routes

import (
	"go-gin-rest-boilerplate/controllers"
	"go-gin-rest-boilerplate/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterProtectedRoutes(engine *gin.RouterGroup) {

	authGroup := engine.Group("/auth")

	authGroup.Use(middlewares.AuthHandler("admin"))
	{
		authGroup.GET("/getmessage", controllers.GetSecretText)
	}

}
