root := ./app/grpc
dirs := v1
name := $(shell basename $(shell pwd))
src  := $(shell cd $(root) && find $(dirs) -name "*.proto")
out  := ./
before := go build -o server main.go

.PHONY: go php js clean

.PHONY: server
server:
	 @export ENV=DEV && ${before} && ./server server

go:
	@$(foreach file, $(src), protoc $(file) -I $(root) --go_out=plugins=grpc:$(out);)

clean:
	@rm -rf ./app/grpc/*/*.go ./server
