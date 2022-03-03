package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/leonardo5621/govote/orm"
	"github.com/leonardo5621/govote/thread_service"
	"github.com/leonardo5621/govote/user_service"
	"github.com/leonardo5621/govote/upvote_service"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":5005"
)

var collection *mongo.Collection

func runHTTPreverseProxy() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := user_service.RegisterUserServiceHandlerFromEndpoint(ctx, mux, "localhost:5005", opts)
	if err != nil {
		log.Fatalf("Server registration failed: %v", err)
	}
	threadRegisterError := thread_service.RegisterThreadServiceHandlerFromEndpoint(ctx, mux, "localhost:5005", opts)
	if threadRegisterError != nil {
		log.Fatalf("Server registration failed: %v", err)
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	log.Fatalln(http.ListenAndServe("localhost:8081", mux))
	//mux := runtime.NewServeMux()
	//opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	//err := user_service.RegisterUserServiceHandlerServer(ctx, mux, &UserServer{})
	//if err != nil {
	//		log.Fatalf("Server registration failed: %v", err)
	//}
	//log.Fatalln(http.ListenAndServe("localhost:8081", mux))
}

func main() {
	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)
	user_service.RegisterUserServiceServer(server, &user_service.UserServer{})
	thread_service.RegisterThreadServiceServer(server, &thread_service.ThreadServer{})
	upvote_service.RegisterUpvoteServiceServer(server, &upvote_service.UpvoteServer{})
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Server listening at %v", lis.Addr())

	go func() {
		fmt.Println("Launching server")
		if err := server.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	go runHTTPreverseProxy()
	client := orm.OpenMongoDBconnection()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	conn, err := grpc.Dial(":5005", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}
	voteClient := upvote_service.NewUpvoteServiceClient(conn)

	requests := []*upvote_service.VoteThreadRequest{
		&upvote_service.VoteThreadRequest{UserId: "621d9471dca46778d6b66486", ThreadId: "621eca5ff0636d9bad2826bd", Votedir: 1 },
		&upvote_service.VoteThreadRequest{UserId: "621d9471dca46778d6b66486", ThreadId: "621eca5ff0636d9bad2826bd", Votedir: 1 },
		&upvote_service.VoteThreadRequest{UserId: "621d9471dca46778d6b66486", ThreadId: "621eca5ff0636d9bad2826bd", Votedir: -1 },
	}
	stream, err := voteClient.VoteThread(context.Background())
	for _, instruction := range requests {
		if err := stream.Send(instruction); err != nil {
			log.Fatalf("%v.Send(%v) = %v: ", stream, instruction, err)
		}
	}
	if err := stream.CloseSend(); err != nil {
		log.Fatalf("%v.CloseSend() got error %v, want %v", stream, err, nil)
	}

	<-ch
	fmt.Println("Closing connection")
	if errDisconnect := client.Disconnect(context.Background()); errDisconnect != nil {
		log.Fatalf("Mongo DB disconnection failed: %v", err)
	}
	server.Stop()
}
