package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if err := demo(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

}

func demo() error {
	// START OMIT
	cmd := exec.Command("sleep", "1000")
	if err := cmd.Start(); err != nil {
		return err
	}

	fmt.Println("PID   :", cmd.Process.Pid) // HL
	cmd.Process.WithHandle(func(handle uintptr) {
		fmt.Println("HANDLE:", handle) // HL
	})
	// END OMIT

	return nil
}
