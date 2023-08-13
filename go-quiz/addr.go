package main

import (
	"fmt"
	"unsafe"
)

type Message struct {
	Level   byte
	Group   byte
	ID      int
	Content string
}

func main() {
	var m Message
	addr1 := uintptr(unsafe.Pointer(&m.ID))
	addr2 := uintptr(unsafe.Pointer(&m.Group))
	fmt.Println(addr1 - addr2)
}
