ifeq ($(GOHOSTOS), windows)
	Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find proto/api -name *.proto")
else
	API_PROTO_FILES=$(shell find proto/api -name *.proto)
endif

path := $(shell pwd)

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.3
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.25.1
	go install github.com/protoc-gen/protoc-gen-validatex@v0.7.0
	go install github.com/protoc-gen/protoc-gen-go-errors@v0.3.0
	go install github.com/google/wire/cmd/wire@v0.6.0
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@v0.7.0
	go install mvdan.cc/gofumpt@latest

.PHONY: config
# generate internal proto
config:
	protoc --proto_path=./proto/conf \
		   --go_out=paths=source_relative:./internal/conf \
		   ./proto/conf/*.proto

.PHONY: common
# generate internal proto
common:
	protoc --proto_path=./proto/common \
 	       --go_out=paths=source_relative:./internal/common \
 	       --go-grpc_out=paths=source_relative:./internal/common \
		   ./proto/common/**/*.proto

.PHONY: errors
# generate errors
errors:
	mkdir -p ./internal/common/xerr
	protoc --proto_path=./proto/xerr \
           --proto_path=./proto/third_party \
           --go_out=paths=source_relative:./internal/common/xerr \
           --go-errors_out=paths=source_relative:./internal/common/xerr \
           ./proto/xerr/*.proto

.PHONY: api
# generate api proto
api:
	mkdir -p ./internal/api
	protoc --proto_path=./proto/api \
 	       --proto_path=./proto/third_party \
 	       --proto_path=./proto/common \
 	       --go_out=paths=source_relative:./internal/api \
 	       --go-grpc_out=paths=source_relative:./internal/api \
 	       --grpc-gateway_out=paths=source_relative:./internal/api \
		   --validatex_out=paths=source_relative:./internal/api \
		   --validatex_opt=i18n_dir=./i18n/validatex,i18n_out_relative_dir=../../i18n/validatex \
	       --openapi_out=fq_schema_naming=true,default_response=false:./swagger_ui \
	       $(API_PROTO_FILES)

.PHONY: stringer
stringer:
	go generate $(path)/pkg/vobj

.PHONY: generate
# generate
generate:
	go generate ./...

.PHONY: wire
wire:
	cd ./cmd/server && wire

.PHONY: all
# generate all
all:
	make stringer;
	make generate;
	make config;
	make common;
	make errors;
	make api;
	make wire;
	go mod tidy;

.PHONY: build
# build
build:
	go build -o ./bin/server ./cmd/server

.PHONY: test
# run all tests
test:
	go test -v ./...

.PHONY: format
format:
	gofmt -s -w .
	golint ./...
	go vet ./...
	go mod tidy
	go mod verify
	goimports -w .
	golangci-lint run ./...
	gofumpt -l -w .