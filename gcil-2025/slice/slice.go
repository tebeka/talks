package slice

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"log/slog"
	"os"
	"path/filepath"

	"mordor/log"
)

// LoadLogs return a slice of logs from all files under root.
func LoadLogs(root string) ([]log.Log, error) {
	var logs []log.Log

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

		gz, err := gzip.NewReader(file)
		if err != nil {
			slog.Warn("gzip", "path", path, "error", err)
			return nil
		}
		defer gz.Close()

		dec := json.NewDecoder(gz)
		lnum := 0
		for {
			lnum++
			var log log.Log
			err := dec.Decode(&log)

			if err == io.EOF {
				break
			}

			if err != nil {
				slog.Warn("decode", "path", path, "line", lnum, "error", err)
				return nil
			}

			logs = append(logs, log)
		}

		return nil
	}

	if err := filepath.Walk(root, walkFn); err != nil {
		return nil, err
	}

	return logs, nil
}
