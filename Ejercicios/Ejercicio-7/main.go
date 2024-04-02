package main

import "fmt"

func main() {
	fmt.Println(MakeNegative(-1))
}

func MakeNegative(x int) int {
	if x < 0 || x == 0{
		fmt.Println("this number is negative or zero")
		return x
	}
	
	return x * -1
}