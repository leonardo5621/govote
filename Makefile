create:
	protoc --proto_path=upvote_pb upvote_pb/*.proto --go_out=.
	protoc --proto_path=upvote_pb upvote_pb/*.proto --go-grpc_out=.

get-plugins:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
	export PATH="$PATH:$(go env GOPATH)/bin"