package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3}
	b := append(a[:1], 10)
	fmt.Println(a, b)
}
