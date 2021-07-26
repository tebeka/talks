package main

import "unsafe"

// START OMIT
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

// END OMIT
