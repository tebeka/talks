# Taxi CLI

A small Go CLI for querying the local taxi-trip Postgres database.

The CLI connects to `localhost:5432` as `postgres`. No password is required.
Start the local database, then run a query:

```sh
make db-start
cat > /tmp/taxi.sql <<'SQL'
SELECT vendor_id, trip_distance FROM trips WHERE vendor_id = 1 LIMIT 5
SQL
go run . query /tmp/taxi.sql
```

The CLI prints each row as `column: value` lines. It prints the row count in
bold. Errors go to stderr. If the SQL text has no `WHERE`, the CLI asks on the
terminal. Type `y` or `yes` to continue. Without a terminal, it stops with an
error.

The database uses trust authentication for this local demo. The port is bound
to loopback. Do not use this setup on a shared network.
