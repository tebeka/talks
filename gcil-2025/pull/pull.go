package pull

import (
	"database/sql"
	"iter"
)

// Transaction inserts records into the database within a transaction.
func Transaction(db *sql.DB, sql string, records iter.Seq[[]any]) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for record := range records {
		_, err := tx.Exec(sql, record...)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
