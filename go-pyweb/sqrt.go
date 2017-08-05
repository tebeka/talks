// go build -buildmode=c-shared -o _sqrt.so sqrt.go
package main

import "math"

import "C"

//export sqrt
func sqrt(n float64) float64 {
	// Newton's square root algorithm
	guess := 1.0
	epsilon := 0.0001

	for math.Abs(guess*guess-n) > epsilon {
		guess = (guess + n/guess) / 2.0
	}

	return guess
}

func main() {}
