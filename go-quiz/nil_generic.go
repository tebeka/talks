package main

import "fmt"

func Max[T int | float64](values []T) T {
	m := values[0]
	for _, v := range values[1:] {
		if v > m {
			m = v
		}
	}
	return m
}

func main() {
	fmt.Println(Max(nil))
}
