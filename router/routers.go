package routers

import (
	"sdkim-go-project/internal/controllers"

	"github.com/gin-gonic/gin"
)

func GetHealth(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "I'm healthy!",
	})
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	public := router.Group("/api")
	{
		public.GET("/", GetHealth)
		public.GET("/me", controllers.GetMe)
	}

	return router

}
