package routes

import (
	"go-gin-rest-boilerplate/middlewares"

	_ "go-gin-rest-boilerplate/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func InitRoutes(engine *gin.Engine) {
	e := engine.Group(BasePath())
	InitMiddleware(e)
	RegisterProtectedRoutes(e)
	RegisterPublicRoutes(e)
	RegisterUtilityRoutes(e)
	// Serve Swagger documentation
	// engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func InitMiddleware(engine *gin.RouterGroup) {
	engine.Use(middlewares.CORSMiddleware())
}

func BasePath() string {
	return "/api/v1"
}
