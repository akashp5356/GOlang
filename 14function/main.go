package main

import (
	"strings"
	"strconv"
	"fmt"
	"bufio"
	"os"
)

func main(){
	reader1:=bufio.NewReader(os.Stdin)
	fmt.Println("Enter number 1:")
	input1,_:=reader1.ReadString('\n')
	num1,_:=strconv.ParseFloat(strings.TrimSpace(input1),64)
	reader2:=bufio.NewReader(os.Stdin)
	fmt.Println("Enter number 2:")
	input2,_:=reader2.ReadString('\n')
	num2,_:=strconv.ParseFloat(strings.TrimSpace(input2),64)
	fmt.Println("Addition is:",add(num1,num2))

	prores,msg := proadd(5,10,5,20)
	fmt.Println("Pro result is:",prores)
	fmt.Println("Pro Message is:",msg)

}

func add(a float64,b float64)float64{
	return a+b
}
//can return two different types
func proadd(values ...int)(int,string){
	total:=0
	for _,val :=range values{
		total+=val
	}
	return total,"Hello"		//return the total and a txt message
}
