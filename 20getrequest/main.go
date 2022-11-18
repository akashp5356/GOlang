package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Get request")
	Performgetrequest()
}
func Performgetrequest() {
	const myurl = "http://localhost:8000/get"
	res, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Println("Status code:", res.StatusCode)
	fmt.Println("Content length:", res.ContentLength)
	con, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(con))
	var responsestring strings.Builder
	count,_:=responsestring.Write(con)
	fmt.Println("Count is:",count)
	fmt.Println(responsestring.String())

}
