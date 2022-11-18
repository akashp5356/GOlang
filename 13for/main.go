package main

import (
	"fmt"
)

func main(){
	fmt.Println("Loops")
	da:=[]string{"A","B","C","D"}	//slice
	for i:=0;i<len(da);i++{
		fmt.Println(da[i])
	}
	fmt.Println("Using range")
	for i:= range da {
		fmt.Println(da[i])
	}
	fmt.Println("Showing index")
	for i:= range da {
		fmt.Printf("Index is %v with data %v\n",i,da[i])
	}
}