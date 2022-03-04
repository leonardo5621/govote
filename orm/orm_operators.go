package orm

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FindOneOperator interface {
	GetFindOneQuery() *primitive.M
	GetDecodeTargetStruct() interface {}
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

type DocumentCounter interface {
	CountDocuments(ctx context.Context, filter interface{},
		opts ...*options.CountOptions) (int64, error)
}