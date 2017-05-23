package main

import (
	"encoding/binary"
	"fmt"
	"unsafe"
)

func main() {
	val := uint32(987654321)

	fmt.Println("Traditional way")
	buf := make([]byte, 16)
	binary.LittleEndian.PutUint32(buf[8:], val)
	fmt.Println(buf)

	fmt.Println("Unsafe way")
	buf = make([]byte, 16)
	p := (*uint32)(unsafe.Pointer(&buf[8]))
	*p = val
	fmt.Println(buf)
}
