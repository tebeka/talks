package main

import (
	//	"fmt"
	"os"
)

func printQuote(author, quote string) {
	os.Stdout.Write([]byte(author))
	os.Stdout.Write([]byte{'\n'})
	os.Stdout.Write([]byte(quote))
}

func main() {
	author := "Heinlein"
	quote := "When one teaches, two learn."
	printQuote(author, quote)
	//fmt.Printf("%s said: %s\n", author, quote)
}
