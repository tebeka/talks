package main

import (
	"cmp"
	"errors"
	"fmt"
	"iter"
)

type node[T any] struct {
	value T
	next  *node[T]
}

type Stack[T any] struct {
	head *node[T]
}

func (s *Stack[T]) Push(v T) {
	s.head = &node[T]{v, s.head}
}

var ErrEmpty = errors.New("empty stack")

func (s *Stack[T]) Pop() (T, error) {
	if s.head == nil {
		var v T
		return v, ErrEmpty
	}

	n := s.head
	s.head = s.head.next
	return n.value, nil
}

func (s *Stack[T]) Items() *StackIterator[T] {
	return &StackIterator[T]{s.head}
}

type StackIterator[T any] struct {
	node *node[T]
}

func (s *StackIterator[T]) Next() (T, bool) {
	if s.node == nil {
		var v T
		return v, false
	}

	n := s.node
	s.node = s.node.next
	return n.value, true
}

func (s *Stack[T]) PrintItems() {
	for n := s.head; n != nil; n = n.next {
		fmt.Println(n.value)
	}
}

func (s *Stack[T]) Do(yield func(v T)) {
	for n := s.head; n != nil; n = n.next {
		yield(n.value)
	}
}

func (s *Stack[T]) Iter() func(func(T) bool) {
	iter := func(yield func(T) bool) { // HL
		for n := s.head; n != nil; n = n.next {
			if !yield(n.value) { // HL
				return
			}
		}
	}

	return iter
}

func (s *Stack[T]) Iter2() func(func(int, T) bool) {
	iter := func(yield func(int, T) bool) { // HL
		for i, n := 0, s.head; n != nil; i, n = i+1, n.next {
			if !yield(i, n.value) { // HL
				return
			}
		}
	}

	return iter
}

func Max[T cmp.Ordered](seq iter.Seq[T]) (T, error) {
	pull, stop := iter.Pull(seq)
	defer stop() // HL

	max, ok := pull()
	if !ok {
		return max, fmt.Errorf("Max of empty sequence")
	}

	for v, ok := pull(); ok; v, ok = pull() { // HL
		if v > max {
			max = v
		}
	}

	return max, nil
}

func main() {
	// START_S OMIT
	var s Stack[int]
	s.Push(10)
	s.Push(20)
	s.Push(30)
	// END_S OMIT

	fmt.Println(s.Pop()) // [
	fmt.Println(s.Pop()) // (
	fmt.Println(s.Pop()) // {
	fmt.Println(s.Pop()) // empty stack

	fmt.Println("PrintItems")
	s.Push(10)
	s.Push(20)
	s.Push(30)

	s.PrintItems()

	fmt.Println("Items")
	// START_ITER OMIT
	it := s.Items()
	for v, ok := it.Next(); ok; v, ok = it.Next() {
		fmt.Println(v)
	}
	// END_ITER OMIT

	fmt.Println("Do")

	// START_DO OMIT
	s.Do(func(n int) {
		fmt.Println(n)
	})
	// END_DO OMIT

	fmt.Println("for")

	// START_FOR OMIT
	for v := range s.Iter() {
		fmt.Println(v)
	}
	// END_FOR OMIT

	fmt.Println("for2")
	// START_ITER2 OMIT
	for i, v := range s.Iter2() {
		fmt.Println(i, v)
	}
	// END_ITER2 OMIT

	// START_MAX OMIT
	m, err := Max(s.Iter())
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Println("max:", m)
	}
	// END_MAX OMIT

	fmt.Println("seq(func")
	seq := s.Iter()
	seq(func(v int) bool {
		fmt.Println(v)
		return true
	})

	// START_2 OMIT
	for i, c := range "Hi" {
		fmt.Printf("%c at %d\n", c, i)
	}
	// END_2 OMIT
}
