FROM golang:1.21-alpine

RUN apk update && apk add --no-cache git make protobuf-dev

RUN go install golang.org/x/tools/gopls@latest &&\
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 &&\
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

ENTRYPOINT [ "/bin/bash" ]
