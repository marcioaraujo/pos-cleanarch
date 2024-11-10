go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest


# para entrar no prompt de comando do evans (Apple M3):
# evans -p 50051 -t ./internal/infra/grpc/protofiles/order.proto


protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto