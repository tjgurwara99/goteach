package main

import (
	"fmt"
	"sync"
	"time"
)

func someFunc() {
	fmt.Println("someFunc called")
}

func someOtherFunc() {
	fmt.Println("someOtherFunc called")
}

func main() {
	var wg sync.WaitGroup // WaitGroup is best when everything needs to be independent
	result := make(chan int)
	wg.Add(2)
	timer1 := time.NewTimer(1 * time.Second)
	timer2 := time.NewTimer(2 * time.Second)
	go func() {
		<-timer1.C
		someFunc()
		result <- 1
		wg.Done()
	}()

	go func() {
		<-timer2.C // change timer1 and timer2 around
		someOtherFunc()
		result <- 1
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(result) // necessary because the for loop blocks the code from executing further and would deadlock usually
	}()

	for i := range result {
		fmt.Println(i)
	}
	fmt.Println("Main function finished!")
}
