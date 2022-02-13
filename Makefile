.PHONY: build clean tool lint help

.PHONY: build
build:
	go build main.go

.PHONY: run
run:
	go run main.go server

.PHONY: cron
cron:
	go run main.go cron

.PHONY: process
process:
	go run main.go process

.PHONY: web
web:
	cd web/ && npm run dev

.PHONY: build-static
build-static:
	cd web/ && npm run build:prod

.PHONY: build-linux
build-linux:
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

.PHONY: build-window
build-window:
    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go

