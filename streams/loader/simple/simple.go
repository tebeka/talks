package simple

import (
	"bufio"
	"log/slog"
	"os"
	"path/filepath"

	"logs/parser"
)

// LoadLogs return a slice of logs from all files under root.
func LoadLogs(root string) ([]parser.Log, error) {
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

	if err := filepath.Walk(root, walkFn); err != nil {
		return nil, err
	}

	return logs, nil
}
