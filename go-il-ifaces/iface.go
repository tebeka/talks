package main

import "unsafe"

// START_IFACE OMIT
type iface struct {
	tab  *itab
	data unsafe.Pointer
}

// END_IFACE OMIT
