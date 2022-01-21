package main

import "fmt"

func main() {
	s := "«Serialization»"
	count := 0
	for _, r := range s {
		fmt.Printf("%X", r)
		count++
	}
	fmt.Println()
	fmt.Println(count, len(s))
}
