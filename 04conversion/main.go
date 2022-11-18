package main

import ("fmt"
"bufio"
"os"
"strconv"
"strings")

func main(){
	reader:=bufio.NewReader(os.Stdin)
	fmt.Println("Enter age:")
	input, _:=reader.ReadString('\n')
	fmt.Println("Your age is:",input)
	num, err:=strconv.ParseFloat(strings.TrimSpace(input),64)
if err!=nil {
	fmt.Println(err)
} else {
	fmt.Println("Adding 10 to your age:",num+10)
}
fmt.Println("Addition is:",10+10)
}