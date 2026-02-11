package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// START OMIT
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGUSR1) // HL
	defer cancel()

	go func() {
		proc, _ := os.FindProcess(os.Getpid())
		proc.Signal(syscall.SIGUSR1) // HL
	}()

	<-ctx.Done()
	fmt.Println("ERROR:", ctx.Err())
	fmt.Println("CAUSE:", context.Cause(ctx))
	// END OMIT
}
