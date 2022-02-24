package main
import (
	"context"
	"log"
	"fmt"
	"net"
	"os"
	"os/signal"

	upvote "github.com/leonardo5621/govote/upvote_pb"
	"google.golang.org/grpc"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	go func(){
		fmt.Println("Launching server")
		if err := server.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	mongoClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Mongo DB connection error: %v", err)
	}
	errConnect := mongoClient.Connect(context.TODO())
	if errConnect != nil {
		log.Fatalf("Mongo DB connection failed: %v", err)
	}
	collection = mongoClient.Database("upvote").Collection("vote")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch
	fmt.Println("Closing connection")
	if errDisconnect := mongoClient.Disconnect(context.TODO()); errDisconnect != nil {
		log.Fatalf("Mongo DB disconnection failed: %v", err)
	}
	server.Stop()
}