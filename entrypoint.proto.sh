protoc --go_out=. \
    --go-grpc_out=. \
    --go-grpc_opt=paths=source_relative \
    proto/*/*.proto
