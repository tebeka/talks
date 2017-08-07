package main

import (
	"fmt"
	"math"
)

// START OMIT
type Point struct {
	X float64
	Y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func (pt *Point) Abs() float64 {
	return math.Sqrt(pt.X*pt.X + pt.Y*pt.Y)
}

func main() {
	pt := NewPoint(4, 3)
	fmt.Println(pt.Abs())
}

// END OMIT
