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

	for _, player := range players { // HL
		player.points += 353
	}
	fmt.Printf("%v\n", players)
	// [{Rick 1000000} {Morty 13}]
}

// END OMIT