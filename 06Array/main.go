package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	//🧪 Example 1: Declaring and Initializing
	var fruit_list [4]string
	fruit_list[0] = ("apple")
	fruit_list[1]=("mango")
	fruit_list[2]=("pineapple")
	fruit_list[3]=("strawberry")
	fmt.Println("the fuits are :",fruit_list)

	//short syntax
	var veg_list =[3]string{"potato","beans","onion"}
	fmt.Println("the veg list is ",veg_list) 
	fmt.Println("the veg list length is ",len(veg_list)) 

	//shotest syntax
	animal:=[4]string{"gorilla","lion","tiger","lion"}
	fmt.Println("the content of animal list are ",animal)


}
