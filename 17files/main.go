package main

import (
	"bufio"
	"io/ioutil"
	"io"
	"fmt"
	"os"
)

func main(){
	reader:=bufio.NewReader(os.Stdin)
	fmt.Println("Enter data to write in file:")
	data, _ :=reader.ReadString('\n')
	//Create creates or truncates the named file. If the file already exists, it is truncated
	file, err :=os.Create("./a.txt")
	if err!=nil{
		fmt.Println(err)
	}
	length,err :=io.WriteString(file,data)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("Length of file is:",length)
	defer file.Close()
read("./a.txt")
}
func read(filename string){
	data,err := ioutil.ReadFile(filename)
if err!=nil{
	panic(err)
}
fmt.Println("Data in file is:",string(data))
}