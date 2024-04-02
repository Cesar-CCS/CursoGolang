package main

import "fmt"

func main() {
	suma, resta := calc(4,3)
	fmt.Println("La suma es:",suma)
	fmt.Println("La resta es:",resta)

    saludo := hello("Cesar")
	fmt.Println(saludo)
}

func hello(name string) string {
	return "Hola,"+name
}

func calc(a, b int) (sum , res int) {
	sum = a + b
	res = a - b
	return 
}