package mapper

import (
	"iter"
	"log/slog"
	"os"

	"logs/loader/seq"
	"logs/parser"
)

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

// LogLines return a sequence of lines from all files found in root.
func LogLines(root string) iter.Seq[string] {
	fn := func(yield func(string) bool) {
		for fileName := range seq.IterFiles(root) {
			file, err := os.Open(fileName)
			if err != nil {
				slog.Warn("open", "path", fileName, "error", err)
				continue
			}
			defer file.Close()

			for line := range seq.IterLines(file) {
				if !yield(line) {
					return
				}
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

// LoadLogs return a sequence of logs from all files under root.
func LoadLogs(root string) (iter.Seq[parser.Log], error) {
	lines := LogLines(root)
	seq := Map(lines, parseLine)
	seq = Filter(seq, func(log parser.Log) bool {
		return !log.Time.IsZero()
	})

	return seq, nil
}
