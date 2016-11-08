package main

// #include <string.h>
// #include <stdlib.h>
import "C"
import "fmt"
import "unsafe"

func free(cstr *C.char) {
	C.free(unsafe.Pointer(cstr))
}

func main() {
	str := "Gophers Rock!"
	cstr:= C.CString(str)
	defer free(cstr)

	dup := C.strdup(C.CString(str))
	defer free(dup)

	gostr := C.GoString(dup)
	fmt.Println(gostr)
}
