package main

import "fmt"

func main() {
	grade := string(GetGrade(80,70,70))
    fmt.Println(grade)
}

func GetGrade(a,b,c int) rune {
	sum := a + b + c
	prom := sum / 3
	
	switch {
	case prom >= 90 && prom <=100:
		return 'A'
	case prom >= 80 && prom < 90:
		return 'B'
	case prom >= 70 && prom < 80:
		return 'C'
	case prom >= 60 && prom < 70:
		return 'D'
	case prom >= 0 && prom < 60:
		return 'F'
	default:
		return '❌'
	}
	
}
