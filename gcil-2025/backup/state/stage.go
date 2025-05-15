package main

import (
	"fmt"
	"iter"
)

type DB struct{}

func (db *DB) Close() error {
	return nil
}

func DBFixture() iter.Seq[*DB] {
	fn := func(yield func(*DB) bool) {
		fmt.Println("setup")
		db := DB{}
		if !yield(&db) {
			return
		}

		// Test code runs

		fmt.Println("teardown")
		db.Close()
		if !yield(nil) {
			return
		}
	}

	return fn
}

func main() {
	fix := DBFixture()
	next, stop := iter.Pull(fix)
	defer stop()

	// Call setup
	db, ok := next()
	if !ok {
		fmt.Println("setup failed")
		return
	}

	_ = db // Work with db

	_, ok = next()
	if !ok {
		fmt.Println("teardown failed")
		return
	}
}
