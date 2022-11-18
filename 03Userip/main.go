package main

import ("fmt"
"bufio"	//implements buffered I/O
"os")

func main() {
	// := is known as walrus operator
	reader:=bufio.NewReader(os.Stdin)
	fmt.Println("Enter you Name:")
	//ReadString reads until the first occurrence of delim in the input. In this case new line is delimiter
	input, _ :=reader.ReadString('\n')
	fmt.Println("Hello ",input)
}