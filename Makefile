ifeq ($(GOHOSTOS), windows)
	Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find proto/api -name *.proto")
else
	API_PROTO_FILES=$(shell find proto/api -name *.proto)
endif

.PHONY: init
# init env
init:
	go mod tidy
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.3
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.25.1
	go install github.com/protoc-gen/protoc-gen-validatex@v0.3.2
	go install github.com/memoria-x/protoc-gen-go-errors@v0.2.0
	go install github.com/google/wire/cmd/wire@v0.6.0

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
 	       --go_out=paths=source_relative:./internal/common/cnst \
 	       --go-grpc_out=paths=source_relative:./internal/common/cnst \
		   ./proto/common/*.proto

.PHONY: errors
# generate errors
errors:
	protoc --proto_path=./proto/xerr \
           --proto_path=./proto/third_party \
           --go_out=paths=source_relative:./internal/common/xerr \
           --go-errors_out=paths=source_relative:./internal/common/xerr \
           ./proto/xerr/*.proto

.PHONY: api
# generate api proto
api:
	protoc --proto_path=./proto/api \
 	       --proto_path=./proto/third_party \
 	       --go_out=paths=source_relative:./internal/api \
 	       --go-grpc_out=paths=source_relative:./internal/api \
 	       --grpc-gateway_out=paths=source_relative:./internal/api \
		   --validatex_out=paths=source_relative:./internal/api \
		   --validatex_opt=i18n_dir=./i18n/validatex,i18n_out_relative_dir=../../i18n/validatex \
	       $(API_PROTO_FILES)

.PHONY: generate
# generate
generate:
	go generate ./...

.PHONY: all
# generate all
all:
	make generate;
	make config;
	make common;
	make errors;
	make api;
	go mod tidy;

.PHONY: test
# run all tests
test:
	go test -v ./...
