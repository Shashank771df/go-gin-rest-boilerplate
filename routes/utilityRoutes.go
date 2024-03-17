package routes

import "github.com/gin-gonic/gin"

func RegisterUtilityRoutes(engine *gin.RouterGroup) {
	//register utility routes
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong.",
		})
	})
	engine.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok.",
		})
	})
}
