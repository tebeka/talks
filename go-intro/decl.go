package main

import (
	"fmt"
)

// START_CONST OMIT
const (
	message = "Hello"
	pi      = 3.14159265358
)

// END_CONST OMIT

func main() {
	// START_MAIN OMIT
	var i int // Initialized to 0

	i = 7
	j := 1 // j is also an int

	a := []int{1, 2, 3}  // slice of ints
	m := map[string]int{ // map string -> int
		"a": 1, "b": 2,
	}

	type Point struct {
		X, Y int
	}

	p1 := new(Point)
	p2 := &Point{1, 2}
	p3 := &Point{Y: 1, X: 2}
	// END_MAIN OMIT

	fmt.Println("message =", message)
	fmt.Println("pi =", pi)
	fmt.Println("i =", i)
	fmt.Println("j =", j)
	fmt.Println("a = ", a)
	fmt.Println("m = ", m)
	fmt.Println("p1 =", p1)
	fmt.Println("p2 =", p2)
	fmt.Println("p3 =", p3)
}
