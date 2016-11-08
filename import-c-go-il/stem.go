package main

import (
	"fmt"
	"bitbucket.org/tebeka/snowball"
)

func main() {
	stemmer, _ := snowball.New("english") // Error ignored for brevity

	word := "running"
	stem := stemmer.Stem(word)
	fmt.Printf("%s -> %s\n", word, stem)
}

