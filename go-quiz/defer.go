package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		defer fmt.Printf("%d ", i)
	}
}
