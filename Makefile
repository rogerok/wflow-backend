include .env

build:
	@go build -o bin/api

run: build
	@./bin/api

test:
	@go test -v ./...

dev: build
	@./bin/api &
	docker compose up

dbseq:
	migrate create -ext sql -dir database/migrations -seq $(name)


migrate:
	migrate -path database/migrations/ -database "postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5432/$(POSTGRES_DB_NAME)?sslmode=disable" -verbose up
