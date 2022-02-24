create:
	protoc --proto_path=proto proto/*.proto --go_out=server/
	protoc --proto_path=proto proto/*.proto --go-grpc_out=server/

