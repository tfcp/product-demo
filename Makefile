.PHONY: build clean tool lint help

build:
	@go build -v .

run:
    @go run main.go server