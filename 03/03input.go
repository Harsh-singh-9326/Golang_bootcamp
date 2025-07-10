package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	/*"Create a reader that listens to the keyboard input."os.Stdin: means "standard input", aka your keyboard.
	bufio.NewReader(...): creates a tool to read what you type.Think of reader as a mic pointed at your keyboard!*/

	fmt.Println("enter your name :")

	name, _ := reader.ReadString('\n')

	/*“Hey reader, listen until the user presses Enter (\n), then save what they typed into a variable called name.”name: what you typed.
	: means "ignore any errors for now" (we’re being lazy here 😄).*/

	fmt.Println("enter your full name :", name)

}
