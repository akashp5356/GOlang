package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://www.google.com:8080/search?Name=Ozil&Position=AMF&Number=10"

func main() {
	result, _ := url.Parse(myurl)	//parses a raw url into a URL structure.
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
qparam:=result.Query()
fmt.Printf("Type of qparam is:%T\n",qparam)
fmt.Println(qparam["Name"])
fmt.Println(qparam["Position"])
fmt.Println(qparam["Number"])

partsofurl:=&url.URL{
	Scheme:"https",
	Host:"www.google.com",
	Path:"/search",
}
//otherurl:=partsofurl.String()
fmt.Println(partsofurl.String())
}
