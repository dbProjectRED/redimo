buf check lint
protoc --proto_path=redimo/v1 --go_out=plugins=grpc:./v1 --go_opt=paths=source_relative redimo/v1/*.proto

