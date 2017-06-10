package main

import (
	"C"
)

//export add
func add(a, b int64) int64 {
	return a + b
}

//export sub
func sub(a, b int64) int64 {
	return a - b
}

func main() {}
