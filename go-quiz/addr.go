package main

import (
	"fmt"
	"unsafe"
)

type Message struct {
	c1 byte
	c2 byte
	i  int64
}

func main() {
	var m Message
	addr1 := uintptr(unsafe.Pointer(&m.i))
	addr2 := uintptr(unsafe.Pointer(&m.c1))
	fmt.Println(addr1 - addr2)
}
