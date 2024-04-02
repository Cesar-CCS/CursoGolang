package main

import "fmt"

func main() {
	fmt.Println("10, 70")
    fmt.Println(Hero(70,0))
}

func Hero(bullets, dragons int) bool {
	// your code
	kill := bullets/2
	for i := 1; i <= kill; i++ {
		dragons -= 1
	}

	fmt.Println(dragons)

	if dragons <= 0 {
		return true
	}

	return false
}