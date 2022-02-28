package firm_service

import 	(
	"fmt"
	"context"
	"encoding/json"
	"github.com/leonardo5621/govote/user_service"
	"github.com/leonardo5621/govote/connect_db"
	"github.com/leonardo5621/govote/orm"
	"github.com/davecgh/go-spew/spew"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
)

type FirmModel struct {
	Id *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string `json: "name" bson:"name,omnitempty"`
	Owner *user_service.User `json: "owner" bson:"owner,omnitempty"`
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
	collection := connect_db.Client.Database("upvote").Collection("firm")
	res := collection.FindOne(connect_db.MongoCtx, bson.M{"_id": oid})
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
	firm := FirmModel {
		Name: firmPayload.GetName(),
		Owner: firmPayload.GetOwner(),
		Enabled: firmPayload.GetEnabled(),
	}
	spew.Dump(firm)
	firmModel := orm.ORModel {
		ModelName: "firm",
		DatabaseName: "upvote",
	}
	firmId, err := firmModel.Create(firm, connect_db.Client, ctx)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	return &CreateFirmResponse{FirmId: firmId}, nil
}
