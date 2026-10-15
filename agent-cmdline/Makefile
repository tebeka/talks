IMAGE = postgres:18.6
NAME = taxi-db

.PHONY: db-start db-stop

db-start:
	docker run -d --name $(NAME) -p 127.0.0.1:5432:5432 -e POSTGRES_HOST_AUTH_METHOD=trust \
		-v $(CURDIR)/db/schema.sql:/docker-entrypoint-initdb.d/schema.sql:ro $(IMAGE)
	until docker exec $(NAME) pg_isready -h localhost -q; do sleep 1; done
	duckdb < db/seed.sql

db-stop:
	docker rm -f $(NAME)


psql:
	docker exec -it $(NAME) psql -U postgres
