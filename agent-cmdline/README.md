# Taxi CLI

A small Go CLI for querying the local taxi-trip Postgres database.

The CLI connects to `localhost:5432` as `postgres`. No password is required.
Start the local database, then run a query:

```sh
make db-start
printf '%s\n' 'SELECT vendor_id, trip_distance FROM trips LIMIT 5' | go run . query
```

The first version prints a colored table and a row count. Query results go to
stdout. Errors go to stderr. Queries run in a read-only transaction.

The database uses trust authentication for this local demo. The port is bound
to loopback. Do not use this setup on a shared network.
