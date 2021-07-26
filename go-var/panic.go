package main

import (
	"fmt"
)

// padding
// padding

// START OMIT
func safeDiv(a, b int) (result int, err error) { // HL
	defer func() {
		// err will be non-nil only if there's a panic
		if e := recover(); e != nil {
			err = fmt.Errorf("div error: %s\n", e) // HL
		}
	}()

	return a / b, nil
}

// END OMIT

func main() {
	// MAIN_START OMIT
	fmt.Println(safeDiv(7, 3))
	fmt.Println(safeDiv(7, 0))
	// MAIN_END OMIT
}
