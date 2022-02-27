package main

import (
	"context"
	"fmt"
	"net/http"
	"log"
	"net"
	"os"
	"os/signal"
	"encoding/json"
	"time"

	user_service "github.com/leonardo5621/govote/user_service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/davecgh/go-spew/spew"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
)

const (
	port = ":5005"
)

var client *mongo.Client
var mongoCtx context.Context

type UserServer struct {
	user_service.UnimplementedUserServiceServer
}

func (s *UserServer) CreateUser(ctx context.Context, req *user_service.CreateUserRequest) (*user_service.CreateUserResponse, error) {
	userPayload := req.GetUser()
	user := user_service.UserModel {
		FirstName: userPayload.GetFirstName(),
		LastName: userPayload.GetLastName(),
		Email: userPayload.GetEmail(),
		Activated: userPayload.GetActivated(),
	}
	spew.Dump(user)
	collection := client.Database("upvote").Collection("user")
	fmt.Println("Connection to DB")
	res, err := collection.InsertOne(mongoCtx, user)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	oid := res.InsertedID.(primitive.ObjectID)
	fmt.Println(oid.Hex())
	fmt.Println("Document inserted")
	jsonMongoModel, err := json.Marshal(user)
	userResponseMessage := new(user_service.User)
	json.Unmarshal(jsonMongoModel, userResponseMessage)
	userResponseMessage.Id = oid.Hex()
	return &user_service.CreateUserResponse{User: userResponseMessage}, nil
}

var collection *mongo.Collection

func runHTTPreverseProxy () {
	ctx := context.Background()
  ctx, cancel := context.WithCancel(ctx)
  defer cancel()

  // Register gRPC server endpoint
  // Note: Make sure the gRPC server is running properly and accessible
  mux := runtime.NewServeMux()
  opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
  err := user_service.RegisterUserServiceHandlerFromEndpoint(ctx, mux,  "localhost:5005", opts)
  if err != nil {
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
	user_service.RegisterUserServiceServer(server, &UserServer{})
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

	
	mongoCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	go runHTTPreverseProxy()

	credentials := options.Credential{
		Username: "root",
		Password: "example",
	}
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(credentials)
	client, err = mongo.Connect(mongoCtx, clientOptions)
	if err != nil {
		log.Fatalf("Mongo DB connection failed: %v", err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch
	fmt.Println("Closing connection")
	if errDisconnect := client.Disconnect(mongoCtx); errDisconnect != nil {
		log.Fatalf("Mongo DB disconnection failed: %v", err)
	}
	server.Stop()
}
