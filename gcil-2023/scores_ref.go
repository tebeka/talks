package main

import (
	"fmt"
)

type Player struct {
	Name   string
	Points int
}

// START OMIT
func main() {
	players := []Player{
		{"Rick", 1_000_000},
		{"Morty", 13},
	}

	for i := range players { // HL
		players[i].Points += 353
	}
	fmt.Printf("%v\n", players)
	// [{Rick 1000353} {Morty 366}]
}

// END OMIT
