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

migrate-create:
	docker-compose run migrate create -ext sql -dir /migrations -seq $(name)

migrate-up:
	docker compose run migrate-up

migrate-down:
	docker-compose run migrate-down

swagger-docs:
	 swag init