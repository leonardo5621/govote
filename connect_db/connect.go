package connect_db;

import (
	"context"
	"time"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var MongoCtx context.Context

func OpenMongoDBconnection() {
	var err error
	MongoCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	credentials := options.Credential{
		Username: "root",
		Password: "example",
	}
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(credentials)
	Client, err = mongo.Connect(MongoCtx, clientOptions)
	if err != nil {
		log.Fatalf("Mongo DB connection failed: %v", err)
	}
}