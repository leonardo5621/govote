package thread_service

import (
	"context"
	//"github.com/davecgh/go-spew/spew"
	"github.com/leonardo5621/govote/orm"
	"github.com/leonardo5621/govote/utilities"
	"github.com/leonardo5621/govote/user_service"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ThreadModel struct {
	Id            *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	OwnerUserId   *primitive.ObjectID `json:"ownerUserId" bson:"ownerUserId,omitempty"`
	Title         string              `json: "title" bson:"title,omnitempty"`
	Description   string              `json: "description" bson:"description,omnitempty"`
	Archived      bool                `json: "archived" bson:"archived,omnitempty"`
	OwnerUserName string              `json: "ownerUserName" bson:"ownerUserName,omnitempty"`
	VoteCount int              `json: "voteCount" bson:"voteCount,omnitempty"`
}

type CommentModel struct {
	Id            *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	AuthorUserId  *primitive.ObjectID `json:"authorUserId" bson:"authorUserId,omitempty"`
	ThreadId      *primitive.ObjectID `json:"threadId" bson:"threadId,omitempty"`
	AuthorUserName string     `json: "authorUserName" bson:"authorUserName,omnitempty"`
	Text   string              `json: "text" bson:"text,omnitempty"`
	VoteCount int              `json: "voteCount" bson:"voteCount,omnitempty"`

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
	// Message parameters validation
	validationError := threadPayload.ValidateAll()
	if validationError != nil {
		return nil, utilities.ReturnValidationError(validationError)
	}
	// Pass the message parameter to a ThreadModel struct, which
	// Has bson Marshal support
	thread, err := orm.ConvertToEquivalentStruct(threadPayload, ThreadModel{})
	assertedThread := thread.(*ThreadModel)
	if err != nil {
		return nil, utilities.ReturnInternalError(err)
	}
	// Get the user name
	userCollection := orm.OrmSession.Client.Database("upvote").Collection("user")
	foundUser, err := orm.FindDocument(threadPayload.GetOwnerUserId(), userCollection, ctx, user_service.UserModel{})
	user := foundUser.(*user_service.UserModel)
	if err != nil {
		return nil, utilities.ReturnInternalError(err)
	}
	assertedThread.OwnerUserName = user.UserName
	// Thread creation
	collection := orm.OrmSession.Client.Database("upvote").Collection("thread")
	threadId, err := orm.Create(assertedThread, collection, ctx)
	if err != nil {
		return nil, utilities.ReturnInternalError(err)
	}
	return &CreateThreadResponse{ThreadId: threadId}, nil
}


func (s *ThreadServer) CreateComment(ctx context.Context, req *CreateCommentRequest) (*CreateCommentResponse, error) {
	commentPayload := req.GetComment()
	// Message parameters validation
	validationError := commentPayload.ValidateAll()
	if validationError != nil {
		return nil, utilities.ReturnValidationError(validationError)
	}
	// Pass the message parameter to a ThreadModel struct, which
	// Has bson Marshal support
	comment, err := orm.ConvertToEquivalentStruct(commentPayload, CommentModel{})
	assertedComment := comment.(*CommentModel)
	if err != nil {
		return nil, utilities.ReturnInternalError(err)
	}
	// Get the user name
	userCollection := orm.OrmSession.Client.Database("upvote").Collection("user")
	foundUser, err := orm.FindDocument(commentPayload.GetAuthorUserId(), userCollection, ctx, user_service.UserModel{})
	user := foundUser.(*user_service.UserModel)
	if err != nil {
		return nil, utilities.ReturnInternalError(err)
	}
	assertedComment.AuthorUserName = user.UserName
	// Check if the thread exist
	threadCollection := orm.OrmSession.Client.Database("upvote").Collection("thread")
	_, errThread := orm.FindDocument(commentPayload.GetThreadId(), threadCollection, ctx, ThreadModel{})
	if errThread != nil {
		return nil, utilities.ReturnInternalError(errThread)
	}
	// Comment creation
	collection := orm.OrmSession.Client.Database("upvote").Collection("comment")
	commentId, err := orm.Create(assertedComment, collection, ctx)
	if err != nil {
		return nil, utilities.ReturnInternalError(err)
	}
	return &CreateCommentResponse{CommentId: commentId}, nil
}

func (s *ThreadServer) GetThreadComments(req *GetThreadRequest, stream ThreadService_GetThreadCommentsServer) error{
	threadId := req.GetThreadId()
	oid, err := primitive.ObjectIDFromHex(threadId)
	if err != nil {
		return utilities.ReturnInternalError(err)
	}
	collection := orm.OrmSession.Client.Database("upvote").Collection("comment")
	query := bson.M{"threadId": oid }
	cursor, err := collection.Find(context.Background(), query)
	if err != nil {
		return utilities.ReturnInternalError(err)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		comment := CommentModel{}
		err := cursor.Decode(&comment)
		if err != nil {
			return utilities.ReturnInternalError(err)
		}
		commentResponse, err := orm.ConvertToEquivalentStruct(comment, Comment{})
		if err != nil {
			return utilities.ReturnInternalError(err)
		}
		stream.Send(&GetThreadCommentsResponse{
			Comment: commentResponse.(*Comment),
		})
	}
	return nil
}

