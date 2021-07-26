package main

import (
	"fmt"
	"log"
	"sort"
)

// MEDIAN_START OMIT
func median(nums []float64) (float64, error) {
	if len(nums) == 0 {
		return 0, fmt.Errorf("median of no numbers")
	}

	sort.Float64s(nums)
	mid := len(nums) / 2
	if len(nums)%2 == 1 {
		return nums[mid], nil
	}

	return (nums[mid-1] + nums[mid]) / 2, nil
}

// MEDIAN_END OMIT

func main() {
	// MAIN_START OMIT
	nums := []float64{3, 1, 4, 1, 5, 9, 2, 6}
	m, err := median(nums)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(nums, "->", m)
	// MAIN_END OMIT
}
