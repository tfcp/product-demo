root := ./app/grpc
dirs := v1
name := $(shell basename $(shell pwd))
src  := $(shell cd $(root) && find $(dirs) -name "*.proto")
out  := ./

.PHONY: go php js clean

server:
	@export ENV=DEV && go run main.go server

go:
	@$(foreach file, $(src), protoc $(file) -I $(root) --go_out=plugins=grpc:$(out);)

clean:
	@rm -rf ./app/grpc/*/*.go
