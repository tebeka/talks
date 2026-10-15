package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	_ "github.com/lib/pq"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "taxi",
		Usage: "query the local taxi-trip database",
		Commands: []*cli.Command{{
			Name:  "query",
			Usage: "run a read-only SQL query from stdin",
			Action: func(ctx context.Context, c *cli.Command) error {
				if c.NArg() != 0 {
					return fmt.Errorf("read SQL from stdin, not from an argument")
				}
				input, err := io.ReadAll(c.Reader)
				if err != nil {
					return err
				}
				statement := strings.TrimSpace(string(input))
				if statement == "" {
					return fmt.Errorf("no SQL provided on stdin")
				}

				db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres sslmode=disable")
				if err != nil {
					return err
				}
				defer db.Close()
				return query(ctx, db, statement)
			},
		}},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintln(os.Stderr, "taxi:", err)
		os.Exit(1)
	}
}

func normalize(val any) any {
	if val == nil {
		return "NULL"
	}

	if data, ok := val.([]byte); ok {
		return string(data)
	}

	return val
}

func query(ctx context.Context, db *sql.DB, statement string) error {
	rows, err := db.QueryContext(ctx, statement)
	if err != nil {
		return err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	count := 0
	for rows.Next() {
		values := make([]any, len(columns))
		dest := make([]any, len(columns))
		for i := range values {
			dest[i] = &values[i]
		}
		if err := rows.Scan(dest...); err != nil {
			return err
		}

		for i, col := range columns {
			val := normalize(values[i])
			fmt.Printf("%s: %v\n", col, val)
		}

		if count > 0 {
			fmt.Println()
		}

		count++
	}

	if err := errors.Join(rows.Err(), rows.Close()); err != nil {
		return err
	}

	fmt.Printf("\033[1m%d row(s)\033[0m\n", count)
	return nil
}
