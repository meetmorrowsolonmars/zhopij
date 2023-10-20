
PARENT_DIR:=$(shell dirname $(CURDIR))
LOCAL_BIN:=$(PARENT_DIR)/bin
VENDOR_PROTO:=$(PARENT_DIR)/vendor.proto

.PHONY: generate
generate:
	$(info Generate GRPC Post stubs...)

	PATH="$(LOCAL_BIN):$(PATH)" protoc -I $(VENDOR_PROTO) -I . \
			--go_out ./internal/pb --go_opt paths=source_relative \
    		--go-grpc_out ./internal/pb --go-grpc_opt paths=source_relative \
			api/v1/post/post.proto && \
		go mod tidy

.PHONY: migrate-up
migrate-up:
	$(info Migrate Post...)

	goose --dir ./tools/migrations postgres "user=$(POSTGRES_POST_USER) password=$(POSTGRES_POST_USER_PASSWORD) database=$(POSTGRES_POST_DB) host=$(POSTGRES_EXTERNAL_HOST) port=$(POSTGRES_EXTERNAL_PORT) sslmode=disable" up
