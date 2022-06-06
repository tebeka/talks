package main

import (
	"io"
	"os"
	"strings"
)

// START_ROT13 OMIT
type Rot13Decoder struct {
	r io.Reader
}

func (r *Rot13Decoder) Read(buf []byte) (int, error) { // HL
	n, err := r.r.Read(buf)
	if err != nil {
		return n, err
	}

	for i, b := range buf[:n] {
		switch {
		case b >= 'A' && b <= 'Z':
			b = ((b-'A')+13)%26 + 'A'
			buf[i] = b
		case b >= 'a' && b <= 'z':
			b = ((b-'a')+13)%26 + 'a'
			buf[i] = b
		}
	}

	return n, nil
}

// END_ROT13 OMIT

func main() {
	// START_MAIN OMIT
	r := strings.NewReader("Neqna Ynof")
	io.Copy(os.Stdout, &Rot13Decoder{r})
	// END_MAIN OMIT
}
