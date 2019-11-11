package main

import "fmt"

func main() {
	var a, b *int

	a = new(int)
	*a = 8
	b = a
	*a++
	fmt.Println(*a, *b)
}
