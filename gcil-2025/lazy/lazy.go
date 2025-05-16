package lazy

import (
	"bufio"
	"encoding/json"
	"io"
	"iter"
	"log/slog"
	"mordor/log"
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

// Lines return a sequence of lines in r.
func Lines(r io.Reader) iter.Seq[string] {
	fn := func(yield func(string) bool) {
		s := bufio.NewScanner(r)
		for s.Scan() {
			line := s.Text()
			if !yield(line) {
				return
			}
		}
		if err := s.Err(); err != nil {
			slog.Warn("scan", "error", err)
			return
		}
	}

	return fn
}

// parseLine parser a log line, on error it'll return empty log.
func parseLine(line string) log.Log {
	var record log.Log
	if err := json.Unmarshal([]byte(line), &record); err != nil {
		return log.Log{}
	}

	return record
}

// LoadLogs return a sequence of logs from file.
func LoadLogs(r io.Reader) iter.Seq[log.Log] {
	lines := Lines(r)
	logs := Map(lines, parseLine)
	logs = Filter(logs, log.IsValid)

	return logs
}
