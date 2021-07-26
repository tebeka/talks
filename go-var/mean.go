package main

import "fmt"

// START_SCALE OMIT
var scale = 1.1

func scaled_mean(values []float64) float64 {
	sum := 0.0
	for _, val := range values {
		sum += val
	}

	// START_RET OMIT
	return sum / float64(len(values)) * scale // HL
	// END_RET OMIT
}

// END_SCALE OMIT

func main() {
	// START_MAIN OMIT
	values := []float64{1, 1, 2, 3, 5, 8, 13, 21}
	fmt.Println(scaled_mean(values))
	// END_MAIN OMIT
}
