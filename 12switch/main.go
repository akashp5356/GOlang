package main

import (
	"time"
	"math/rand"
	
	"fmt"
)

func main() {
	fmt.Println("Switch Case")
rand.Seed(time.Now().UnixNano())
num:=rand.Intn(3)
//fmt.Println("Number is :",num)
switch num{
case 0:fmt.Println("Number is 0")
case 1:fmt.Println("Number is 1")
case 2:fmt.Println("Number is 2")
}
}
