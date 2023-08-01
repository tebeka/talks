package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

var text = `
Monday's child is fair of face,
Tuesday's child is full of grace.
Wednesday's child is full of woe,
Thursday's child has far to go.
Friday's child is loving and giving,
Saturday's child works hard for a living.
And the child born on the Sabbath day
Is bonny and blithe, good and gay.
`

func main() {

	r := strings.NewReader(text)
	term := "full"

	lnum := 0
	s := bufio.NewScanner(r)
	for s.Scan() {
		lnum++
		if strings.Contains(s.Text(), term) {
			fmt.Printf("%d: %s\n", lnum, s.Text())
		}
	}
	if err := s.Err(); err != nil {
		log.Fatalf("error: %s", err)
	}
}
