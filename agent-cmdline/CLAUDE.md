## What this is

Demo code for the talk "CLI Design for Humans and AI" (`desc.md`). It is a Go CLI (`taxi query FILE`) that runs SQL from a file against a local Postgres database of NYC taxi trips. `playbook.md` has the talk flow, design rules and the demo commands for each step. Read it before adding a step.

This is demo code, simplicity the brevity are the main goals.

## Step-per-directory layout

Each talk step is a standalone `main` package in its own directory. Each step is a full copy of the previous one plus one small change, so the diff between neighbors is what the talk shows:

`initial/` → `tty/` (color only when stdout is a TTY) → `json/` (`-format text|json`) → `yes/` (`-yes` skips the confirmation prompt)

To add a step, `cp -r` the latest directory, make only the new change, and leave earlier steps alone. Then add the demo commands to the list at the top of `playbook.md` (above the `---`). Keep each step small so `diff prev/main.go next/main.go` stays readable on a slide. It's OK to change previous stages in order to arrange the code so the diffs will be small.

All steps share one `go.mod` (module `taxi`, using `urfave/cli/v3`, `lib/pq`, `x/term`). The connection string is hard-coded: `localhost:5432`, user `postgres`, trust auth, no password.

The confirmation prompt is triggered by `hasLimit`, a naive substring check for `WHERE`/`LIMIT`. This is deliberate for the demo. The playbook notes it is not a real safety check.

## Commands

```sh
make db-start          # Postgres 18.6 in docker (taxi-db) + load the parquet file via duckdb
make db-stop
make psql              # psql shell inside the container
go run ./<step> query 3.sql          # e.g. ./yes, ./json
go run ./yes query -yes count.sql    # count.sql has no WHERE/LIMIT, so it triggers the prompt
go vet ./...
```

There are no tests. Check changes by running the step against the database, both on a terminal and piped (`| cat -v`), because output depends on whether stdout is a TTY.

## Data

- `db/schema.sql` runs as the container init script. It creates `trips` plus lookup tables (`vendor`, `rate_code`, `payment_type`, `store_and_fwd`).
- `db/seed.sql` is a DuckDB script that loads `db/yellow_tripdata_2026-05.parquet` (about 4M rows, gitignored) into `trips`. `make db-start` needs `duckdb` on PATH.
