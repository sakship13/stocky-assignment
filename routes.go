package routes

import (
	"github.com/gin-gonic/gin"
	"stocky/controllers"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/reward", controllers.PostReward)
	router.GET("/stats/:userId", controllers.GetStats)
}
