package user_service

import 	(
	"fmt"
	"context"
	"encoding/json"
	"github.com/leonardo5621/govote/orm"
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
	collection := orm.OrmSession.Client.Database("upvote").Collection("user")
	res := collection.FindOne(ctx, bson.M{"_id": oid})
	if decodeErr := res.Decode(&searchUser); decodeErr != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find record with Object Id %s: %v", oid, err))
	}
	
	modelToAssert, err := json.Marshal(searchUser)
	responseModel := new(User)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Conversion failed: %v", err))
	}
	json.Unmarshal(modelToAssert, &responseModel)
	return &GetUserResponse{User: responseModel}, nil

}

func (s *UserServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	userPayload := req.GetUser()
	validationError := userPayload.ValidateAll()
	if validationError != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Invalid parameters: %v", validationError),
		)
	}
	user, err := orm.ConvertToEquivalentStruct(userPayload, UserModel{})
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Error: %v", err),
		)
	}
	collection := orm.OrmSession.Client.Database("upvote").Collection("user")
	userId, err := orm.Create(user, collection, ctx)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	return &CreateUserResponse{Id: userId}, nil
}
