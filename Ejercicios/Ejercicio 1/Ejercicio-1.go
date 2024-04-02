package main

import (
	"fmt"
	"math"
)

const decimales = 100

func main(){
	var altura float64;
	var base float64;

	fmt.Print("Ingresa la Altura: ")
	fmt.Scanln(&altura)
	fmt.Print("Ingresa la Base: ")
	fmt.Scanln(&base)

	hipotenusa := math.Sqrt(math.Pow(altura, 2) + math.Pow(base, 2))
	//hipotenusa := math.Hypot(altura, base)
	println(hipotenusa)
	area := (base * altura) / 2
	perimetro := base + altura + hipotenusa

	fmt.Printf("El area del triangulo es: %.2f\n", area)
	fmt.Printf("El perimetro del triangulo es: %.2f\n", perimetro)
}