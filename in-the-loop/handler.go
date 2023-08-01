package main

import (
	"fmt"
)

func main() {
	p := provider{max: 5}
	handler(&p)
}

func handler(p Provider) {
	for {
		msg := p.Next()
		if msg == nil {
			break
		}

		fmt.Printf("%+v\n", msg)
	}
}

type Message struct {
	ID      int
	Content string
}

type Provider interface {
	Next() *Message
}

type provider struct {
	max int
	n   int
}

func (p *provider) Next() *Message {
	if p.n == p.max {
		return nil
	}
	p.n++
	return &Message{
		ID:      p.n,
		Content: fmt.Sprintf("message #%d", p.n),
	}

}
