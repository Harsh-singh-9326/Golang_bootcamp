package main

import "fmt"

// public value
var Name string = "harsh singh"

func main() {
	//taking string as value
	var username string = "harshS"
	fmt.Println(username)
	fmt.Printf("the type of this value is %T \n", username)

	//taking num as variable
	var age int = 23
	fmt.Println(age)
	fmt.Printf("the type of this value is %T \n", age)

	//taking boolean as variable
	var isloogedin bool = true
	fmt.Println(isloogedin)
	fmt.Printf("the type of this value is %T /n", isloogedin)

	//implicit type
	var website = "www.harsh.com"
	fmt.Println(website)
	fmt.Printf("the type of this value is %T \n", website)

	fmt.Println(Name)

	//walrus operator nicknamed ✅ Allowed only inside functions.❌ Not allowed for package-level variables.❌ Cannot re-declare all existing variables:
	a := 5
   // a := 6  ❌ Error: no new variables on left side
	fmt.Println(a)


}
