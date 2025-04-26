package main

import (
	"fmt"
	"time"
)

func main() {
	// Start two goroutines
	go greeter("Hello")
	go greeter("World")

	// Main thread sleeping to let goroutines finish
	time.Sleep(2 * time.Second)
	fmt.Println("Main function finished")
}

func greeter(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s, i)
		time.Sleep(300 * time.Millisecond) // simulate some work
	}
}