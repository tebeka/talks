package main

import (
	"runtime"
	"runtime/debug"
)

func call() {
	debug.PrintStack()
	runtime.Breakpoint()
}

func main() {
	call()
}
