IMAGE = postgres:18.6
NAME = taxi-db

.PHONY: db db-stop

db:
	docker run -d --name $(NAME) -p 5432:5432 -e POSTGRES_PASSWORD=postgres \
		-v $(CURDIR)/schema.sql:/docker-entrypoint-initdb.d/schema.sql:ro $(IMAGE)
	until docker exec $(NAME) pg_isready -h localhost -q; do sleep 1; done
	duckdb < seed.sql

db-stop:
	docker rm -f $(NAME)
