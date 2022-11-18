package routes

import (
	"dice/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.GET("/api/dice", controller.Getdice)
}
