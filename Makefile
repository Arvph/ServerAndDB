.PHONY: all

all:
	go build -v ./cmd/api


up:
	migrate -path migrations -database "postgres://localhost:5434/restapi?sslmode=disable&user=postgres&password=admin" up

down:
	migrate -path migrations -database "postgres://localhost:5434/restapi?sslmode=disable&user=postgres&password=admin" down