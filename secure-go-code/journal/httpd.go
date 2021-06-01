package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var (
	db       *DB
	lastHTML = `<!DOCTYPE html>
<html>
	<head>
		<title>Journal</title>
	</head>
	<body>
		<h2>Last Entry</h2>
		<p>
		%s
		</p>
	</body>
</html>
`
)

func lastHandler(w http.ResponseWriter, r *http.Request) {
	lastText := "No entries"

	entry, err := db.Last()
	if err == nil {
		time := entry.Time.Format("2006-01-02T15:04")
		lastText = fmt.Sprintf("[%s] %s by %s", time, entry.Content, entry.User)
	}
	fmt.Fprintf(w, lastHTML, lastText)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var e Entry

	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	e.Time = time.Now()
	id, err := db.Add(e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	resp := map[string]int{
		"id": id,
	}
	json.NewEncoder(w).Encode(resp)
}

func queryHandler(w http.ResponseWriter, r *http.Request) {
	start := r.URL.Query().Get("start")
	end := r.URL.Query().Get("end")

	startTime, err := time.Parse("2006-01-02", start)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	endTime, err := time.Parse("2006-01-02", end)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	entries, err := db.Query(startTime, endTime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entries)
}

func checkHealth() error {
	return db.Health()
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if err := checkHealth(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprintf(w, "OK\n")
}

func main() {
	var err error
	db, err = NewDB("journal.db")
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/last", lastHandler).Methods("GET")
	r.HandleFunc("/add", addHandler).Methods("POST")
	r.HandleFunc("/query", queryHandler).Methods("GET")
	r.HandleFunc("/health", healthHandler).Methods("GET")

	http.Handle("/", r)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
