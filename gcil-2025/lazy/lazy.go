package lazy

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"iter"
	"log/slog"
	"mordor/log"
	"os"
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

// Lines return a sequence of lines file.
func Lines(fileName string) iter.Seq[string] {
	fn := func(yield func(string) bool) {
		file, err := os.Open(fileName)
		if err != nil {
			slog.Warn("open", "path", fileName, "error", err)
			return
		}
		defer file.Close()

		gz, err := gzip.NewReader(file)
		if err != nil {
			slog.Warn("gzip", "path", fileName, "error", err)
			return
		}
		defer gz.Close()

		s := bufio.NewScanner(gz)
		for s.Scan() {
			line := s.Text()
			if !yield(line) {
				return
			}
		}
		if err := s.Err(); err != nil {
			slog.Warn("scan", "path", fileName, "error", err)
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
func LoadLogs(fileName string) (iter.Seq[log.Log], error) {
	lines := Lines(fileName)
	logs := Map(lines, parseLine)
	logs = Filter(logs, log.IsValid)

	return logs, nil
}
