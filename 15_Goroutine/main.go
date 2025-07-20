package main

import (
	"fmt"
	"time"
)
/*
A goroutine is a lightweight thread of execution managed by the Go runtime.
It lets you run functions concurrently — meaning multiple things can happen at the same time in your program.
*/

func sayHello() {
	fmt.Println("Hello from goroutine")
}

func main() {
	go sayHello()              // Runs in the background
	fmt.Println("Hello main")  // Main thread runs this

	time.Sleep(1 * time.Second) // Wait for goroutine to finish (or else main ends too early)
}
