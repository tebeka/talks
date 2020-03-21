package main

import (
	"unsafe"
)

/*
#include <stdio.h>

void show(char *msg) {
	printf("C says: %s\n", msg);
}
*/
import "C"

func main() {
	msg := "Go is better"
	C.show((*C.char)(unsafe.Pointer(&msg)))
}
