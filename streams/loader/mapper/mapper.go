package mapper

import (
	"bufio"
	"fmt"
	"iter"
	"log/slog"
	"os"
	"path/filepath"

	"logs/parser"
)

// IterLines return a sequence of lines from all files found in root.
func IterLines(root string) iter.Seq[string] {
	fn := func(yield func(string) bool) {
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
				if !yield(s.Text()) {
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

// Map transforms sequence of F to sequence of T.
func Map[F any, T any](seq iter.Seq[F], mapper func(F) T) iter.Seq[T] {
	fn := func(yield func(T) bool) {
		for v := range seq {
			t := mapper(v)
			if !yield(t) {
				return
			}
		}
	}

	return fn
}

// parseLine parser a log line, on error it'll return empty log.
func parseLine(line string) parser.Log {
	log, err := parser.ParseLine(line)
	if err != nil {
		return parser.Log{}
	}

	return log
}

// Filter return a sequence of items in seq that pred returned true.
func Filter[T any](seq iter.Seq[T], pred func(T) bool) iter.Seq[T] {
	fn := func(yield func(T) bool) {
		for v := range seq {
			if !pred(v) {
				continue
			}

			if !yield(v) {
				return
			}
		}
	}

	return fn
}

// LoadLogs return a sequence of logs from all files under root.
func LoadLogs(root string) (iter.Seq[parser.Log], error) {
	seq := Map(IterLines(root), parseLine)
	seq = Filter(seq, func(log parser.Log) bool {
		return !log.Time.IsZero()
	})

	return seq, nil
}
