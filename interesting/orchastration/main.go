package main

import (
	"fmt"
	"sync"
)

func someFunc() {
	fmt.Println("someFunc called")
}

func someOtherFunc() {
	fmt.Println("someOtherFunc called")
}

func main() {
	// WaitGroup is best when everything needs to be independent
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		someFunc()
		wg.Done()
	}()

	go func() {
		someOtherFunc()
		wg.Done() // what happens if I forget this?
	}()

	wg.Wait()
	fmt.Println("Main function finished!")
}
