package main

type DB {}
func (d *DB) Close() error {
	return nil
}

func NewDB() (*DB, error) {
	return nil, nil
}

// START_CREATE OMIT
func createDB(tb *testing.TB) *DB {
	db, err := NewDB()
	if err != nil {
		tb.Fatal(err)
	}
	tb.Cleanup(func() {db.Close()})

	return db
}
// END_CREATE OMIT
