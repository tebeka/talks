package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
)

func main() {
	if err := run(); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return
		}

		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	host := flag.String("h", envOr("PGHOST", "localhost"), "database server host")
	port := flag.String("p", envOr("PGPORT", "5432"), "database server port")
	user := flag.String("U", envOr("PGUSER", "postgres"), "database user")
	database := flag.String("d", envOr("PGDATABASE", "postgres"), "database name")
	passwd := flag.String("W", envOr("PGPASSWORD", ""), "password")
	flag.Usage = func() {
		name := filepath.Base(os.Args[0])
		fmt.Fprintf(os.Stderr, "Usage: %s [-h host] [-p port] [-U user] [-d database] [-W password] [sql-file] (or SQL via stdin)\n", name)
		flag.PrintDefaults()
	}

	flag.Parse()
	if flag.NArg() > 1 {
		return errors.New("expected at most one SQL file")
	}

	var query []byte
	if flag.NArg() == 1 {
		file := flag.Arg(0)
		info, err := os.Stat(file)
		if err != nil {
			return fmt.Errorf("stat SQL file: %w", err)
		}
		if !info.Mode().IsRegular() {
			return fmt.Errorf("SQL file %q is not a regular file", file)
		}
		query, err = os.ReadFile(filepath.Clean(file))
		if err != nil {
			return fmt.Errorf("read SQL file: %w", err)
		}
	} else {
		if stdinIsTerminal() {
			return errors.New("provide SQL in a file or pipe it on stdin")
		}
		var err error
		query, err = io.ReadAll(os.Stdin)
		if err != nil {
			return fmt.Errorf("read SQL from stdin: %w", err)
		}
	}
	if len(bytes.TrimSpace(query)) == 0 {
		return errors.New("SQL input is empty")
	}

	portNumber, err := strconv.ParseUint(*port, 10, 16)
	if err != nil || portNumber == 0 {
		return fmt.Errorf("invalid database port %q", *port)
	}

	config, err := pgx.ParseConfig("")
	if err != nil {
		return fmt.Errorf("configure PostgreSQL connection: %w", err)
	}
	config.Host = *host
	config.Port = uint16(portNumber)
	config.User = *user
	config.Password = *passwd
	config.Database = *database

	db := stdlib.OpenDB(*config)
	defer db.Close()
	ctx := context.Background()
	if err := db.PingContext(ctx); err != nil {
		return fmt.Errorf("connect to PostgreSQL: %w", err)
	}

	rows, err := db.QueryContext(ctx, string(query))
	if err != nil {
		return fmt.Errorf("execute SQL: %w", err)
	}
	defer rows.Close()
	if err := writeJSONRows(os.Stdout, rows); err != nil {
		return fmt.Errorf("write JSON results: %w", err)
	}

	return nil
}

func writeJSONRows(out io.Writer, rows *sql.Rows) error {
	columns, err := rows.Columns()
	if err != nil {
		return err
	}
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		return err
	}

	if _, err := io.WriteString(out, "["); err != nil {
		return err
	}
	first := true
	for rows.Next() {
		values := make([]any, len(columns))
		scanTargets := make([]any, len(columns))
		for i := range values {
			scanTargets[i] = &values[i]
		}
		if err := rows.Scan(scanTargets...); err != nil {
			return err
		}

		row := make(map[string]any, len(columns))
		for i, column := range columns {
			row[column] = jsonValue(values[i], columnTypes[i].DatabaseTypeName())
		}
		encoded, err := json.Marshal(row)
		if err != nil {
			return err
		}
		if !first {
			if _, err := io.WriteString(out, ","); err != nil {
				return err
			}
		}
		if _, err := out.Write(encoded); err != nil {
			return err
		}
		first = false
	}
	if err := rows.Err(); err != nil {
		return err
	}
	_, err = io.WriteString(out, "]\n")
	return err
}

func jsonValue(value any, databaseType string) any {
	typeName := strings.ToUpper(databaseType)
	switch value := value.(type) {
	case []byte:
		switch typeName {
		case "JSON", "JSONB":
			return json.RawMessage(value)
		case "NUMERIC", "DECIMAL":
			return json.Number(value)
		case "BYTEA":
			return value
		default:
			return string(value)
		}
	case string:
		switch typeName {
		case "JSON", "JSONB":
			return json.RawMessage(value)
		case "NUMERIC", "DECIMAL":
			return json.Number(value)
		}
	}
	return value
}

func envOr(name, fallback string) string {
	if value := os.Getenv(name); value != "" {
		return value
	}

	return fallback
}

func stdinIsTerminal() bool {
	info, err := os.Stdin.Stat()
	return err == nil && info.Mode()&os.ModeCharDevice != 0
}
