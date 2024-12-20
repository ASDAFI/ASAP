dependencies:
	$(GOGET) -u google.golang.org/grpc
	$(GOGET) -u github.com/grpc-ecosystem/grpc-gateway
	$(GOGET) -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	$(GOGET) -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	$(GOGET) -u github.com/golang/protobuf/protoc-gen-go

build-proto-go:
	find ./proto/messages -name "*.proto" 2>/dev/null | xargs realpath | xargs -I {} protoc $(PROTOC_IMPORT_PATH) --go_out=plugins=grpc:$(BUILD_PROTO_DIRECTORY) {}
	find ./proto/services -name "*.proto" 2>/dev/null | xargs realpath | xargs -I {} protoc $(PROTOC_IMPORT_PATH) --grpc-gateway_out=logtostderr=true:$(BUILD_PROTO_DIRECTORY) --go_out=plugins=grpc:$(BUILD_PROTO_DIRECTORY) {}

farm:
	@echo "Building Farm"
	$(GOBUILD) ./...
	$(GOBUILD) -o $(BINARY_NAME) -v

docker-build:
	docker build -t $(DOCKER_IMAGE_NAME):latest .

all: build-proto-go

all-by-docker:
	docker run --rm --name proto-compiler -it --user $$(id -u):$$(id -g) $(DOCKER_VOLUMES) $(DOCKER_IMAGE_NAME) all

GOCMD=go
GOMOD=GO111MODULE=on $(GOCMD) mod
GOGET=GO111MODULE=on $(GOCMD) "get"

BUILD_DIRECTORY=build
BUILD_PATH=git.carriot.ir/
PROJECT_PATH=git.carriot.ir/farm/

BUILD_PROTO_DIRECTORY=../
GOOGLE_APIS_DIR="$$(find $(GOPATH) -wholename "*github.com/grpc-ecosystem/grpc-gateway*/third_party/googleapis" 2>/dev/null | head -n 1)"
PROTOC_IMPORT_PATH=-I${GOOGLE_APIS_DIR} -I $$PWD/proto -I/usr/local/include
GOBUILD=$(GOCMD) build
BINARY_NAME=farm
BINARY_UNIX=$(BINARY_NAME)_unix
DOCKER_VOLUMES=-v $$PWD:/go/src/$(PROJECT_PATH)
DOCKER_IMAGE_NAME=git.carriot.ir:8000/farm
