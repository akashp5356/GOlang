package main

import (
	"fmt"
)

func main(){
	//defer line is executed at the end of the function
	defer fmt.Println("World")
	defer fmt.Println("New")
// lifo order for  multiple defer
	fmt.Println("Hello")
	mydefer()
	fmt.Println("")

}

func mydefer(){
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}