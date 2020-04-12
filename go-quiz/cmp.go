package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

func main() {
	p1 := Point{1, 2}
	p2 := Point{2, 1}
	fmt.Println(p1 > p2)
}
