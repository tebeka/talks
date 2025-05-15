package slice

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"log/slog"
	"os"

	"mordor/log"
)

// LoadLogs return a slice of logs from file.
func LoadLogs(fileName string) ([]log.Log, error) {
	file, err := os.Open(fileName)
	if err != nil {
		slog.Warn("open", "path", fileName, "error", err)
		return nil, err
	}
	defer file.Close()

	gz, err := gzip.NewReader(file)
	if err != nil {
		slog.Warn("gzip", "path", fileName, "error", err)
		return nil, err
	}
	defer gz.Close()

	var logs []log.Log
	dec := json.NewDecoder(gz)
	lnum := 0

	for {
		lnum++
		var record log.Log
		err := dec.Decode(&record)

		if err == io.EOF {
			break
		}

		if err != nil {
			slog.Warn("decode", "path", fileName, "line", lnum, "error", err)
			return nil, err
		}

		if !log.IsValid(record) {
			continue
		}

		logs = append(logs, record)
	}

	return logs, nil
}
