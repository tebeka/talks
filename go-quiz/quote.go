package main

import "fmt"

func main() {
	a, b := 'a', `b`
	fmt.Printf("%T %T\n", a, b)
}
