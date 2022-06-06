package main

// START_IFACE OMIT
type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// END_IFACE OMIT
