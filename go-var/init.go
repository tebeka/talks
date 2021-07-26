package main

import (
	"fmt"
)

// START OMIT
func init() {
	fmt.Printf("A ")
}

func init() {
	fmt.Print("B ")
}

// END OMIT

func main() {
	fmt.Println()
}
