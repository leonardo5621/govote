package user_service

import 	(
	"fmt"
	"context"
	"encoding/json"
	"github.com/leonardo5621/govote/orm"
	"github.com/leonardo5621/govote/connect_db"
	"go.mongodb.org/mongo-driver/bson"
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

func (s *UserServer) GetUser(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error) {
	id := req.GetUserId()
	searchUser := new(UserModel)
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	collection := connect_db.Client.Database("upvote").Collection("user")
	res := collection.FindOne(connect_db.MongoCtx, bson.M{"_id": oid})
	if decodeErr := res.Decode(&searchUser); decodeErr != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find record with Object Id %s: %v", oid, err))
	}
	modelToAssert, err := json.Marshal(searchUser)
	responseModel := new(User)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to Json: %v", err))
	}
	json.Unmarshal(modelToAssert, &responseModel)
	return &GetUserResponse{User: responseModel}, nil

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
		DatabaseName: "upvote",
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
