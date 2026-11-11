IMAGE = postgres:18-alpine
NAME = taxi-db
DOCKER ?= docker

.PHONY: db-start db-stop psql

default:
	$(error pick a target)

db/yellow_tripdata_2026-05.parquet:
	curl -o $@ https://d37ci6vzurychx.cloudfront.net/trip-data/yellow_tripdata_2026-05.parquet

db-start: db/yellow_tripdata_2026-05.parquet db-stop
	$(DOCKER) run -d --name $(NAME) -p 127.0.0.1:5432:5432 -e POSTGRES_HOST_AUTH_METHOD=trust \
		-v $(CURDIR)/db/schema.sql:/docker-entrypoint-initdb.d/schema.sql:ro $(IMAGE)
	until $(DOCKER) exec $(NAME) pg_isready -h localhost -q; do sleep 1; done
	duckdb < db/seed.sql

db-stop:
	$(DOCKER) rm -f $(NAME)

psql:
	$(DOCKER) exec -it $(NAME) psql -U postgres

lint-sql:
	sqlfluff lint
