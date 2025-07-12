package main

import "fmt"

func main() {
	//simple if_else
	i := 10

	if i >= 18 {
		fmt.Println("i am an adult")

	} else if i >= 12 {
		fmt.Println("i a am teeneger")
	} else {
		fmt.Println("just an kid")
	}
	//to check whether number is even or odd
	num := 14
	if num%2 == 0 {
		fmt.Println("number is even", num)
	} else {
		fmt.Println("its odd")
	}
}
