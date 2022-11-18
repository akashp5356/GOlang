package main

import ("fmt"
"sort")

func main(){
	//initialize slice
	var fruits=[]string {"apple","grapes"}
	fmt.Println(fruits)
	fruits=append(fruits,"watermelon","berries","papaya")
	fmt.Println(fruits)
	fruits=append(fruits[1:3])
	fmt.Println(fruits)

	 num:= make([]int,4)
	 num[0]=9
	 num[1]=2
	 num[2]=54
	 num[3]=98
	fmt.Println(num)
	 num=append(num,78,0,102)
	 fmt.Println(num)
sort.Ints(num)
fmt.Println(num)
var remo=[]string {"A","B","C","D","E"}
fmt.Println(remo)
//removing B from list
var index int=1
//triple dots used for accepting more arguments than it was meant for
remo=append(remo[:index],remo[index+1:]...)
fmt.Println(remo)	
}