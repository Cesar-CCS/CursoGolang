package main

import "fmt"

func main() {
	fmt.Println(Past(1, 1, 1))
}

func Past(h, m, s int) int {
    // your code here
	if h > 23 && m > 59 && s > 59{
		return 0
	}

	h = (h * 3600 * 1000)
	m = (m * 60 * 1000)
	s = s * 1000

	return h + m + s
}