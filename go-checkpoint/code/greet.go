package main

import "fmt"

type Greeter interface {
	Greet() string
}

type GI int

func (gi GI) Greet() string {
	return fmt.Sprintf("Hi from %d", gi)
}

type GF float64

func (gf GF) Greet() string {
	return fmt.Sprintf("Hi from %.2f", gf)
}

func main() {
	gs := []Greeter{GI(42), GF(2.718282)}
	for _, g := range gs {
		fmt.Println(g.Greet())
	}
}
