package main

import "fmt"

func main() {
	greeting := "sup?"
	for range greeting {
		fmt.Printf("x")
	}

	fmt.Println()
}
