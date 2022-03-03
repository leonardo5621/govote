package user_service

import (
	"context"
	"github.com/leonardo5621/govote/orm"
	"github.com/leonardo5621/govote/utilities"
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

type UserServer struct {
	UnimplementedUserServiceServer
}

func (s *UserServer) GetUser(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error) {
	id := req.GetUserId()
	collection := orm.OrmSession.Client.Database("upvote").Collection("user")
	user, err := orm.FindDocument(id, collection, ctx, UserModel{})
	if err != nil {
		return nil, utilities.ReturnInternalError(err)
	}
	// Pass the parameters returned by the search to the User struct, which
	// Is the defined in the response method
	responseModel, err := orm.ConvertToEquivalentStruct(user, User{})
	if err != nil {
		return nil, utilities.ReturnInternalError(err)
	}
	return &GetUserResponse{User: responseModel.(*User)}, nil

}

func (s *UserServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	userPayload := req.GetUser()
	validationError := userPayload.ValidateAll()
	if validationError != nil {
		return nil, utilities.ReturnValidationError(validationError)
	}
	// Pass the message parameter to a UserModel struct, which
	// Has bson Marshal support
	user, err := orm.ConvertToEquivalentStruct(userPayload, UserModel{})
	if err != nil {
		return nil, utilities.ReturnInternalError(err)
	}
	collection := orm.OrmSession.Client.Database("upvote").Collection("user")
	userId, err := orm.Create(user, collection, ctx)
	if err != nil {
		return nil, utilities.ReturnInternalError(err)
	}
	return &CreateUserResponse{Id: userId}, nil
}
