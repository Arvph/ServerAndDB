.PHONY: all postgres createdb dropdb migrateup migratedown

all:
	go build -v ./cmd/api

postgres:
	docker run --name postgres16 -p 5434:5434 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin -d postgres:16-alpine

createdb:
	docker exec -it postgres16 createdb --username=admin --owner=admin test_db

dropdb:
	docker exec -it postgres16 dropdb --force test_db

migrateup:
	migrate -path migration -database "postgres://localhost:5434/restapi?sslmode=disable&user=admin&password=admin" -verbose up

migratedown:
	migrate -path migration -database "postgres://localhost:5434/restapi?sslmode=disable&user=admin&password=admin" -verbose down
