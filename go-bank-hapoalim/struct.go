package main

import (
	"fmt"
)

// Point is a 2D point
// Point will be exported from the package
type Point struct {
	X int // X is exported struct field
	Y int
}

// NewPoint return a *Point and possible error if x or y are out of bounds
func NewPoint(x, y int) (*Point, error) {
	if x > 1000 || y > 1000 || x < 0 || y < 0 {
		return nil, fmt.Errorf("(%d, %d) out of bounds [0,1000]", x, y)
	}
	pt := &Point{
		X: x,
		Y: y,
	}
	return pt, nil // Go's compiler does escape analysis and will allocate pt on the heap
}

// p is the method receiver
// We use mostly pointer receivers, especially when mutating
func (p *Point) Move(dx, dy int) {
	fmt.Printf("moving: %d, %d\n", dx, dy)
	p.X += dx
	p.Y += dy
}

func main() {
	p := Point{1, 2}
	fmt.Printf("%+v\n", p)

	p.Move(10, 20)
	fmt.Printf("%+v\n", p)
}
