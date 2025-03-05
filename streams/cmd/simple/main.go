package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"logs/parser"
)

func main() {
	var logs []parser.Log
	nErr := 0

	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		s := bufio.NewScanner(file)
		for s.Scan() {
			if log, err := parser.ParseLine(s.Text()); err != nil {
				nErr++
				continue
			} else {
				logs = append(logs, log)
			}
		}
		if err := s.Err(); err != nil {
			return err
		}

		return nil
	}

	start := time.Now()
	if err := filepath.Walk("logs", walkFn); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	duration := time.Since(start)
	fmt.Printf("%d logs (%d errors) in %v\n", len(logs), nErr, duration)
}
