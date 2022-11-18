package main

import("fmt")

func main(){
	num :=20
	var ptr=&num
	//pointer ensures that the actual value is used instead of the copy
	fmt.Println("Pointer is:",ptr)
	fmt.Println("Value of pointer is:",*ptr)
*ptr+=10
fmt.Println("Value after adding 10 is:",num)
}