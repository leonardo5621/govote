package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	upvote "github.com/leonardo5621/govote/upvote_service"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

const (
	port = ":5005"
)

type UpvoteServer struct {
	upvote.UnimplementedUpvoteServiceServer
}

var collection *mongo.Collection

func main() {
	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)
	upvote.RegisterUpvoteServiceServer(server, &UpvoteServer{})
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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	credentials := options.Credential{
		Username: "root",
		Password: "example",
	}
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(credentials)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Mongo DB connection failed: %v", err)
	}

	collection = client.Database("upvote").Collection("vote")
	res, err := collection.InsertOne(ctx, bson.D{{"name", "pi"}, {"value", 3.14159}})
	if err != nil {
		log.Fatalf("Failed insert: %v", err)
	}
	id := res.InsertedID
	fmt.Print(id)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch
	fmt.Println("Closing connection")
	if errDisconnect := client.Disconnect(context.TODO()); errDisconnect != nil {
		log.Fatalf("Mongo DB disconnection failed: %v", err)
	}
	server.Stop()
}
