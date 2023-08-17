package main

import (
	"fmt"
)

func main() {
	n := 1
	{
		n := n + 1
		fmt.Print(n, " ")
	}
	fmt.Println(n)
}
