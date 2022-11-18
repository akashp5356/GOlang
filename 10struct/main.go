package main

import ("fmt")

func main(){
	fmt.Println("Structs in go")
	data:=Details{"ABC",20,987456321}	//user defined values provided to struct
	fmt.Println(data)
	fmt.Printf("Details with naming conventions: %+v\n",data)
	fmt.Printf("I am %v \nMy age is %v \nContact is %v",data.Name,data.Age,data.Mob)
}
//initalize structure
type Details struct{
	Name string
	Age int
Mob int
}