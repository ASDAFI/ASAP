postgres:
	cd ./db/postgres; bash setup.sh

mqtt:
	cd ./mqtt; bash setup.sh

add-device:
	docker exec mqtt mosquitto_passwd -b /mosquitto/config/pwfile $(serial) $(password)

proto-dependencies:
	$(GOGET) github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	$(GOGET) github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	$(GOGET) github.com/golang/protobuf/protoc-gen-go

build-proto-go:
	find ./proto/messages -name "*.proto" 2>/dev/null | xargs realpath | xargs -I {} protoc $(PROTOC_IMPORT_PATH) --go_out=plugins=grpc:$(BUILD_PROTO_DIRECTORY) {}
	find ./proto/services -name "*.proto" 2>/dev/null | xargs realpath | xargs -I {} protoc $(PROTOC_IMPORT_PATH) --grpc-gateway_out=logtostderr=true:$(BUILD_PROTO_DIRECTORY) --swagger_out=logtostderr=true:$(BUILD_PROTO_DIRECTORY) --go_out=plugins=grpc:$(BUILD_PROTO_DIRECTORY) {}

farm:
	@echo "Building Farm"
	$(GOBUILD) ./...
	$(GOBUILD) -o $(BINARY_NAME) -v

GOCMD=go
GOMOD=GO111MODULE=on $(GOCMD) mod
GOGET=GO111MODULE=on $(GOCMD) "get"
BUILD_PROTO_DIRECTORY=../
GOPATH=/home/$(USER)/go/pkg/mod
GOOGLE_APIS_DIR="$$(find $(GOPATH) -wholename "*github.com/grpc-ecosystem/grpc-gateway*/third_party/googleapis" 2>/dev/null | head -n 1)"
PROTOC_IMPORT_PATH=-I${GOOGLE_APIS_DIR} -I $$PWD/proto -I/usr/local/include
GOBUILD=$(GOCMD) build
BINARY_NAME=farm
BINARY_UNIX=$(BINARY_NAME)_unix
