package main

import (
	"fmt"
	"io"
	"os"
)

func Open(path string) io.Writer {
	var f *os.File
	return f // FIXME
}

func main() {
	f := Open("/dev/zero")
	fmt.Println(f == nil)
}
