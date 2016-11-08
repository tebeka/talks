package main

// #include <math.h>
// #cgo LDFLAGS: -lm
import "C"

import (
	"fmt"
)

func main() {
	i, err := C.sqrt(-1)
	if err != nil {
		fmt.Printf("error: %s\n", err) // Go will use strerror for message
	}
	fmt.Printf("i = %v\n", i)
}
