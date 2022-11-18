package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)
var signals=[]string{"test"}
var wg sync.WaitGroup	//usually are pointers
var mut sync.Mutex

func main() {
	//go greet("Hello")
	//greet("World")
	websitelist:=[]string{
		"https://google.com",
		"https://facebook.com",
		"https://youtube.com",
	}
	for _, v := range websitelist {
		go getstatus(v)	//go keyword used for routine
		wg.Add(1)
	}
	wg.Wait()	//dont exit yet 
	fmt.Println(signals)
}

func greet(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(10 * time.Millisecond)
		fmt.Println(s)
	}
}

func getstatus(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("In endpoint")
	} else {
		mut.Lock()
		signals=append(signals,endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}

}
