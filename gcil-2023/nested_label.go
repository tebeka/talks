package main

import "fmt"

func main() {
	// START OMIT
	mat := [][]int{
		{1, -2},
		{3, -4},
		{5, -6},
	}

	found := false
loop: // HL
	for r := range mat {
		for c := range mat[0] {
			if v := mat[r][c]; v < 0 {
				found = true
				fmt.Println("found", v)
				break loop // HL
			}
		}
	}
	fmt.Println("negatives:", found)
	// END OMIT
}
