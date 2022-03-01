package thread_service

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

type ThreadModel struct {
	Id *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title string `json: "title" bson:"title,omnitempty"`
	Description string `json: "description" bson:"description,omnitempty"`
	Owner *user_service.User `json: "owner" bson:"owner,omnitempty"`
	Archived bool `json: "archived" bson:"archived,omnitempty"`
}

type ThreadServer struct {
	UnimplementedThreadServiceServer
}

func (s *ThreadServer) GetThread(ctx context.Context, req *GetThreadRequest) (*GetThreadResponse, error) {
	id := req.GetThreadId()
	searchThread := new(ThreadModel)
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	collection := orm.OrmSession.Client.Database("upvote").Collection("thread")
	res := collection.FindOne(ctx, bson.M{"_id": oid})
	if decodeErr := res.Decode(&searchThread); decodeErr != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find record with Object Id %s: %v", oid, err))
	}
	modelToAssert, err := json.Marshal(searchThread)
	responseModel := new(Thread)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to Json: %v", err))
	}
	json.Unmarshal(modelToAssert, &responseModel)
	return &GetThreadResponse{Thread: responseModel}, nil

}

func (s *ThreadServer) CreateThread(ctx context.Context, req *CreateThreadRequest) (*CreateThreadResponse, error) {
	threadPayload := req.GetThread()
	thread := ThreadModel {
		Title: threadPayload.GetTitle(),
		Owner: threadPayload.GetOwner(),
		Description: threadPayload.GetDescription(),
	}
	collection := orm.OrmSession.Client.Database("upvote").Collection("user")
	threadId, err := orm.Create(thread, collection, ctx)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	return &CreateThreadResponse{ThreadId: threadId}, nil
}
