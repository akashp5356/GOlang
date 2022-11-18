package main

import (
	"gin/middleware"
	"gin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	routes.InitRoutes(router)
	router.Run("localhost:8082") // listen and serve  "localhost:8082"
}
