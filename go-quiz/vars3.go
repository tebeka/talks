package main

import "fmt"

func main() {
	a, b := 3, 5
	a, b = a+b, a
	fmt.Println(a, b)
}
