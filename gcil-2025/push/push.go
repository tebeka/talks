package push

import (
	"database/sql"
)

// Transaction inserts records to database in a transaction.
type Transaction struct {
	sql string
	db  *sql.DB
	tx  *sql.Tx
}

// NewTransaction creates a new Transaction instance.
func NewTransaction(db *sql.DB, sql string) (*Transaction, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	t := Transaction{
		sql: sql,
		db:  db,
		tx:  tx,
	}
	return &t, nil
}

// Add adds a record to the transaction.
func (t *Transaction) Add(record []any) error {
	_, err := t.tx.Exec(t.sql, record...)
	if err != nil {
		t.tx.Rollback()
	}

	return err
}

// Close commits the transaction.
func (t *Transaction) Close() error {
	if t.tx == nil {
		return nil
	}

	return t.tx.Commit()
}
