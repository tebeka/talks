package main

import "fmt"

func main() {

	mat := [][]int{
		{1, -2, 3},
		{4, -5, 6},
	}

	found := false
	for r := range mat {
		for c := range mat[0] {
			if v := mat[r][c]; v < 0 {
				found = true
				fmt.Println("found", v)
				break
			}
		}
	}
	fmt.Println("negatives:", found)
}
