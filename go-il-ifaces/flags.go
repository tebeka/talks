package main

import "fmt"

// START_FLAG OMIT
type Flag uint8

const (
	Read Flag = 1 << iota
	Write
	Execute
)

func (f Flag) String() string { // HL
	switch f {
	case Read:
		return "read"
	case Write:
		return "write"
	case Execute:
		return "execute"
	default:
		return fmt.Sprintf("<Flag %d>", f) // HL
	}
	// Exercise: Support Read|Write ☺
}

// END_FLAG OMIT

func main() {
	// START_MAIN OMIT
	fmt.Println(Execute)
	fmt.Printf("s: %s\n", Execute)
	fmt.Printf("v: %v\n", Execute)
	fmt.Printf("+v: %+v\n", Execute)
	fmt.Printf("#v: %#v\n", Execute)
	fmt.Printf("d: %d\n", Execute)
	// END_MAIN OMIT
}
