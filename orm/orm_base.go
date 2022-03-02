package orm

import (
	"context"
	"encoding/json"
	"fmt"
	//"github.com/davecgh/go-spew/spew"
	//"github.com/leonardo5621/govote/utilities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"reflect"
	"time"
)

type OrmDB struct {
	Client *mongo.Client
}

var OrmSession *OrmDB

func OpenMongoDBconnection() *mongo.Client {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	credentials := options.Credential{
		Username: "root",
		Password: "example",
	}
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(credentials)
	Client, err := mongo.Connect(ctx, clientOptions)
	OrmSession = &OrmDB{
		Client: Client,
	}
	if err != nil {
		log.Fatalf("Mongo DB connection failed: %v", err)
	}
	return Client
}

func Create(model interface{}, collection DocumentInserter, ctx context.Context) (string, error) {
	res, err := collection.InsertOne(ctx, model)
	if err != nil {
		return "", status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	stringId := res.InsertedID.(primitive.ObjectID)
	return stringId.Hex(), nil
}

func FindDocument(id string, collection DocumentFinder, ctx context.Context, foundDocumentInstance interface{}) (interface{}, error) {
	targetStructPtr := reflect.New(reflect.TypeOf(foundDocumentInstance)).Interface()
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	res := collection.FindOne(ctx, bson.M{"_id": oid})
	if decodeErr := res.Decode(targetStructPtr); decodeErr != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find record with Object Id %s: %v", oid, err))
	}
	return targetStructPtr, nil
}

func ConvertToEquivalentStruct(initialStruct interface{}, targetStruct interface{}) (interface{}, error) {
	targetStructPtr := reflect.New(reflect.TypeOf(targetStruct)).Interface()
	marshalledRequest, err := json.Marshal(initialStruct)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	json.Unmarshal(marshalledRequest, targetStructPtr)
	return targetStructPtr, nil
}

type DocumentInserter interface {
	InsertOne(ctx context.Context, document interface{},
		opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
}

type DocumentFinder interface {
	FindOne(ctx context.Context, filter interface{},
		opts ...*options.FindOneOptions) *mongo.SingleResult
}

type QueryFinder interface {
	Find(ctx context.Context, filter interface{},
		opts ...*options.FindOptions) (cur *mongo.Cursor, err error)
}
