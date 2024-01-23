BINARY_NAME=app


build:
	go mod tidy
	go build -o bin/$(BINARY_NAME) ./cmd/$(BINARY_NAME)/main.go

run: build
	./bin/$(BINARY_NAME)

test:
	go test -v ./...

.PHONY: build run test
