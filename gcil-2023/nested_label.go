package main

import "fmt"

func main() {
	// START OMIT
	mat := [][]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}

	hasEven := false
loop: // HL
	for r := range mat {
		for c := range mat[0] {
			if mat[r][c]%2 == 0 {
				hasEven = true
				fmt.Println("found")
				break loop // HL
			}
		}
	}
	fmt.Println("has even:", hasEven)
	// END OMIT
}
