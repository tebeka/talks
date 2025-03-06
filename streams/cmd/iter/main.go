package main

import (
	"bufio"
	"fmt"
	"iter"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"

	"logs/parser"
)

func iterLogs(root string) iter.Seq[parser.Log] {
	fn := func(yield func(parser.Log) bool) {
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
				log, err := parser.ParseLine(s.Text())
				if err != nil {
					slog.Warn("parse", "path", path, "text", s.Text(), "error", err)
					continue
				}
				if !yield(log) {
					return fmt.Errorf("stop")
				}
			}
			if err := s.Err(); err != nil {
				slog.Warn("scan", "error", err)
				return nil
			}

			return nil
		}

		filepath.Walk(root, walkFn)
	}

	return fn
}

func main() {

	n := 0
	for log := range iterLogs("logs") {
		_ = log
		n++
	}

	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	alloc_mb := mem.Alloc / (1 << 20)

	fmt.Printf("%d logs (%dmb)\n", n, alloc_mb)
}
