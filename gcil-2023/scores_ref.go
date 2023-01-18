package main

import (
	"fmt"
)

// START OMIT
func main() {
	var players = []struct {
		name   string
		points int
	}{
		{"Rick", 1_000_000},
		{"Morty", 13},
	}

	for i := range players { // HL
		players[i].points += 353
	}
	fmt.Printf("%v\n", players)
	// [{Rick 1000353} {Morty 366}]
}

// END OMIT
