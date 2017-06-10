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

func (p *Point) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func main() {
	p := NewPoint(4, 3)
	fmt.Println(p.Abs())
}

// END OMIT
