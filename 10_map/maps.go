package main

import (
	"fmt"
)

func main() {
	fmt.Println(`In Go (Golang), a map is a built-in data 
  type that associates keys with values, 
  like a dictionary or hash table in other languages.`)

	//mapName := make(map[keyType]valueType)
	fruit_list := make(map[string]int)

	//initalizing a map
	fruit_list["apple"] = 3
	fruit_list["mango"] = 5
	fruit_list["lecchi"] = 7

	//accesing all elements and elements
	fmt.Println(fruit_list)

	count, ok := fruit_list["apple"]
	fmt.Println("count is", count)
	//to check if value exist or not

	if ok {
		fmt.Println("the value exist is ", count)
	} else {
		fmt.Println("this value doesnt exist")
	}
	//deleting a value
	delete(fruit_list, "mango")
	fmt.Println("the total map is", fruit_list)

	//LOOP OR ITERATION
	for key, value := range fruit_list {
		fmt.Println(key, value)
	}

}
