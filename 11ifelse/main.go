package main


import ("fmt"
"os"
"bufio"
"strings"
"strconv")

func main(){
	reader:=bufio.NewReader(os.Stdin)
	fmt.Println("Enter age:")
input,_:=reader.ReadString('\n')
num,err:=strconv.ParseFloat(strings.TrimSpace(input),64)
if err!=nil{
	fmt.Println("Error")
} else {
	if num>18{
		fmt.Println("You are 18+")
	} else if num<18{
		fmt.Println("You are under 18")
	} else {
		fmt.Println("You are 18")
	}
}
}