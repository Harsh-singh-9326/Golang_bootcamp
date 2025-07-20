package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Mark this goroutine as done when it finishes

	fmt.Printf("Worker %d starting\n", id)
	// Simulate some work
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup // Create a WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)         // Tell WaitGroup: one more goroutine is starting
		go worker(i, &wg) // Start the goroutine and pass pointer to wg
	}

	wg.Wait() // Block here until all goroutines call Done()
	fmt.Println("All workers completed")
}
