.PHONY: test build run

test:
	go test ./...

build:
	go build -o ./cmd/server/main.out ./cmd/server/main.go

run: build
	./cmd/server/main.out -w ./public/html
