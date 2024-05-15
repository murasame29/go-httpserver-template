BINARY_NAME=app
ENV_FILE=

build:
	docker compose build

run: build
	docker compose up

rund: build
	docker compose up -d

test:
	go test -v ./...
