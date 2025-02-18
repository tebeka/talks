package main

import (
	"fmt"
)

func main() {
	// START OMIT
	total := 0
	for i := 0; i < 1000; i++ { // HL
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}
	fmt.Println(total)
	// END OMIT
}
