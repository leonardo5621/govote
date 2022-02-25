package user_service

import 	(
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserModel struct {
	Id *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FirstName string `json: "firstName" bson:"firstName,omnitempty"`
	LastName string `json: "lastName" bson:"lastName,omnitempty"`
	Email string `json: "email" bson:"email,omnitempty"`
	Activated bool `json: "activated" bson:"activated,omnitempty"`
}

// type UserServer struct {
// 	UnimplementedUserServiceServer
// }

// func (s *UserServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
// 	userPayload := req.GetUser()
// 	user := UserModel {
// 		FirstName: userPayload.GetFirstName(),
// 		LastName: userPayload.GetLastName(),
// 		Email: userPayload.GetEmail(),
// 		Activated: userPayload.GetActivated(),
// 	}
// }
