package slice

import (
	"encoding/json"
	"io"
	"log/slog"

	"mordor/log"
)

// LoadLogs return a slice of logs from r.
func LoadLogs(r io.Reader) ([]log.Log, error) {
	var logs []log.Log
	dec := json.NewDecoder(r)
	lnum := 0

	for {
		lnum++
		var record log.Log
		err := dec.Decode(&record)

		if err == io.EOF {
			break
		}

		if err != nil {
			slog.Error("decode", "line", lnum, "error", err)
			return nil, err
		}

		if !log.IsValid(record) {
			continue
		}

		logs = append(logs, record)
	}

	return logs, nil
}
