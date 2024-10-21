# Load environment variables from .env file
ifneq (,$(wildcard .env))
    include .env
    export
endif

postgres:
	docker run --name postgres17 -p ${DB_PORT}:5432 -e POSTGRES_PASSWORD=${DB_PASSWORD} -d postgres:17-alpine

createdb:
	docker exec -it postgres17 createdb --username=${DB_USER} --owner=${DB_USER} ${DB_NAME}

dropdb:
	docker exec -it postgres17 dropdb --username=${DB_USER} ${DB_NAME}

migrateup:
	migrate -path backend/internal/db/migrations -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_ADDRESS}/${DB_NAME}?sslmode=disable" -verbose up

migratedown:
	migrate -path backend/internal/db/migrations -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_ADDRESS}/${DB_NAME}?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test