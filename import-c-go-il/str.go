package main

// #include <string.h>
import "C"

import (
	"fmt"
)

func main() {
	str := "Gophers Rock!"
	dup := C.strdup(C.CString(str))
	fmt.Printf("dup: %s\n", dup)
}

