package main

import (
	"encoding/json"
	"fmt"
	"iter"
	"log/slog"
	"logs/loader/mapper"
	"net/http"
	"os"
	"strconv"
)

// Limit returns a sequence of up to "n" items from seq.
func Limit[T any](seq iter.Seq[T], n int) iter.Seq[T] {
	fn := func(yield func(T) bool) {
		i := 0
		for v := range seq {
			if !yield(v) {
				return
			}

			i++
			if i == n {
				return
			}
		}
	}

	return fn
}

func logsHandler(w http.ResponseWriter, r *http.Request) {
	logs, err := mapper.LoadLogs("logs")
	if err != nil {
		slog.Error("load logs", "error", err)
		http.Error(w, "load logs", http.StatusInternalServerError)
		return
	}

	if expr := r.URL.Query().Get("filter"); expr != "" {
		pred, err := createFilter(expr)
		if err != nil {
			slog.Error("filter", "error", err, "expr", expr)
			http.Error(w, "bad filter", http.StatusBadRequest)
			return
		}
		logs = mapper.Filter(logs, pred)
	}

	if s := r.URL.Query().Get("limit"); s != "" {
		limit, err := strconv.Atoi(s)
		if err != nil || limit <= 0 {
			slog.Error("limit", "error", err, "value", limit)
			http.Error(w, "bad limit", http.StatusBadRequest)
			return
		}

		logs = Limit(logs, limit)
	}

	w.Header().Set("content-type", "application/jsonlines")
	ctrl := http.NewResponseController(w)
	enc := json.NewEncoder(w)
	for log := range logs {
		if err := enc.Encode(log); err != nil {
			slog.Error("encode", "error", err, "log", log)
			return
		}
		ctrl.Flush()
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /logs", logsHandler)

	addr := ":8080"
	srv := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	slog.Info("server starting", "address", addr)
	if err := srv.ListenAndServe(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
