PROTO_DIR=api/v1/proto
GEN_DIR=src/pkg/gen

PROTO_FILES=$(wildcard $(PROTO_DIR)/*.proto)

.PHONY: proto clean deps

deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

proto:
	@echo "Generating protobuf..."
	protoc \
    	-I=$(PROTO_DIR) \
    	--go_out=$(GEN_DIR) --go_opt=paths=source_relative \
    	--go-grpc_out=$(GEN_DIR) --go-grpc_opt=paths=source_relative \
    	$(PROTO_DIR)/services.proto
	@echo "Done."

clean:
	@echo "Cleaning generated files..."
	rm -rf $(GEN_DIR)/*