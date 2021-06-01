package main

import (
	"database/sql"
	_ "embed"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var (
	//go:embed sql/schema.sql
	schemaSQL string

	//go:embed sql/add.sql
	addSQL string

	//go:embed sql/query.sql
	querySQL string

	//go:embed sql/last.sql
	lastSQL string
)

type DB struct {
	db *sql.DB
}

func NewDB(dsn string) (*DB, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(schemaSQL); err != nil {
		db.Close()
		return nil, err
	}

	return &DB{db}, nil
}

func strDate(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func (d *DB) Add(entry Entry) (int, error) {
	sql := fmt.Sprintf(addSQL, strDate(entry.Time), entry.User, entry.Content)
	res, err := d.db.Exec(sql)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	return int(id), err
}

func (d *DB) Query(start, end time.Time) ([]Entry, error) {
	sql := fmt.Sprintf(querySQL, strDate(start), strDate(end))
	rows, err := d.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []Entry
	for rows.Next() {
		var e Entry
		if err := rows.Scan(&e.Time, &e.User, &e.Content); err != nil {
			return nil, err
		}
		entries = append(entries, e)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return entries, nil
}

func (d *DB) Last() (Entry, error) {
	row := db.db.QueryRow(lastSQL)
	var e Entry
	if err := row.Scan(&e.Time, &e.User, &e.Content); err != nil {
		return Entry{}, err
	}

	return e, nil
}

func (d *DB) Health() error {
	return d.db.Ping()
}

func (d *DB) Close() error {
	return d.db.Close()
}
