# Taxi CLI

A small Go CLI for querying the local taxi-trip Postgres database.

The CLI connects to `localhost:5432` as `postgres`. No password is required.
Start the local database, then run a query:

```sh
make db-start
cat > /tmp/taxi.sql <<'SQL'
SELECT vendor_id, trip_distance FROM trips WHERE vendor_id = 1 LIMIT 5
SQL
go run ./tty query /tmp/taxi.sql
```

The initial version is in `initial/`. The `tty/` version colors the row count
only when stdout is a terminal. It prints plain text when output is piped.
Both versions print each row as `column: value` lines and send errors to stderr.

The database uses trust authentication for this local demo. The port is bound
to loopback. Do not use this setup on a shared network.
