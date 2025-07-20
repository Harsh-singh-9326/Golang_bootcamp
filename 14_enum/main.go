package main

import "fmt"

// Define a custom type
type Status int

// Use iota to create enum values
const (
	Pending  Status = iota // 0
	Approved               // 1
	Rejected               // 2
)

func main() {
	var s Status = Approved
	fmt.Println(s) // Output: 1
}
