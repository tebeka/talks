package main

import (
	"fmt"
)

// padding

func main() {
	// START OMIT
	x := 1
	{
		x := 7 // HL
		fmt.Println("inner", x)

	}
	fmt.Println("outer", x) // 2
	// END OMIT
}
