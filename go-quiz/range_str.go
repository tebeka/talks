// https://twitter.com/classic_addetz/status/1662769796692299780
package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Who’s "
		ch <- "on "
		ch <- "first?"
	}()

	for s := range <-ch {
		fmt.Print(s, " ")
	}
	fmt.Println()
}
