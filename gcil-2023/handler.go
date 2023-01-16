package main

import (
	"fmt"
)

type Message struct {
	User string
	Body string
}

// AAA OMIT
type Provider interface {
	// Next return next message, nil when done
	Next() *Message
}

// BBB OMIT

// START OMIT
func handler(p Provider) {
	for { // HL
		msg := p.Next()
		if msg == nil { // HL
			break
		}
		fmt.Printf("%+v\n", msg)
	}
}

// END OMIT

type prov struct {
	i     int
	count int
}

func (p *prov) Next() *Message {
	if p.i == p.count {
		return nil
	}
	p.i++

	msgs := []string{
		":(){ :|:& };:",
		"sudo rm -rf /",
		"sudo shutdown -h now",
	}

	return &Message{
		User: "eliot",
		Body: msgs[p.i-1],
	}
}

func main() {
	p := prov{count: 3}
	handler(&p)
}
