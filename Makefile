.PHONY: build test clean run lint

BINARY_NAME=ghfetcher
VERSION=1.0.0

build:
	mkdir -p bin
	go build -o bin/$(BINARY_NAME) cmd/ghfetcher/main.go

run:
	go run cmd/ghfetcher/main.go --user torvalds

run-with-repos:
	go run cmd/ghfetcher/main.go --user torvalds --repos --format json

test:
	go test ./...

clean:
	rm -rf bin/

lint:
	gofmt -d .
	golangci-lint run

install: build
	sudo cp bin/$(BINARY_NAME) /usr/local/bin/

version:
	@echo $(VERSION)

help:
	@echo "Available commands:"
	@echo "  build        - Build the binary"
	@echo "  run          - Run with default user"
	@echo "  run-with-repos - Run with repositories"
	@echo "  test         - Run tests"
	@echo "  lint         - Check code quality"
	@echo "  install      - Install to /usr/local/bin"