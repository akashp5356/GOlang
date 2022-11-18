package main

import (
	"net/url"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Post request")
 	Performpostformrequest()
}
func Performpostformrequest() {
	const myurl = "http://localhost:8000/postform"
data:=url.Values{}
data.Add("Name","Mark")
data.Add("Number","65478965")
data.Add("Mail","mark@go.edv")
res,err:=http.PostForm(myurl,data)
if err!=nil{
	panic(err)
}
defer res.Body.Close()
con,_:=ioutil.ReadAll(res.Body)
fmt.Println(string(con))
}
