package main

import (
	"fmt"
)

type Tree struct {
	value       int
	left, right *Tree
}

func walk(tree *Tree, ch chan int) {
	if tree == nil {
		return
	}
	ch <- tree.value
	walk(tree.left, ch)
	walk(tree.right, ch)
}

func leaves(tree *Tree) chan int {
	ch := make(chan int)
	go func() {
		walk(tree, ch)
		close(ch)
	}()
	return ch
}

func main() {
	tree := &Tree{1, &Tree{2, nil, nil}, &Tree{3, nil, nil}}
	for v := range leaves(tree) {
		fmt.Println(v)
	}
}
