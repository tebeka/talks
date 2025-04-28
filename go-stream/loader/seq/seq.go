package seq

import (
	"bufio"
	"io"
	"iter"
	"log/slog"
	"os"
	"path/filepath"

	"logs/parser"
)

func Lines(r io.Reader) iter.Seq[string] {
	fn := func(yield func(string) bool) {
		s := bufio.NewScanner(r)
		for s.Scan() {
			if !yield(s.Text()) {
				return
			}
		}

		if err := s.Err(); err != nil {
			slog.Error("scan", "error", err)
			return
		}
	}

	return fn
}

// LoadLogs return a sequence of logs from all files under root.
func LoadLogs(root string) (iter.Seq[parser.Log], error) {
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

			for line := range Lines(file) {
				log, err := parser.ParseLine(line)
				if err != nil {
					slog.Warn("parse", "path", path, "text", line, "error", err)
					continue
				}

				if !yield(log) {
					return filepath.SkipAll
				}
			}

			return nil
		}

		if err := filepath.Walk(root, walkFn); err != nil {
			slog.Error("walk", "root", root, "error", err)
		}
	}

	return fn, nil
}
