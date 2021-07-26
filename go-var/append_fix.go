package main

import (
	"fmt"
)

// APPEND_START OMIT
func appendN(values []int, val int, n int) []int { // HL
	for i := 0; i < n; i++ {
		values = append(values, val)
	}
	return values // HL
}

// APPEND_END OMIT

func main() {
	// MAIN_START OMIT
	values := []int{1, 2}
	values = appendN(values, 3, 3) // HL
	fmt.Println(values)
	// MAIN_END OMIT
}
