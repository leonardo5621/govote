package user_service

import (
	"context"
	"github.com/leonardo5621/govote/orm"
	"github.com/leonardo5621/govote/utilities"
)


type UserServer struct {
	UnimplementedUserServiceServer
}

func (s *UserServer) GetUser(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error) {
	userFinder := &UserFinder{}
	err := userFinder.Init(req.GetUserId())
	if err != nil {
		return nil, utilities.ReturnValidationError(err)
	}
	collection := orm.OrmSession.Client.Database("upvote").Collection("user")
	user, err := orm.FindDocument(userFinder, collection, ctx)
	if err != nil {
		return nil, utilities.ReturnInternalError(err)
	}
	// Pass the document returned by the search to the User struct, which
	// Is the defined in the response method
	responseModel, err := orm.ConvertStruct(user, User{})
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
	
	user := &UserModel{}
	user.Init(userPayload)

	collection := orm.OrmSession.Client.Database("upvote").Collection("user")
	userId, err := orm.Create(user, collection, ctx)
	if err != nil {
		return nil, utilities.ReturnInternalError(err)
	}
	return &CreateUserResponse{Id: userId}, nil
}
