.PHONY:install clean tool lint

.PHONY: install
install:
	@cd . && go build -o server main.go

.PHONY: server
all: install
server:
	./server server

.PHONY: cron
all: install
cron:
	./server cron

.PHONY: process
all: install
process:
	./server process

.PHONY: web
web:
	cd web/ && npm run dev

.PHONY: clean
clean:
	@rm -rf ./server ./sre-* ./web/web.zip

.PHONY: rice
rice:
	go get github.com/GeertJohan/go.rice
	go get github.com/GeertJohan/go.rice/rice
	cd tools/rice
	rice embed-go

.PHONY: build-static
build-static:
	cd web/ && npm run build:prod

.PHONY: build-linux
build-linux:
	@cd . && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o server-linux main.go

.PHONY: build-window
build-window:
	@cd . && CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o server-windows main.go
