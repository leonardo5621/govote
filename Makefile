create_upvote_service:
	protoc --proto_path=protobuffers protobuffers/upvote.proto --go_out=.
	protoc --proto_path=protobuffers protobuffers/upvote.proto --go-grpc_out=.

create_firm_service:
	protoc --proto_path=protobuffers protobuffers/firm.proto --go_out=.
	protoc --proto_path=protobuffers protobuffers/firm.proto --go-grpc_out=.

create_user_service:
	protoc --proto_path=protobuffers protobuffers/user.proto --go_out=.
	protoc --proto_path=protobuffers protobuffers/user.proto --go-grpc_out=.

create_thread_service:
	protoc --proto_path=protobuffers protobuffers/thread.proto --go_out=.
	protoc --proto_path=protobuffers protobuffers/thread.proto --go-grpc_out=.

get-plugins:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
	export PATH="$PATH:$(go env GOPATH)/bin"