package main

import "fmt"

func main() {
	// START OMIT
	total := 0
	a, b := 1, 1
	for a <= 4_000_000 { // HL
		if a%2 == 0 {
			total += a
		}
		a, b = b, a+b
	}
	fmt.Println(total)
	// END OMIT
}
