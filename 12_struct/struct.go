package main

import (
	"fmt"

	
)

// A struct is a custom data type that groups related fields (variables) under one name.

// Only fields with capitalized names can be seen outside the package.
type person struct {
	Name     string //exported start with capital letter
	Age      int    //exported start with capital letter
	educated bool   //cant  exported dont start with capital letter
}
type post struct {
	id          int
	tittle      string
	description string
	is_posted   bool
}

type rectangle struct {
	width  float32
	height float32
}

func (r rectangle) Area() float32 {
	return r.width * r.height
}
func main() {

	//assigning the value
	Person := person{Name: "harsh", Age: 22, educated: true}
	fmt.Println(Person)

	Post := post{id: 1, tittle: "learning golang", description: "enjoying a lot", is_posted: true}
	fmt.Println(Post)

	r := rectangle{width: 15,height: 10}
	fmt.Println(r.Area())


}
