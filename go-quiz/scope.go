package main

import (
	"fmt"
)

func main() {
	n := 1
	{
		n := n + 1
		fmt.Printf("%d ", n)
	}
	fmt.Println(n)
}
