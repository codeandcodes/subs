protoc -I . --proto_path=./protos --proto_path=$HOME/workspace/googleapis \
    --openapiv2_out ./backend/httpserver/static/ \
    --openapiv2_opt logtostderr=true \
    protos/api.proto

protoc --proto_path=./protos --proto_path=$HOME/workspace/googleapis \
    --go_out=protos --go_opt=paths=source_relative \
    --go-grpc_out=protos --go-grpc_opt=paths=source_relative \
    --grpc-gateway_out=protos --grpc-gateway_opt=paths=source_relative \
    protos/*.proto
