package main

import("fmt"
"time")

func main(){
	pt:=time.Now()
	fmt.Println("Time is:",pt)
	fmt.Println(pt.Format("02-01-2006 Monday 15:04:05"))
}