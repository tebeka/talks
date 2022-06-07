package main

import "fmt"

// START_MAX OMIT
func MaxInts(values []int) (int, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("MaxInts of empty slice")
	}
	max := values[0]
	for _, v := range values[1:] {
		if v > max {
			max = v
		}
	}
	return max, nil
}

func MaxFloats(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("MaxFloats of empty slice")
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
