package main

import (
	"fmt"
)

// APPEND_START OMIT
func appendN(values []int, val int, n int) {
	for i := 0; i < n; i++ {
		values = append(values, val)
	}
}

// APPEND_END OMIT

func main() {
	// MAIN_START OMIT
	values := []int{1, 2}
	appendN(values, 3, 3)
	fmt.Println(values)
	// MAIN_END OMIT
}
