package user_service

import 	(
	"fmt"
	"context"
	"github.com/leonardo5621/govote/orm"
	"github.com/leonardo5621/govote/connect_db"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
)

type UserModel struct {
	Id *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FirstName string `json: "firstName" bson:"firstName,omnitempty"`
	LastName string `json: "lastName" bson:"lastName,omnitempty"`
	Email string `json: "email" bson:"email,omnitempty"`
	Activated bool `json: "activated" bson:"activated,omnitempty"`
}

type UserServer struct {
	UnimplementedUserServiceServer
}

func (s *UserServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	userPayload := req.GetUser()
	user := UserModel {
		FirstName: userPayload.GetFirstName(),
		LastName: userPayload.GetLastName(),
		Email: userPayload.GetEmail(),
		Activated: userPayload.GetActivated(),
	}
	 
	userModel := orm.ORModel {
		ModelName: "user",
		DatabaseName: "upvte",
	}
	userId, err := userModel.Create(user, connect_db.Client, ctx)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	return &CreateUserResponse{Id: userId}, nil
}
