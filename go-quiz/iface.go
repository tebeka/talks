package main

import (
	"fmt"
	"unsafe"
)

type Writer interface {
	Write()
}

type ReadWriter interface {
	Read()
	Writer
}

type File struct{}

func (File) Read()  { fmt.Println("Read") }
func (File) Write() { fmt.Println("Write") }

func main() {
	var rw ReadWriter = File{}
	w := *(*Writer)(unsafe.Pointer(&rw))
	w.Write()
}
