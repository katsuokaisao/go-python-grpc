FROM golang:1.20

WORKDIR /root

# PROTOC_VERSION のバージョン戦略は未定
ENV \
  PROTOC_VERSION=25.1

RUN set -ex; \
  # protoc
  apt-get update; \
  apt-get -qq install unzip net-tools; \
  rm -rf /var/lib/apt/lists/*; \
  wget https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-aarch_64.zip -O protobuf.zip; \
  unzip protobuf.zip -d protobuf;

ENV \
  PATH=/root/protobuf/bin:$PATH

WORKDIR /home/src/go/

RUN set -ex; \
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest; \
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

COPY ./src/go/go.mod ./src/go/go.sum ./

RUN go mod tidy

CMD ["go", "run", "main.go"]
