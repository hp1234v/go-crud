package main

import (
	"fmt"
)

func main() {
	ch := make(chan string) // create a channel of strings

	go func() {
		fmt.Println("Goroutine: sending value...")
		ch <- "Hello from Goroutine!"
	}()

	fmt.Println("Main: waiting for value...")
	value := <-ch // blocks until it receives
	fmt.Println("Main: got value -", value)
}