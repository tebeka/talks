package main

import "fmt"

// START_MAX OMIT
type Number interface { // HL
	float64 | int // HL
} // HL

func Max[T Number](values []T) (T, error) { // HL
	if len(values) == 0 {
		var zero T // HL
		return zero, fmt.Errorf("Max of empty slice")
	}

	max := values[0]
	for _, v := range values[1:] {
		if v > max {
			max = v
		}
	}
	return max, nil
}

// END_MAX OMIT

func main() {
	// START_MAIN OMIT
	iVals := []int{2, 7, 1, 8}
	fmt.Println(Max(iVals))

	fVals := []float64{2, 7, 1, 8}
	fmt.Println(Max(fVals))
	// END_MAIN OMIT
}
