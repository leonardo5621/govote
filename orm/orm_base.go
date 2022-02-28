package orm;

import (
	"fmt"
	"context"
	//"reflect"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ORMBase interface {
	FindByPk(id string, client *mongo.Client, mongoCtx context.Context) (interface{}, error)
	Create(interface{}, *mongo.Client, context.Context) (string, error)
}

type ORModel struct {
	ModelName string;
	DatabaseName string;
}

func (m ORModel) Create(model interface{}, client *mongo.Client, mongoCtx context.Context) (string, error) {
	collection := client.Database(m.DatabaseName).Collection(m.ModelName)
	res, err := collection.InsertOne(mongoCtx, model)
	if err != nil {
		return "", status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	stringId := res.InsertedID.(primitive.ObjectID)
	return stringId.Hex(), nil
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