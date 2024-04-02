package main

import (
	"fmt"
	"math/rand"
)

func main() {
	play()
	newGame()
}

func play(){
	random := rand.Intn(100)
    var numberUser int
	var intentos int
	const maxIntentos = 10

	for intentos < maxIntentos{
		fmt.Printf("Ingrese un numero del 1 al 100 (intentos: %d)\n", intentos + 1)
		fmt.Scanln(&numberUser)

		if numberUser == random {
			fmt.Printf("Ganaste, %d es el numero que pensé", random)
			return
		} else if numberUser < random {
			fmt.Println("El numero a adivinar es mayor")
		}else if numberUser > random{
			fmt.Println("El numero a adivinar es menor")
		}
		intentos++
	}

	fmt.Printf("Perdiste no tienes mas intentos, %d es el numero que pensé", random)
}

func newGame(){
	var eleccion string 
	fmt.Println("\n¿Quieres jugar nuevamente? (s/n): ")
	fmt.Scanln(&eleccion)

	for eleccion == "s" {
		play()
        fmt.Println("\n¿Quieres jugar nuevamente? (s/n): ")
        fmt.Scanln(&eleccion)
	}
}