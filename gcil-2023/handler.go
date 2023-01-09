package main

type Message struct {
	// ...
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
		if msg == nil { // no more
			break
		}
		// TODO: Handle msg
	}
}

// END OMIT

func main() {
}
