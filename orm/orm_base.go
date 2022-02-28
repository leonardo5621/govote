package orm;

import (
	"fmt"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ORMBase interface {
	FindByPk(id string) (interface{}, error)
	Create(interface{}, *mongo.Client, context.Context) (string, error)
}

type ORModel struct {
	ModelName string;
	DatabaseName string;
}

type Record struct {
	Id string`json: id`;
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