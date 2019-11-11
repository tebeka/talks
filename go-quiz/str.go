package main

import "fmt"

func main() {
	city := "Kraków"
	for c := range city {
		fmt.Printf("%v ", c)
	}
	fmt.Println()
}
