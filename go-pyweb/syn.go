package main

import (
	"fmt"
)

// START_FUNC OMIT
func add(x, y int) int {
	return x + y
}

// END_FUNC OMIT
func main() {
	// START_FOR OMIT
	for i := 0; i < 3; i++ {
		fmt.Printf("%d + %d = %d\n", i, i, add(i, i))
	}

	j := 0
	for j < 3 {
		j = add(j, 1)
	}
	fmt.Printf("j = %d\n", j)
	// END_FOR OMIT
}
