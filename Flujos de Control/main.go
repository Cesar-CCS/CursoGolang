package main

import (
	"fmt"
	//"runtime"
	//"time"
)

func main() {
	//Bucle con condicion
	var i int
	for i <= 10{
		fmt.Println(i)
		i++
	}

	//Bucle clasico
	for j := 0; j <= 10; j++ {
        fmt.Println(j)
    }
}

/*func main() {
	/*t := time.Now()
	hora := t.Hour()

	if t := time.Now(); t.Hour() < 12 {
		fmt.Println("¡Mañana!")
	} else if t.Hour() < 17 {
		fmt.Println("¡tarde!")
	} else {
		fmt.Println("¡noche!")
	}

	//os := runtime.GOOS

	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Go run -> Windows")
	case "linux":
		fmt.Println("Go run -> Linux")
	case "darwin":
		fmt.Println("Go run -> macOS")
	default:
		fmt.Println("Go run -> Otro OS")
	}
}*/