FROM golang:1.16.2

WORKDIR /go/src/git.carriot.ir/farm/

RUN apt-get update && apt-get install -y \
    protobuf-compiler \
 && rm -rf /var/lib/apt/lists/*

RUN wget https://github.com/grpc/grpc-web/releases/download/1.2.0/protoc-gen-grpc-web-1.2.0-linux-x86_64 \
  && mv protoc-gen-grpc-web-1.2.0-linux-x86_64 /usr/local/bin/protoc-gen-grpc-web \
  && chmod a+x /usr/local/bin/protoc-gen-grpc-web

COPY Makefile ./

RUN GO111MODULE=on go mod init

RUN export GOPROXY=https://goproxy.io && make dependencies

VOLUME /go/src/git.carriot.ir/farm/

ENTRYPOINT ["/usr/bin/make"]

