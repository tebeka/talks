package main

// #include <string.h>
// #include <stdlib.h>
import "C"

import (
	"fmt"
)

func main() {
	str := "Gophers Rock!"
	dup := C.strdup(C.CString(str))

	gostr := C.GoString(dup)
	fmt.Println(gostr)
}
