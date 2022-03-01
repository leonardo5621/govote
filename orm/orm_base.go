package orm;

import (
	"fmt"
	"log"
	"time"
	"context"
	"encoding/json"
	"reflect"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/davecgh/go-spew/spew"
)

type OrmDB struct {
	Client *mongo.Client
}

var OrmSession *OrmDB

func OpenMongoDBconnection() (*mongo.Client){
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	credentials := options.Credential{
		Username: "root",
		Password: "example",
	}
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(credentials)
	Client, err := mongo.Connect(ctx, clientOptions)
	OrmSession = &OrmDB {
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

func ConvertToEquivalentStruct(initialStruct interface{}, targetStruct interface{}) (interface{}, error) {
	targetStructPtr := reflect.New(reflect.TypeOf(targetStruct)).Interface()
	marshalledRequest, err := json.Marshal(initialStruct)
	spew.Dump(marshalledRequest)
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

// func (m ORModel) FindByPk(searchedEntity interface{},  client *mongo.Client, mongoCtx context.Context) (interface{}, error) {
// 	modelPtr := reflect.New(reflect.TypeOf(searchedEntity))
// 	modelMap, ok := searchedEntity.(map[string]interface{})
// 	if !ok {
// 			return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Type assertion map[string]interface{} failed;  got %T", modelMap))
// 	}
// 	id, ok := modelMap["Id"]
// 	if !ok {
// 		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Id not given"))
// 	}
// 	oid, err := primitive.ObjectIDFromHex(id.(string))
// 	if err != nil {
// 		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
// 	}
// 	collection := client.Database(m.DatabaseName).Collection(m.ModelName)
// 	res := collection.FindOne(mongoCtx, bson.M{"_id": oid})
// 	if decodeErr := res.Decode(&modelPtr); decodeErr != nil {
// 		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find record with Object Id %s: %v", oid, err))
// 	}
// 	return modelPtr.Interface(), nil
// }