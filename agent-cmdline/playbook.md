- Show command line usage
    go run ./initial query 3.sql
    go run ./initial query 3.sql | cat -v
- 
    go run ./tty query 3.sql
    go run ./tty query 3.sql | cat -v
---
# Talk playbook: CLI Design for Humans and AI

## Goal

Build a Go command-line tool that queries the local Postgres taxi-trip database. Improve it in small steps so it works well for a person at a terminal and for an AI agent running commands without a terminal.

The project already includes the Postgres schema and seed setup. DuckDB loads the taxi Parquet file into Postgres.

## Main points

### 1. Non-interactive operation

- Do not require input prompts to run normal read-only queries.
- Do not wait for a person when stdin is not a terminal.
- Use command arguments and options for required inputs. Give a clear error when input is missing.
- Do not ask for confirmation for `SELECT` queries.
- Ask a person to confirm state-changing operations such as `UPDATE`, `DELETE`, `INSERT`, or `DROP` in interactive use.
- In non-interactive use, do not wait for input or silently run a state-changing operation. Return a clear error that says approval is required.
- Enforce write safety with database permissions. A confirmation prompt is not a security boundary.
- Do not detect writes by checking only the first SQL word. SQL can contain multiple statements or a write inside a `WITH` clause.

### 2. Human and agent output

- Detect whether stdout is a terminal. Use a readable human format for terminal use and a compact format for piped or agent use.
- Provide an explicit format option. An explicit user choice must take precedence over automatic detection.
- Consider an explicit option to force or disable agent defaults.
- Keep query results on stdout. Send progress, warnings, and diagnostics to stderr.
- Use compact tables for easy reading. Offer JSON or NDJSON when callers need structured data.
- Keep repeated calls quiet. Give a short mode notice only when it helps. Put full behavior details in help.

### 3. Result size and query cost

- Set useful row, byte, and cell-size limits.
- Never make a partial result look complete. Say whether the result is complete, limited, sampled, or stopped early. Report counts when known.
- Avoid returning every row from a large table by default.
- For a later enhancement, show a query estimate or progress for long-running queries.
- A result hash could help compare large results, but it is optional and not needed for the first demo.

### 4. Errors that support recovery

- Return a non-zero exit status for a failed command.
- Keep errors on stderr, separate from result data.
- Use a stable machine-readable error form in agent mode. Include a short message, a stable error code, and a useful correction hint when possible.
- For Postgres errors, include SQLSTATE and position when available.
- Keep the error compact. Do not repeat the same information in several fields.

### 5. Schema discovery

- Give the caller a way to list tables and inspect columns and types.
- This helps an agent learn the database before it writes a query.
- Approximate row counts or query plans can be later enhancements.

## DuckDB Agent Mode lessons

The DuckDB post describes agent mode for its CLI. Useful ideas for this talk are:

- Agent detection can use known environment variables together with stdout not being a terminal.
- Automatic behavior should stay out of the way when the user selects an explicit output format.
- Compact Markdown tables reduce formatting noise. Column types can be shown with names.
- Truncation must be explicit. DuckDB reports omitted rows and can stop large results early.
- Errors can use a compact JSON form on stderr.
- Schema inspection, query estimates, and progress can reduce guesswork.
- Give one short startup notice. Keep full instructions in help instead of repeating them on every call.

The post reports 59% fewer output tokens in its experiment, with all answers correct. It also reports no reduction in total cost or runtime. Use this as evidence that compact output can reduce output tokens, not as a claim that it always makes agent work cheaper or faster.

## Initial 45-minute flow

### 0–4 min: Set up the problem

- Introduce the taxi-trip database and the query task.
- State the goal: one CLI should work for people and agents.
- Explain that agents often run a command without a terminal and read only captured output.

### 4–10 min: Build the basic query flow

- Run a useful query against the trips data.
- Show the starting human-oriented output.
- Include a way to inspect tables and columns so the query is not based on guesswork.

### 10–16 min: Define the non-interactive and write-safety rules

- Run a read-only query with no confirmation prompt.
- Show confirmation only for a state-changing operation in interactive use.
- Show that a non-interactive state-changing operation stops with a clear approval-required error.
- Explain that database permissions enforce write safety.

### 16–23 min: Adapt output to the caller

- Detect terminal versus piped output.
- Show a human-readable terminal view and a compact agent view.
- Show that an explicit output-format option takes precedence.

### 23–28 min: Separate output streams

- Keep query data on stdout.
- Send progress and diagnostics to stderr.
- Pipe or capture the command output to show that status text does not corrupt the result.

### 28–34 min: Bound results

- Add row or output limits.
- Show a large result and make its truncation explicit.
- Explain how the caller can request more data safely.

### 34–40 min: Make errors recoverable

- Cause a query error, such as using an unknown column.
- Show the exit status and structured error on stderr.
- Use the error hint to correct and rerun the query.

### 40–45 min: Run an end-to-end agent task

- Ask an agent a question about taxi trips.
- Have it inspect the schema, run a query, and explain the result.
- Include one recoverable query error or one limited result.
- Recap the design rules: no blocking prompts, clear output, visible limits, useful errors, and safe write controls.

## First implementation priorities

1. Query the local database and inspect its schema.
2. Make read-only use non-interactive.
3. Separate stdout and stderr.
4. Add compact output and an explicit format option.
5. Add result limits with clear truncation details.
6. Add structured, actionable errors.
7. Add safe behavior for state-changing SQL if write support is part of the demo.

Keep query-cost estimates, progress reporting, result hashes, and advanced agent detection as optional follow-up work. Do not let them displace the core demo.

---
https://duckdb.org/2026/10/09/agent-mode
- y when there's no where clause
- security: read only user
