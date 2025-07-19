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
	fmt.Println("the square is ", result1)
	fmt.Println("the square  is ", result2)

	//takingonefunction to another
	/*✅ Syntax for Passing a Function into Another Function
	 func outerFunction(fn func(paramType1, paramType2) returnType, arg1 paramType1, arg2 paramType2) returnType {
	    // Use the passed function
	    return fn(arg1, arg2)
	}
	*/
	total := func(fn func(int, int) (int, int), a, b int) int {
		sqr1, sqr2 := fn(a, b)
		return sqr1 + sqr2
	}
	result0 := total(square, 2, 4)
	fmt.Println(result0)

	greet := func(name string) string {
		return "hello mrs welcome  " + name

	}
	name := greet("harsh")
	fmt.Println(name)

	/*A variadic function is a function that can accept a variable number of arguments of the same type.
	  🔧 Syntax:
	  func functionName(args ...Type) {
	     // args is a slice of Type
	 }*/
	total1 := 0
	value1 := func(a ...int) int {
		for _, a := range a {
			total1 += a
		}
		return total1

	}
	result9 := value1(1, 2, 3, 4, 5, 6)
	fmt.Println(result9)

}
