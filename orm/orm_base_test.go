package orm

import (
	"fmt"
	"context"
	"reflect"
	"testing"
	"gotest.tools/v3/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
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

func (m MockDBoperator) CountDocuments(ctx context.Context, filter interface{},
	opts ...*options.CountOptions) (int64, error) {
	bsonQuery := filter.(bson.M)
	isInDB, ok := bsonQuery["isInDB"]
	if !ok {
		return 0, status.Errorf(codes.NotFound, fmt.Sprintf("Value not set correctly"))
	}
	dbExists := isInDB.(bool)
	if dbExists {
		return 1, nil
	} else {
		return 0, nil
	}
}

func (m MockDBoperator) Find(ctx context.Context, filter interface{},
	opts ...*options.FindOptions) (cur *mongo.Cursor, err error) {
		foundDoc, err := bson.Marshal(bson.D{{"find", "this"}})
		if err != nil {
			return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Marshaling failed: %v", err))
		}
		return &mongo.Cursor{
			Current: foundDoc,
		}, nil
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
	_, err := FindDocument("621eca5ff0636d9bad2826bd", collection, ctx, mockStruct)
	if err == nil {
		t.Errorf("Search failed")
	}
}

func TestStructconversion(t *testing.T) {
	mockJsonSupport := MockWithOnlyJsonSupport{
		Id:   "123",
		Name: "test",
	}
	convertedModel, err := ConvertToStructWithBsonSupport(mockJsonSupport, MockWithBsonSupport{})
	if err != nil {
		t.Errorf("Conversion failed: %v", err)
	}
	mockWithBson := convertedModel.(*MockWithBsonSupport)
	if mockWithBson.Name != mockJsonSupport.Name || mockWithBson.Id != mockJsonSupport.Id {
		t.Errorf("Wrong values after conversion, Id= %v , Name= %v", mockWithBson.Id, mockWithBson.Name)
	}
}

func TestDocumentcounter(t *testing.T) {
	searchQuery := bson.M{"isInDB": true}
	ctx := context.Background()
	collection := MockDBoperator{}
	check, err := CheckDocumentExists(searchQuery, collection, ctx)
	if err != nil || check == false {
		t.Errorf("Check document failed to find element")
	}
	searchQueryNotFind := bson.M{"isInDB": false}
	checkNotInDB, err := CheckDocumentExists(searchQueryNotFind, collection, ctx)
	if err != nil || checkNotInDB == true {
		t.Errorf("Check document failed, no element should have been found")
	}
}

func TestQueryfinder(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	ctx := context.Background()
	defer mt.Close()
	mt.Run("Query test", func(mt *mtest.T) {
		collection := mt.Coll
		oId1 := primitive.NewObjectID()
		searchQuery := bson.M{"name": "test"}

		mock1 := mtest.CreateCursorResponse(1, "mock.test", mtest.FirstBatch, bson.D{
			{"_id", oId1},
			{"name", "test"},
		})
		killCursors := mtest.CreateCursorResponse(0, "mock.test", mtest.NextBatch)
		mt.AddMockResponses(mock1, killCursors)

		res, err := FindByQuery(searchQuery, collection, ctx)
		if err != nil {
			t.Error(err)
		}
		foundName := res[0]["name"]
		assert.Equal(t, "test", foundName)
	})
}
