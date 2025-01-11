include .env

air:
	air

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

migrate-up:
	migrate -path database/migrations/ -database "postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5432/$(POSTGRES_DB_NAME)?sslmode=disable" -verbose up

migrate-down:
	migrate -path database/migrations/ -database "postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5432/$(POSTGRES_DB_NAME)?sslmode=disable" down 1

#TODO: findout how to run migrations via docker

#migrate-create:
#	docker compose run --user migrate create -ext sql -dir /migrations -seq $(name)
#
#migrate-up:
#	docker compose run migrate-up
#
#migrate-down:
#	docker compose run migrate-down

db-up:
	docker compose up db

swagger-docs:
	 swag init
