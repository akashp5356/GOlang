package main

import (
	"fmt"
	"math/big"

	//"math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Math in GO")
	// var num1 int=2
	// var num2 float64=4
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println("Dice is:",rand.Intn(6)+1)
	myrandom, _ := rand.Int(rand.Reader, big.NewInt(6))
fmt.Println(myrandom)
}
