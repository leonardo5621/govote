package main

import (
	"context"
	"fmt"
	"net/http"
	"log"
	"net"
	"os"
	"os/signal"
	//"time"
	//"github.com/leonardo5621/govote/orm"

	user_service "github.com/leonardo5621/govote/user_service"
	"github.com/leonardo5621/govote/connect_db"
	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	//"github.com/davecgh/go-spew/spew"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/status"
	//"google.golang.org/grpc/codes"
)

const (
	port = ":5005"
)

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
	user_service.RegisterUserServiceServer(server, &user_service.UserServer{})
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

	
	//mongoCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	go runHTTPreverseProxy()
	go connect_db.OpenMongoDBconnection()
	// credentials := options.Credential{
	// 	Username: "root",
	// 	Password: "example",
	// }
	// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(credentials)
	// client, err = mongo.Connect(mongoCtx, clientOptions)
	// if err != nil {
	// 	log.Fatalf("Mongo DB connection failed: %v", err)
	// }

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch
	fmt.Println("Closing connection")
	if errDisconnect := connect_db.Client.Disconnect(connect_db.MongoCtx); errDisconnect != nil {
		log.Fatalf("Mongo DB disconnection failed: %v", err)
	}
	server.Stop()
}
