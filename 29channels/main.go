package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in GO")
	mych := make(chan int)
	wg := &sync.WaitGroup{}
	// mych<-5
	// fmt.Println(<-mych)
	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		//fmt.Println(<-mych)
		val, ischannelopen := <-mych
		fmt.Println(ischannelopen)
		fmt.Println(val)

		wg.Done()
	}(mych, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		wg.Done()
		mych <- 0 //if channel has value 0 assigned the ischannelopen will be true and if there is no value assigned it will be 0 and ischannelopen will be false
		// mych<-6
		close(mych)
	}(mych, wg)
	wg.Wait()
}
