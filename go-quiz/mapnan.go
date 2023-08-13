package main

import (
	"fmt"
	"math"
)

func main() {
	m := map[float64]string{
		math.NaN(): "MacLeod",
	}
	delete(m, math.NaN())
	fmt.Println(m)
}
