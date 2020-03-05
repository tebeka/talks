package main

import (
	"fmt"
	"reflect"
)

func alloc(size int) []int {
	return make([]int, size)
}

func main() {
	var expected []int
	out := alloc(0)
	fmt.Println(reflect.DeepEqual(expected, out))

}
