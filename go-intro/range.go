package main

import (
	"fmt"
)

func main() {
	// START OMIT
	a := []int{6, 7}
	for i := range a { // iterate over indexes
		fmt.Println(i)
	}
	for i, j := range a { // both index and value
		fmt.Println(i, j)
	}

	b := "xyz"
	for i, c := range b { // c is a rune
		fmt.Println(i, c)
	}

	c := map[string]int{
		"a": 1, "b": 2,
	}
	for k, v := range c { // i is a string (keys)
		fmt.Println(k, v)
	}
	// END OMIT
}
