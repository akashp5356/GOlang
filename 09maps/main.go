package main

import ("fmt")

func main(){
	//often known as key-value pair
	num:=make(map[int]string)
	num[1]="One"
	num[2]="Two"
	num[3]="Three"
	num[4]="Four"
	fmt.Println("Values are:",num)
	delete(num,2)
	fmt.Println("After deleting 2 :",num)
	//iterating through for
	for key,value:= range num{
		fmt.Printf("Key is %v with Value %v \n",key,value)
	}
}