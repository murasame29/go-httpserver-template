BINARY_NAME=app
ENV_FILE=

build:
	go mod tidy
	go build -o bin/$(BINARY_NAME) ./cmd/$(BINARY_NAME)/main.go

run: build
	./bin/$(BINARY_NAME) -env=$(ENV_FILE)

test:
	go test -v ./...

.PHONY: build run test
