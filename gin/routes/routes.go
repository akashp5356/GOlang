package routes

import (
	"gin/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {

	//	router.GET("/api/movies/:id", controller.Getmovie)
	router.GET("/api/movies", controller.Getallmovies)
	router.POST("/api/movie", controller.Createmovie)
	router.PUT("/api/movie/:id", controller.Markwatched)
	router.DELETE("/api/movie/:id", controller.Deletemovie)
	router.DELETE("/api/deleteallmovies", controller.Deleteallmovies)
}
