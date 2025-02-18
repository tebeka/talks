package main

import "fmt"

func main() {
	// START OMIT
	cart := []string{"bread", "butter", "beer"}

	// indices
	for i := range cart { // HL
		fmt.Println(i)
	}

	// index + value
	for i, v := range cart { // HL
		fmt.Println(i, v)
	}

	// values
	for _, v := range cart { // HL
		fmt.Println(v)
	}

	// END OMIT

}
