package orm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/leonardo5621/govote/utilities"
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

func FindDocument(finder FindOneOperator, collection DocumentFinder, ctx context.Context) (interface{}, error) {
	targetStructPtr := finder.GetDecodeTargetStruct()
	if targetStructPtr == nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Target struct not set"))
	}
	query := finder.GetFindOneQuery()
	if query == nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Find query not set"))
	}
	res := collection.FindOne(ctx, finder.GetFindOneQuery())
	if err := res.Decode(targetStructPtr); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find record with Object Id %v", err))
	}
	return targetStructPtr, nil
}

func CheckDocumentExists(filter interface{}, collection DocumentCounter, ctx context.Context) (bool, error) {
	query := filter.(bson.M)
	count, err := collection.CountDocuments(ctx, query)
	if err != nil {
		return false, utilities.ReturnInternalError(err)
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

func FindByQuery(filter interface{}, collection QueryFinder, ctx context.Context, ) ([]primitive.M, error) {
	query := filter.(bson.M)
	cur, err := collection.Find(ctx, query)
	if err != nil {
		return nil, utilities.ReturnInternalError(err)
	}
	var documentsFiltered []bson.M
	if err = cur.All(ctx, &documentsFiltered); err != nil {
		return nil, utilities.ReturnInternalError(err)
	}
	return documentsFiltered, nil
}

func ConvertStruct(initialStruct interface{}, targetStruct interface{}) (interface{}, error) {
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

func CheckCollectionForDocId(id string, collection *mongo.Collection, ctx context.Context) (error) {
	docId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return utilities.ReturnInternalError(err)
	}
	exists, err := CheckDocumentExists(bson.M{"_id": docId}, collection, ctx)
	if err != nil {
		return utilities.ReturnInternalError(err)
	} else if !exists {
		return status.Errorf(codes.NotFound, fmt.Sprintf("DocId %v not found in collection", docId))
	}
	return nil
}
