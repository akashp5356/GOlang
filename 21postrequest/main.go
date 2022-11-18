package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Post request")
 	Performpostrequest()
}
func Performpostrequest() {
	const myurl = "http://localhost:8000/post"
	requestbody :=strings.NewReader(`
	{"name":"golang",
	"price":0,
	"platform":"youtube"}
	`)
res,err:=http.Post(myurl,"application/json",requestbody)
if err!=nil{
	panic(err)
}
defer res.Body.Close()
con,_:=ioutil.ReadAll(res.Body)
fmt.Println(string(con))
}
