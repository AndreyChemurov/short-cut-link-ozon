BUILD_FILE_PATH = cmd/ozon/main.go

LINK_SHORTER_FILENAME = link_shorter
LINK_SHORTER_PROTO_PATH = internal/api/grpc/proto/$(LINK_SHORTER_FILENAME).proto

build: $(BUILD_FILE_PATH)
	@go build -o ./bin/main $(BUILD_FILE_PATH)

run: ./bin/main
	@./bin/main

lint:
	@golangci-lint run -v

grpc-gen: $(LINK_SHORTER_PROTO_PATH)
	@protoc -I internal/api/grpc/proto \
		--go_out=internal/api/grpc/gen \
		--go-grpc_out=internal/api/grpc/gen \
		--plugin=$(GOBIN)/protoc-gen-go \
		$(LINK_SHORTER_PROTO_PATH)
