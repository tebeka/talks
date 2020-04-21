package main

import (
	"fmt"
	"math"
	"sort"
)

// Point is a 2D point
type Point struct {
	X float64
	Y float64
}

func (p Point) distFrom0() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func main() {
	points := []Point{
		Point{1, 2},
		Point{2, 3},
		Point{1, 7},
		Point{2, 2},
	}
	less := func(i, j int) bool {
		return points[i].distFrom0() < points[j].distFrom0()
	}
	sort.Slice(points, less)
	for _, p := range points {
		fmt.Printf("%+v → %.2f\n", p, p.distFrom0())
	}
}
