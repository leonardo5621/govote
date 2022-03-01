package firm_service

import 	(
	"fmt"
	"context"
	"encoding/json"
	"github.com/leonardo5621/govote/user_service"
	"github.com/leonardo5621/govote/orm"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
)

type FirmModel struct {
	Id *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string `json: "name" bson:"name,omnitempty"`
	OwnerUserId *primitive.ObjectID `json: "ownerUserId" bson:"ownerUserId,omnitempty"`
	OwnerUserName string `json: "ownerUserName" bson:"ownerUserName,omnitempty"`
	Enabled bool `json: "enabled" bson:"enabled,omnitempty"`
}

type FirmServer struct {
	UnimplementedFirmServiceServer
}

func (s *FirmServer) GetFirm(ctx context.Context, req *GetFirmRequest) (*GetFirmResponse, error) {
	id := req.GetFirmId()
	searchFirm := new(FirmModel)
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	collection := orm.OrmSession.Client.Database("upvote").Collection("firm")
	res := collection.FindOne(ctx, bson.M{"_id": oid})
	if decodeErr := res.Decode(&searchFirm); decodeErr != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find record with Object Id %s: %v", oid, err))
	}
	modelToAssert, err := json.Marshal(searchFirm)
	responseModel := new(Firm)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to Json: %v", err))
	}
	json.Unmarshal(modelToAssert, &responseModel)
	return &GetFirmResponse{Firm: responseModel}, nil

}

func (s *FirmServer) CreateFirm(ctx context.Context, req *CreateFirmRequest) (*CreateFirmResponse, error) {
	firmPayload := req.GetFirm()
	ownerUserId, err := primitive.ObjectIDFromHex(firmPayload.GetOwnerUserId())
	firm := FirmModel {
		Name: firmPayload.GetName(),
		OwnerUserId: &ownerUserId,
		Enabled: firmPayload.GetEnabled(),
	}
	userInfo := user_service.UserModel{}
	userCollection := orm.OrmSession.Client.Database("upvote").Collection("user")
	res := userCollection.FindOne(ctx, bson.M{"_id": ownerUserId})
	if decodeErr := res.Decode(&userInfo); decodeErr != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find record with Object Id %s: %v", ownerUserId, decodeErr))
	}
	firmCollection := orm.OrmSession.Client.Database("upvote").Collection("firm")
	firm.OwnerUserName = userInfo.LastName
	firmInsertResponse, err := firmCollection.InsertOne(ctx, firm)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	firmId := firmInsertResponse.InsertedID.(primitive.ObjectID)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	return &CreateFirmResponse{FirmId: firmId.Hex()}, nil
}
