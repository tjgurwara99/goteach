package main

import "fmt"

func main() {
	for {
		go fmt.Print(0)
		go fmt.Print(1)
	}
}
