package main

import (
    "fmt"
    "math"
)

func main() {
	numbers := [3]int {1, 2, 2}
	result := SquareSum(numbers[:])
	fmt.Println("El resultado de la suma:", result)
}

func SquareSum(numbers []int) int {
	var sum int
	for _, number := range numbers{
		sum = sum +  int(math.Pow(float64(number), 2))
	}
	return sum
}