package user_service

import (
	"testing"
	"reflect"
	"gotest.tools/v3/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreationvalidation(t *testing.T) {
	user := User{
		FirstName: "tester",
		LastName:  "tester last",
		Email:     "testemail",
	}
	validationErrEmail := user.Validate()
	if validationErrEmail == nil {
		t.Error("Email validation failed")
	}
	user2 := User{
		FirstName: "",
		LastName:  "tester last",
		Email:     "test@email.com",
	}
	validationErrName := user2.Validate()
	if validationErrName == nil {
		t.Error("First name validation failed")
	}
}

func TestGetvalidation(t *testing.T) {
	userReq := &GetUserRequest{
		UserId: "12323@$!!",
	}
	idValidation := userReq.Validate()
	if idValidation == nil {
		t.Error("Id validation failed, special characters passed")
	}
	userReqCorrect := &GetUserRequest{
		UserId: "123232wrsedSASF",
	}
	idValidationErr := userReqCorrect.Validate()
	if idValidationErr != nil {
		t.Errorf("Id validation failed, allowed characters were not recognized, %v", idValidationErr)
	}
}

func TestInituser(t *testing.T) {
	userPayload := &User{
		FirstName: "test1",
		LastName: "test2",
		Email: "test@test.com",
		UserName: "None",
	}
	user := &UserModel{}
	err := user.Init(userPayload)
	if err != nil {
		t.Error(err)
	}
}

func TestUserfinder(t *testing.T) {
	oid := primitive.NewObjectID()
	userFinder := &UserFinder{}
	err := userFinder.Init(oid.Hex())
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, reflect.DeepEqual(userFinder.UserId, &oid), true)	
	assert.Equal(t, reflect.DeepEqual(*userFinder.GetFindOneQuery(), bson.M{"_id": userFinder.UserId }), true)
	assert.Equal(
		t,
		reflect.TypeOf(userFinder.GetDecodeTargetStruct()).Kind(),
		reflect.TypeOf(&UserModel{}).Kind(),
	)
}
