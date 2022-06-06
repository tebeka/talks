package main

import "io"

// START_WS OMIT
type WriteSyncer interface {
	io.Writer
	Sync() error
}

type Logger struct {
	level Level
	w     WriteSyncer // HL
}

// END_WS OMIT
