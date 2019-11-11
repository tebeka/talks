package main

import "fmt"

func main() {
	city := "Kraków"
	for _, c := range city {
		fmt.Printf("%T\n", c)
		break
	}
}
