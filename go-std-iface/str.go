package main

import (
	"bytes"
	"fmt"
)

// START_LIST OMIT
type List struct {
	Value string
	Next  *List
}

func (l *List) String() string { // HL
	var buf bytes.Buffer
	for node := l; node != nil; node = node.Next {
		fmt.Fprintf(&buf, "[%s]", node.Value)
		if node.Next != nil {
			fmt.Fprint(&buf, "->")
		}
	}
	return buf.String()
}

// END_LIST OMIT

func main() {
	// START_MAIN OMIT
	l := &List{"Who's", &List{"on", &List{"first?", nil}}}
	fmt.Println(l)
	// END_MAIN OMIT
}
