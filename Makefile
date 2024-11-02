build:
	@go build -o bin/api cmd/api/main.go

test:
	@go test -v ./...

run:
	@go run cmd/api/main.go

start: build
	@./bin/api

migration:
	@migrate create -ext sql -dir db/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run db/main.go up

migrate-down:
	@go run db/main.go down