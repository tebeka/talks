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

// IterLines return an iterator of lines from the reader.
func IterLines(r io.Reader) iter.Seq[string] {
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

// IterFiles return an iterator of files under root.
func IterFiles(root string) iter.Seq[string] {
	fn := func(yield func(string) bool) {
		walkFn := func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			if !yield(path) {
				return filepath.SkipAll
			}

			return nil
		}

		if err := filepath.Walk(root, walkFn); err != nil {
			slog.Error("walk", "root", root, "error", err)
		}
	}

	return fn
}

// LoadLogs return a sequence of logs from all files under root.
func LoadLogs(root string) (iter.Seq[parser.Log], error) {
	fn := func(yield func(parser.Log) bool) {
		for fileName := range IterFiles(root) {
			file, err := os.Open(fileName)
			if err != nil {
				slog.Warn("open", "path", fileName, "error", err)
				continue
			}
			defer file.Close()

			for line := range IterLines(file) {
				log, err := parser.ParseLine(line)
				if err != nil {
					slog.Warn("parse", "path", fileName, "text", line, "error", err)
					continue
				}

				if !yield(log) {
					return
				}
			}
		}
	}

	return fn, nil
}
