package thread_service

import 	(
	"context"
	"github.com/leonardo5621/govote/utilities"
	"github.com/leonardo5621/govote/orm"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ThreadModel struct {
	Id *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	OwnerUserId *primitive.ObjectID `json:"ownerUserId" bson:"ownerUserId,omitempty"`
	FirmId *primitive.ObjectID `json:"firmId" bson:"firmId,omitempty"`
	Title string `json: "title" bson:"title,omnitempty"`
	Description string `json: "description" bson:"description,omnitempty"`
	Archived bool `json: "archived" bson:"archived,omnitempty"`
	FirmName string `json: "firmName" bson:"firmName,omnitempty"`
	OwnerUserName string `json: "ownerUserName" bson:"ownerUserName,omnitempty"`

}

type ThreadServer struct {
	UnimplementedThreadServiceServer
}

func (s *ThreadServer) GetThread(ctx context.Context, req *GetThreadRequest) (*GetThreadResponse, error) {
	id := req.GetThreadId()
	collection := orm.OrmSession.Client.Database("upvote").Collection("thread")
	thread, err := orm.FindDocument(id, collection, ctx, ThreadModel{})
	if err != nil {
		return nil, utilities.ReturnInternalError(err)
	}
	responseModel, err := orm.ConvertToEquivalentStruct(thread, Thread{})
	if err != nil {
		return nil, utilities.ReturnInternalError(err)
	}
	return &GetThreadResponse{Thread: responseModel.(*Thread)}, nil

}

func (s *ThreadServer) CreateThread(ctx context.Context, req *CreateThreadRequest) (*CreateThreadResponse, error) {
	threadPayload := req.GetThread()
	validationError := threadPayload.ValidateAll()
	if validationError != nil {
		return nil, utilities.ReturnValidationError(validationError)
	}
	thread, err := orm.ConvertToEquivalentStruct(threadPayload, ThreadModel{})
	if err != nil {
		return nil, utilities.ReturnInternalError(err)
	}
	collection := orm.OrmSession.Client.Database("upvote").Collection("thread")
	threadId, err := orm.Create(thread, collection, ctx)
	if err != nil {
		return nil, utilities.ReturnInternalError(err)
	}
	return &CreateThreadResponse{ThreadId: threadId}, nil
}