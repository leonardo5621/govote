package firm_service

import 	(
	"context"
	//"encoding/json"
	"github.com/leonardo5621/govote/utilities"
	"github.com/leonardo5621/govote/orm"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	collection := orm.OrmSession.Client.Database("upvote").Collection("firm")
	firm, err := orm.FindDocument(id, collection, ctx, FirmModel{})
	responseModel, err := orm.ConvertToEquivalentStruct(firm, Firm{})
	if err != nil {
		return nil, utilities.ReturnInternalError(err)
	}
	return &GetFirmResponse{Firm: responseModel.(*Firm)}, nil

}

func (s *FirmServer) CreateFirm(ctx context.Context, req *CreateFirmRequest) (*CreateFirmResponse, error) {
	firmPayload := req.GetFirm()
	validationError := firmPayload.ValidateAll()
	if validationError != nil {
		return nil, utilities.ReturnValidationError(validationError)
	}
	_, err := primitive.ObjectIDFromHex(firmPayload.GetOwnerUserId())
	if err != nil {
		return nil, utilities.ReturnValidationError(err)
	}
	firm, err := orm.ConvertToEquivalentStruct(firmPayload, FirmModel{})
	if err != nil {
		return nil, utilities.ReturnInternalError(err)
	}
	collection := orm.OrmSession.Client.Database("upvote").Collection("firm")
	firmId, err := orm.Create(firm, collection, ctx)
	if err != nil {
		return nil, utilities.ReturnInternalError(err)
	}
	return &CreateFirmResponse{FirmId: firmId}, nil
}
