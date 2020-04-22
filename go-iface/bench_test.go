package main

import (
	"testing"
)

type Point struct {
	X int
	Y int
}

func (p *Point) Move(dx, dy int) {
	p.X += dx
	p.Y += dy
}

type Mover interface {
	Move(dx, dy int)
}

var (
	dx = 3
	dy = 5
)

func BenchmarkStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := Point{0, 0}
		p.Move(dx, dy)
		if p.X < dx {
			b.Fatal(p.X)
		}
	}
}

func BenchmarkInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := Point{0, 0}
		var ip Mover = &p
		ip.Move(dx, dy)
		if p.X < dx {
			b.Fatal(p.X)
		}
	}
}
