package main

import (
	"io/ioutil"
	"fmt"
	"net/http"
)

const url ="https://www.google.com/"
func main(){
res,err := http.Get(url)
if err!=nil{
	panic(err)
}
fmt.Printf("Response is of type %T\n",res)
data,err:=ioutil.ReadAll(res.Body)
if err!=nil{
	panic(err)
}
fmt.Println(string(data))
res.Body.Close()
}