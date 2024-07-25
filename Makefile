.PHONEY:
.SILENT:

LOCAL_BIN:=$(CURDIR)/bin

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.59.1

lint:
	./bin/golangci-lint run ./... --config .golangci.pipeline.yaml

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.21.1

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u github.com/joho/godotenv
	go get -u github.com/jackc/pgx/v4
	go get -u github.com/Masterminds/squirrel

generate: get-deps install-deps
	make generate-chat-api

generate-chat-api:
	mkdir -p ./pkg/chat/v1
	protoc --proto_path=./api/chat/v1 --go_out=./pkg/chat/v1 \
	--go_opt=paths=source_relative --plugin=protoc-gen-go=./bin/protoc-gen-go \
	--go-grpc_out=./pkg/chat/v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=./bin/protoc-gen-go-grpc \
	./api/chat/v1/chat.proto

local-docker-compose-up:
	docker compose --env-file ./config/.env stop chat-migrator 
	docker compose --env-file ./config/.env stop chat-server
	docker compose --env-file ./config/.env rm -f chat-migrator
	docker compose --env-file ./config/.env rm -f chat-server
	docker compose --env-file ./config/.env build chat-migrator 
	docker compose --env-file ./config/.env build chat-server
	docker compose --env-file ./config/.env up --force-recreate -d

build:
	go build -o ./bin/auth ./cmd/chat/main.go

run: build
	./bin/auth -config-path ./config/.env