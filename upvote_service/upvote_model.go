package upvote_service

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Upvote interface {
	GetVotedir() int32
	GetSearchQuery() primitive.M
	GetResourceId() *primitive.ObjectID
	GetUserId() *primitive.ObjectID 
} 

type UpvoteThreadModel struct {
	Id       *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserId   *primitive.ObjectID `json:"userId" bson:"userId,omitempty"`
	ThreadId *primitive.ObjectID `json:"threadId" bson:"threadId,omitempty"`
	Votedir  int32         `json: "votedir" bson:"votedir,omnitempty"`
}

func (u *UpvoteThreadModel) GetVotedir() int32 {
	if u != nil {
		return 0
	}
	return u.Votedir
}

func (u *UpvoteThreadModel) GetSearchQuery() primitive.M {
	if u != nil {
		return nil
	}
	return bson.M{"threadId": u.ThreadId, "userId": u.UserId }
}

func (u *UpvoteThreadModel) GetResourceId() *primitive.ObjectID {
	if u != nil {
		return nil
	}
	return u.ThreadId
}

func (u *UpvoteThreadModel) GetUserId() *primitive.ObjectID {
	if u != nil {
		return nil
	}
	return u.UserId
}