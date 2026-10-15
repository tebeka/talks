package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"
	"github.com/urfave/cli/v3"
	"golang.org/x/term"
)

func main() {
	cmd := &cli.Command{
		Name:  "taxi",
		Usage: "query the local taxi-trip database",
		Commands: []*cli.Command{{
			Name:      "query",
			Usage:     "run a read-only SQL query from a file",
			ArgsUsage: "FILE",
			Action: func(ctx context.Context, c *cli.Command) error {
				if c.NArg() != 1 {
					return fmt.Errorf("expected one SQL file")
				}
				input, err := os.ReadFile(c.Args().First())
				if err != nil {
					return err
				}
				statement := strings.TrimSpace(string(input))
				if statement == "" {
					return fmt.Errorf("SQL file is empty")
				}

				if !hasLimit(statement) {
					if err := confirm(); err != nil {
						return err
					}
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
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func isYes(answer string) bool {
	for _, n := range []string{"y", "yes"} {
		if strings.EqualFold(answer, n) {
			return true
		}
	}

	return false
}

func confirm() error {
	fmt.Fprint(os.Stderr, "Query has no WHERE or LIMIT clause. Run it? [y/N] ")
	var answer string

	if _, err := fmt.Fscanln(os.Stdin, &answer); err != nil || (!isYes(answer)) {
		return errors.New("query canceled")
	}

	return nil
}

func display(val any) any {
	if val == nil {
		return "NULL"
	}

	return val
}

func hasLimit(statement string) bool {
	for _, name := range []string{"WHERE", "LIMIT"} {
		if strings.Contains(statement, name) {
			return true
		}
	}

	return false
}

var isTTY = term.IsTerminal(int(os.Stdout.Fd()))

func highlight(s string) string {
	if !isTTY {
		return s
	}

	return "\033[33m" + s + "\033[0m"
}

func formatRow(columns []string, values []any) error {
	for i, col := range columns {
		val := display(values[i])
		fmt.Printf("%s: %v\n", col, val)
	}
	fmt.Println()

	return nil
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
		for i, val := range values {
			if data, ok := val.([]byte); ok {
				values[i] = string(data)
			}
		}

		if err := formatRow(columns, values); err != nil {
			return err
		}

		count++
	}

	if err := errors.Join(rows.Err(), rows.Close()); err != nil {
		return err
	}

	total := fmt.Sprintf("%d row(s)", count)
	fmt.Println(highlight(total))

	return nil
}
