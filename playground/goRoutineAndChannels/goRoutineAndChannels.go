package main

import (
	"fmt"
	"time"
)

func worker(id int, ch chan string) {
	for i := 1; i <= 3; i++ {
		time.Sleep(time.Duration(id) * time.Second) // Different workers take different time
		ch <- fmt.Sprintf("Worker %d finished task %d", id, i)
	}
}

func main() {
	ch := make(chan string)

	// Start 3 workers
	go worker(1, ch)
	go worker(2, ch)
	go worker(3, ch)

	// Receive 9 messages (3 tasks x 3 workers)
	for i := 0; i < 9; i++ {
		fmt.Println(<-ch) // Receive result from channel
	}

	fmt.Println("All work done 🚀")
}