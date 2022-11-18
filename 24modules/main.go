package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello")
	greet()
	r := mux.NewRouter()
	r.HandleFunc("/", servermsg).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", r))
}

func greet() {
	fmt.Println("Good morning")
}

func servermsg(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello Hi</h1>"))
}
