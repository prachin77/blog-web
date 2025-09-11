.PHONY: generate_pb_files client server

generate_pb_files:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	       --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	       proto/auth.proto proto/blog.proto

client:
	cd client && go run client_main.go

server:
	cd server && go run server_main.go
