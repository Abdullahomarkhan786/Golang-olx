.PHONY:build run

build:
	@go build -o  bin/api ./cmd/api
run:build
	@./bin/api

migrate-up:
	@go run ./cmd/migrate up

migrate-down:
	@go run ./cmd/migrate down

# usage: make migrate-force VERSION=<n>  (use -1 to reset to no version)
migrate-force:
	@go run ./cmd/migrate force $(VERSION)
