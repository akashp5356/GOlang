package main

import (
	"fmt"
	"log"
	"mongoapi/router"
	"net/http"
)

func main() {
	fmt.Println("Server is Starting ...")
	r:=router.Router()
	log.Fatal(http.ListenAndServe(":4000",r))
	fmt.Println("Listening at port 4000")
}


