buf check lint
protoc --proto_path=dbProjectRED/redimo/v1 --go_out=plugins=grpc:./v1 --go_opt=paths=source_relative dbProjectRED/redimo/v1/*.proto

