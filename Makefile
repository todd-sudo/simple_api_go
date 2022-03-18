.PHONY: build

run:build start

start:
	./.bin/app

build:
	go build -o .bin/app -v cmd/web/main.go

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := run