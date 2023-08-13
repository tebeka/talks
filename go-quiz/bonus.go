package main

import (
	"fmt"
)

func main() {
	var players = []struct {
		name   string
		points int
	}{
		{"Wednesday", 1_000_000},
		{"Pugsley", 13},
	}

	for _, player := range players {
		player.points += 123
	}
	fmt.Printf("%v\n", players)
}
