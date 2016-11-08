package main

import "fmt"

// #include "callback.h"
import "C"

//export go_func
func go_func() {
	fmt.Println("Hello from Go")
}

func main() {
	C.c_func();
}

