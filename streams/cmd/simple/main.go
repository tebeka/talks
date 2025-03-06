package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"

	"logs/parser"
)

func main() {
	var logs []parser.Log

	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			slog.Warn("open", "path", path, "error", err)
			return nil
		}
		defer file.Close()

		s := bufio.NewScanner(file)
		for s.Scan() {
			if log, err := parser.ParseLine(s.Text()); err != nil {
				slog.Warn("parse", "path", path, "text", s.Text(), "error", err)
				continue
			} else {
				logs = append(logs, log)
			}
		}
		if err := s.Err(); err != nil {
			slog.Warn("scan", "error", err)
			return nil
		}

		return nil
	}

	if err := filepath.Walk("logs", walkFn); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	alloc_mb := mem.Alloc / (1 << 20)

	fmt.Printf("%d logs (%dmb)\n", len(logs), alloc_mb)
}
