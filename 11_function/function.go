package main

import "fmt"

// named function
func add(a int, b int) int {
	return a + b
}

func main() {
	result := add(2, 3)
	fmt.Println("the result is ", result)

	//annonymous function
	square := func(a, b int) (int, int) {
		return a * a, b * b
	}
	result1, result2 := square(2, 3)
	fmt.Println("the square is ",result1)
	fmt.Println("the square  is ",result2)
	
	//takingonefunction to another
	// total := func(square func,a ,b int)


}
