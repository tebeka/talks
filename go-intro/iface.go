package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

// START_IFACE OMIT
type Abser interface {
	Abs() float64
}

type Floater float64

func (f Floater) Abs() float64 {
	v := float64(f) // cast for return type
	if v < 0 {
		return -v
	}
	return v
}

func printAbser(a Abser) {
	fmt.Println(a.Abs())
}

// END_IFACE OMIT

func main() {
	// START_MAIN OMIT
	p := &Point{4, 3}
	f := Floater(-2)

	printAbser(p)
	printAbser(f)
	// END_MAIN OMIT
}
