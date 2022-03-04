package user_service

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserModel struct {
	Id        *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FirstName string              `json: "firstName" bson:"firstName,omnitempty"`
	LastName  string              `json: "lastName" bson:"lastName,omnitempty"`
	Email     string              `json: "email" bson:"email,omnitempty"`
	UserName  string              `json: "userName" bson:"userName,omnitempty"`
	Activated bool                `json: "activated" bson:"activated,omnitempty"`
}

func (um *UserModel) Init(userPayload *User) error {
	um.FirstName = userPayload.GetFirstName()
	um.LastName = userPayload.GetLastName()
	um.UserName = userPayload.GetUserName()
	um.Email = userPayload.GetEmail()
	return nil
}

type UserFinder struct {
	UserId *primitive.ObjectID
}

func (uf *UserFinder) Init(userId string) error {
	oid, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	uf.UserId = &oid
	return nil
}

func (uf *UserFinder) GetFindOneQuery() *primitive.M {
	if uf.UserId == nil {
		return nil
	}
	return &bson.M{"_id": uf.UserId }
}

func (uf *UserFinder) GetDecodeTargetStruct() interface {} {
	return &UserModel{}
}
