package main

import (
	"fmt"
)

const (
	RequestType = 1 << iota
	ReplyType
)

type Reply struct{}

func (r *Reply) Type() byte {
	return ReplyType
}

func main() {
	var r *Reply
	fmt.Println(r.Type())
}
