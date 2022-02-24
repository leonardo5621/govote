package main
import (
	"context"
	"log"
	"math/rand"
	"net"
	upvote "upvote_pb"
	"google.golang.org/grpc"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	port = ":5005"
)

var collection *mongo.Collection

func main() {
	options := []grpc.ServerOption{}
	server := grpc.NewServer(...options)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("Mongo DB connection error: %v", err)
	}
	errConnect := mongoClient.Connect(context.TODO())
	if errConnect != nil {
		log.Fatal("Mongo DB connection failed: %v", err)
	}
	collection = mongoClient.Database("upvote").Collection("vote")
	s := grpc.NewServer()
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}