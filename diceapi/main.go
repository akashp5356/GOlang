package main

import (
	"dice/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Math in GO")
	router := gin.Default()
	routes.InitRoutes(router)
	router.Run("localhost:8082")


}
