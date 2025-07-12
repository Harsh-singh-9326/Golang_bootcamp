package main

import "fmt"

func main(){
	fmt.Println("welcome to go for loop")
	//classic for loop 
	for i:=1;i<3;i++{
		fmt.Println(i)
	}

	//go while loop
		
	i:=1
	for i<=3{
		if i==2 {
			i++
			continue
		}
		fmt.Println(i)
		i++

	}

	//1.22 range

	for i:= range 10{  //dont have to declared i value it will range to its range similar to python
		fmt.Println(i)
	}


}
