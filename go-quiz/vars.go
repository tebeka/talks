package main

import "fmt"

func main() {
	a, b := 1, 2
	b, a := 3, 4
	fmt.Println(a, b)
}
