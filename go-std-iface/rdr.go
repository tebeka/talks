package main

// START_GO OMIT
type Reader interface {
	Read(p []byte) (int, error)
}

// END_GO OMIT

// START_PY OMIT
type Reader interface {
	Read(n int) ([]byte, error)
}

// END_PY OMIT
