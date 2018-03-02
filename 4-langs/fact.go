package main

import "fmt"

func fact(n int) int {
	total := 1
	for n > 0 {
		total *= n
		n--
	}
	return total
}

func main() {
	fmt.Println(fact(10))
}
