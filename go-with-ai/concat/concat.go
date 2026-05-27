package main

import (
	"fmt"
)

func ConcatSlices(a, b []string) []string {
	return append(a, b...)
}

func main() {
	a, b := []string{"It's", "a"}, []string{"bug!"}
	fmt.Println(ConcatSlices(a, b)) // [It's a bug!]

	fmt.Println(ConcatSlices(a[:1], b)) // [It's bug!]
	fmt.Println(a)                      // [It's bug!]
}
