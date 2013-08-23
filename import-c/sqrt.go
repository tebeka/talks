package main

import "fmt"

// #include <math.h>
// #cgo LDFLAGS: -lm
import "C"


func main() {
	v := 16.0
	s := C.sqrt(C.double(v))
	fmt.Printf("sqrt(%f) = %f\n", v, s)
}
