package controller

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Getdice(c *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	var myrandom = rand.Intn(6) + 1
	fmt.Println("Your dice rolled to:",myrandom)
	c.JSON(http.StatusOK, gin.H{"Rolled Dice is": myrandom})
}
