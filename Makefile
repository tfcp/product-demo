.PHONY: build clean tool lint help

.PHONY: build
build:
	go build main.go

.PHONY: run
run:
	go run main.go server

.PHONY: web
web:
	cd web/ && npm run dev

.PHONY: build-linux
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

.PHONY: build-window
    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
