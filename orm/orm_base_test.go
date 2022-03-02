package orm

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
	"testing"
)

type MockDBoperator struct{}

func (m MockDBoperator) FindOne(ctx context.Context, filter interface{},
	opts ...*options.FindOneOptions) *mongo.SingleResult {
	return &mongo.SingleResult{}
}

func (m MockDBoperator) InsertOne(ctx context.Context, document interface{},
	opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	id := primitive.NewObjectID()
	return &mongo.InsertOneResult{InsertedID: id}, nil
}

type MockWithBsonSupport struct {
	Id   string `json:"id" bson:"_id,omitempty"`
	Name string `json: "firstName" bson:"firstName,omnitempty"`
}

type MockWithOnlyJsonSupport struct {
	Id   string `json:"id"`
	Name string `json: "firstName"`
}

func TestCreation(t *testing.T) {
	collection := MockDBoperator{}
	ctx := context.Background()
	mockStruct := struct{ Name string }{Name: "Example"}
	returnedId, err := Create(mockStruct, collection, ctx)
	if err != nil {
		t.Errorf("Creation failed: %v", err)
	}
	if reflect.ValueOf(returnedId).Kind() != reflect.ValueOf("").Kind() {
		t.Error("Returned Id should be a string")
	}
}

func TestFindone(t *testing.T) {
	collection := MockDBoperator{}
	ctx := context.Background()
	mockStruct := struct{ Name string }{Name: "Example"}
	foundInstance, err := FindDocument("", collection, ctx, mockStruct)
	if foundInstance != nil {
		t.Errorf("No instance should have been found")
	}
	t.Log(err)
}

func TestStructconversion(t *testing.T) {
	mockJsonSupport := MockWithOnlyJsonSupport{
		Id:   "123",
		Name: "test",
	}
	convertedModel, err := ConvertToEquivalentStruct(mockJsonSupport, MockWithBsonSupport{})
	if err != nil {
		t.Errorf("Conversion failed: %v", err)
	}
	mockWithBson := convertedModel.(*MockWithBsonSupport)
	if mockWithBson.Name != mockJsonSupport.Name || mockWithBson.Id != mockJsonSupport.Id {
		t.Errorf("Wrong values after conversion, Id= %v , Name= %v", mockWithBson.Id, mockWithBson.Name)
	}
}
