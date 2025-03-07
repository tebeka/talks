package seq

import (
	"bufio"
	"fmt"
	"iter"
	"log/slog"
	"os"
	"path/filepath"

	"logs/parser"
)

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

		if err := filepath.Walk(root, walkFn); err != nil {
			slog.Error("walk", "root", root, "error", err)
		}
	}

	return fn, nil
}
