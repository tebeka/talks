package main

import (
	"fmt"
	"strings"
	"time"
)

// Player is a player in the 7boom game
type Player struct {
	ID   int
	in   chan int
	out  chan int
	done chan bool
}

// NewPlayer return a new player
func NewPlayer(id int, done chan bool) *Player {
	return &Player{
		ID:   id,
		out:  make(chan int, 1),
		done: done,
	}
}

// isBoom returns true if n is multiple of 7 or has a 7 digit
func isBoom(n int) bool {
	if n%7 == 0 {
		return true
	}

	if strings.Contains(fmt.Sprintf("%d", n), "7") {
		return true
	}

	return false
}

// Play is the player game loop
func (p *Player) Play() {
	for {
		select {
		case v := <-p.in:
			// Simulate work
			time.Sleep(1 * time.Second)
			if isBoom(v) {
				fmt.Printf("Player %d: BOOM\n", p.ID)
			} else {
				fmt.Printf("Player %d: %d\n", p.ID, v)
			}
			p.out <- v + 1
		case <-p.done:
			fmt.Printf("Player %d: QUIT\n", p.ID)
			return
		}
	}
}

// makeChain creates a chain of players, return the 1st player and done channel
// It will also invode the Play method of each player in a gouroutine
func makeChain(n int) (*Player, chan bool) {
	var prev *Player
	var first *Player
	done := make(chan bool)

	// Create chain of players
	for i := 0; i < n; i++ {
		player := NewPlayer(i, done)
		if prev != nil {
			player.in = prev.out
		}

		if first == nil {
			first = player
		} else {
			go player.Play()
		}
		prev = player
	}
	first.in = prev.out
	go first.Play()
	return first, done
}

func main() {
	first, done := makeChain(3)

	fmt.Println("Play time!")
	first.in <- 1
	time.Sleep(70 * time.Second) // Let them play

	fmt.Println("Stopping Game")
	close(done)
	time.Sleep(200 * time.Millisecond) // for QUIT prints
}
