package main

import (
	"context"
	"database/sql"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/lib/pq"
	"github.com/urfave/cli/v3"
	"golang.org/x/term"
)

const dsn = "host=localhost port=5432 user=postgres dbname=postgres sslmode=disable"

const maxRows = 1000

func main() {
	cmd := &cli.Command{
		Name:  "taxi",
		Usage: "query the local taxi-trip database",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "format",
				Usage: "output format: text or json",
				Value: "text",
			},
		},
		Before: func(ctx context.Context, c *cli.Command) (context.Context, error) {
			switch c.String("format") {
			case "json":
				formatRow = formatRowJSON
				formatTotal = formatTotalJSON
				formatError = formatErrorJSON
			case "text":
			default:
				return ctx, fmt.Errorf("format must be text or json")
			}

			return ctx, nil
		},
		Commands: []*cli.Command{{
			Name:      "query",
			Usage:     "run a read-only SQL query from a file",
			ArgsUsage: "FILE",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "yes",
					Usage: "run without asking for confirmation",
				},
				&cli.IntFlag{
					Name:  "limit",
					Usage: "maximum number of rows to print, 0 for all",
					Value: maxRows,
				},
			},
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

				if !c.Bool("yes") && !hasLimit(statement) {
					if err := confirm(); err != nil {
						return err
					}
				}

				db, err := sql.Open("postgres", dsn)
				if err != nil {
					return err
				}
				defer db.Close()

				if !c.IsSet("limit") {
					n, err := estimateRows(ctx, db, statement)
					if err != nil {
						return err
					}
					if n > maxRows {
						return fmt.Errorf("query returns about %d rows, use -limit N to get the first N rows (-limit 0 for all) or update your query", n)
					}
				}

				return query(ctx, db, int(c.Int("limit")), statement)
			},
		}, {
			Name:      "schema",
			Usage:     "list tables, or show the columns of TABLE",
			ArgsUsage: "[TABLE]",
			Action: func(ctx context.Context, c *cli.Command) error {
				db, err := sql.Open("postgres", dsn)
				if err != nil {
					return err
				}
				defer db.Close()

				if c.NArg() == 0 {
					return query(ctx, db, 0, tablesSQL)
				}

				return query(ctx, db, 0, columnsSQL, c.Args().First())
			},
		}},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		formatError(err)
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
	if !term.IsTerminal(int(os.Stdin.Fd())) {
		return errors.New("query has no WHERE or LIMIT clause, use -yes to run it")
	}

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

var formatRow = formatRowText
var formatTotal = formatTotalText
var formatError = formatErrorText

func highlight(s string) string {
	if !isTTY {
		return s
	}

	return "\033[33m" + s + "\033[0m"
}

func formatRowJSON(columns []string, values []any) error {
	row := make(map[string]any, len(columns))
	for i, col := range columns {
		row[col] = values[i]
	}

	return json.NewEncoder(os.Stdout).Encode(row)
}

func formatRowText(columns []string, values []any) error {
	for i, col := range columns {
		val := display(values[i])
		fmt.Printf("%s: %v\n", col, val)
	}
	fmt.Println()

	return nil
}

func formatTotalJSON(count int, truncated bool) error {
	return json.NewEncoder(os.Stdout).Encode(map[string]any{"_meta": map[string]any{"count": count, "truncated": truncated}})
}

func formatTotalText(count int, truncated bool) error {
	total := fmt.Sprintf("%d row(s)", count)
	if truncated {
		total += ", truncated (use -limit 0 for all)"
	}
	fmt.Println(highlight(total))

	return nil
}

func errorHint(err error) string {
	pqErr, ok := errors.AsType[*pq.Error](err)
	if !ok {
		return ""
	}

	if pqErr.Hint != "" {
		return pqErr.Hint
	}

	switch pqErr.Code {
	case "42P01", "42703": // undefined_table, undefined_column
		return `run "taxi schema" to list tables, "taxi schema TABLE" to list columns`
	}

	return ""
}

func formatErrorJSON(err error) {
	out := map[string]string{"error": err.Error()}
	if hint := errorHint(err); hint != "" {
		out["hint"] = hint
	}

	json.NewEncoder(os.Stderr).Encode(out)
}

func formatErrorText(err error) {
	fmt.Fprintln(os.Stderr, "error:", err)
	if hint := errorHint(err); hint != "" {
		fmt.Fprintln(os.Stderr, "hint:", hint)
	}
}

var (
	//go:embed tables.sql
	tablesSQL string

	//go:embed columns.sql
	columnsSQL string
)

func estimateRows(ctx context.Context, db *sql.DB, statement string) (int, error) {
	var data []byte
	if err := db.QueryRowContext(ctx, "EXPLAIN (FORMAT JSON) "+statement).Scan(&data); err != nil {
		return 0, err
	}

	var plans []struct {
		Plan struct {
			Rows int `json:"Plan Rows"`
		}
	}
	if err := json.Unmarshal(data, &plans); err != nil {
		return 0, err
	}

	return plans[0].Plan.Rows, nil
}

func query(ctx context.Context, db *sql.DB, limit int, statement string, args ...any) error {
	rows, err := db.QueryContext(ctx, statement, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	count, truncated := 0, false
	for rows.Next() {
		if limit > 0 && count == limit {
			truncated = true
			break
		}

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

	return formatTotal(count, truncated)
}
