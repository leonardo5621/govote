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

type MockWithBsonSupport struct {
	Id   *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string `json: "firstName" bson:"firstName,omnitempty"`
}

type MockWithOnlyJsonSupport struct {
	Id   string `json:"id"`
	Name string `json: "firstName"`
}

type MockFinder struct {
	Id *primitive.ObjectID
}

func (mf *MockFinder) Init(Id string) error {
	oid, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	mf.Id = &oid
	return nil
}

func (mf *MockFinder) GetFindOneQuery() *primitive.M {
	if mf.Id == nil {
		return nil
	}
	return &bson.M{"_id": mf.Id }
}

func (mf *MockFinder) GetDecodeTargetStruct() interface {} {
	return &MockWithBsonSupport{}
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
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("Find one test", func(mt *mtest.T) {
		collection := mt.Coll
		oid := primitive.NewObjectID()
		ctx := context.Background()
		mockFinder := &MockFinder{}
		mockFinder.Id = &oid

		doc, err := FindDocument(mockFinder, collection, ctx)
		if err != nil && doc == nil{
			t.Log(err)
		} else {
			t.Errorf("Result should have been nil")
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "mock.test", mtest.FirstBatch, bson.D{
			{"_id", oid},
			{"name", "test"},
		}))

		foundDoc, err := FindDocument(mockFinder, collection, ctx)
		if err != nil {
			t.Error(err)
		}
		reflect.DeepEqual(&MockWithBsonSupport{
			Id: &oid, Name: "test",
		}, foundDoc)
		
	})
}

func TestStructconversion(t *testing.T) {
	oid := primitive.NewObjectID()
	mockJsonSupport := MockWithOnlyJsonSupport{
		Id:   oid.Hex(),
		Name: "test",
	}
	convertedModel, err := ConvertStruct(mockJsonSupport, MockWithBsonSupport{})
	if err != nil {
		t.Errorf("Conversion failed: %v", err)
	}
	mockWithBson := convertedModel.(*MockWithBsonSupport)
	if mockWithBson.Name != mockJsonSupport.Name {
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
		firstSearchResult, err := FindByQuery(searchQuery, collection, ctx)

		if err != nil && firstSearchResult == nil{
			t.Log(err)
		} else {
			t.Errorf("Result should have been nil")
		}

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

func TestCheckforid(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("Check for Id test", func(mt *mtest.T) {
		collection := mt.Coll
		oid := primitive.NewObjectID()
		ctx := context.Background()

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "mock.test", mtest.FirstBatch, bson.D{{"_id", oid}}))
		err := CheckCollectionForDocId("123", collection, ctx)
		if err != nil {
			t.Log(err)
		} else {
			t.Errorf("A not valid objectID error should have been returned")
		}

		foundDoc := CheckCollectionForDocId(oid.Hex(), collection, ctx)
		if foundDoc == nil {
			t.Error("Error should have been nil")
		}
	})
}
