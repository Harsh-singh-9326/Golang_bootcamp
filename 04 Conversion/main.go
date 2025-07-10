package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter your review on this simple code :")
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n')
	fmt.Println("your review to our code is ", value)
	//the value is in string so we have to convert into int
	fmt.Printf("what is your type %T \n", value)

	total_review, err := strconv.ParseFloat(strings.TrimSpace(value), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("add 1 review my rule my choice ", total_review+1)
	}
}
