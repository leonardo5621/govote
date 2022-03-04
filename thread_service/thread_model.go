package thread_service

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (tm *ThreadModel) Init(threadPayload *ThreadCreationPayload) error {
	tm.Title = threadPayload.GetTitle()
	oid, err := primitive.ObjectIDFromHex(threadPayload.GetOwnerUserId())
	if err != nil {
		return status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	tm.OwnerUserId = &oid
	tm.Description = threadPayload.GetDescription()
	return nil
}

type CommentModel struct {
	Id            *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	AuthorUserId  *primitive.ObjectID `json:"authorUserId" bson:"authorUserId,omitempty"`
	ThreadId      *primitive.ObjectID `json:"threadId" bson:"threadId,omitempty"`
	AuthorUserName string     `json: "authorUserName" bson:"authorUserName,omnitempty"`
	Text   string              `json: "text" bson:"text,omnitempty"`
	VoteCount int              `json: "voteCount" bson:"voteCount,omnitempty"`
}

func (cm *CommentModel) Init(commentPayload *CreateCommentPayload) error {
	cm.Text = commentPayload.GetText()
	oid, err := primitive.ObjectIDFromHex(commentPayload.GetAuthorUserId())
	if err != nil {
		return status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	cm.AuthorUserId = &oid
	oidThread, err := primitive.ObjectIDFromHex(commentPayload.GetThreadId())
	if err != nil {
		return status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	cm.ThreadId = &oidThread
	return nil
}

type ThreadFinder struct {
	ThreadId *primitive.ObjectID
}

func (tf *ThreadFinder) Init(threadId string) error {
	oid, err := primitive.ObjectIDFromHex(threadId)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	tf.ThreadId = &oid
	return nil
}

func (tf *ThreadFinder) GetFindOneQuery() *primitive.M {
	if tf.ThreadId == nil {
		return nil
	}
	return &bson.M{"_id": tf.ThreadId }
}

func (tf *ThreadFinder) GetDecodeTargetStruct() interface {} {
	return &ThreadModel{}
}