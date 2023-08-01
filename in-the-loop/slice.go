package main

import "fmt"

func main() {
	cart := []string{"bread", "butter", "beer"}
	// indices
	for i := range cart {
		fmt.Println(i)
	}
	// index + value
	for i, v := range cart {
		fmt.Println(i, v)
	}
	// values
	for _, v := range cart {
		fmt.Println(v)
	}
}
