create_upvote_service:
	protoc --proto_path=protobuffers protobuffers/upvote.proto --go_out=.
	protoc --proto_path=protobuffers protobuffers/upvote.proto --go-grpc_out=.

create_firm_service:
	protoc \
  -I . \
  -I=${GOPATH}/src \
  -I=${GOPATH}/src/github.com/protoc-gen-validate \
  --go_out=":." \
  --validate_out="lang=go:." \
	--go-grpc_out=. \
	--grpc-gateway_out firm_service \
	 --grpc-gateway_opt logtostderr=true \
	--grpc-gateway_opt paths=source_relative \
	./protobuffers/firm.proto
	mv firm_service/protobuffers/firm.* firm_service
	rm -r firm_service/protobuffers	

create_user_service:
	protoc \
  -I . \
  -I=${GOPATH}/src \
  -I=${GOPATH}/src/github.com/protoc-gen-validate \
  --go_out=":." \
  --validate_out="lang=go:." \
	--go-grpc_out=. \
	--grpc-gateway_out user_service \
	 --grpc-gateway_opt logtostderr=true \
	--grpc-gateway_opt paths=source_relative \
	./protobuffers/user.proto
	mv user_service/protobuffers/user.* user_service
	rm -r user_service/protobuffers	

create_thread_service:
	protoc \
  -I . \
  -I=${GOPATH}/src \
  -I=${GOPATH}/src/github.com/protoc-gen-validate \
  --go_out=":." \
  --validate_out="lang=go:." \
	--go-grpc_out=. \
	--grpc-gateway_out thread_service \
	 --grpc-gateway_opt logtostderr=true \
	--grpc-gateway_opt paths=source_relative \
	./protobuffers/thread.proto
	mv thread_service/protobuffers/thread.* thread_service
	rm -r thread_service/protobuffers

get-plugins:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
	export PATH="$PATH:$(go env GOPATH)/bin"