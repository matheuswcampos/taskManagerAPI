.PHONY: install run test

install:
	go mod download

run:
	go run ./app

test:
	go test ./...
