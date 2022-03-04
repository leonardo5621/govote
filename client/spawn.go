package main

import (
	"fmt"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/leonardo5621/govote/orm"
	"github.com/leonardo5621/govote/thread_service"
	"github.com/leonardo5621/govote/user_service"
	"github.com/leonardo5621/govote/utilities"
)

func SpawnUser(i int)  (*primitive.ObjectID, error) {
	user := &user_service.UserModel{
		FirstName: fmt.Sprintf("client%v", i),
		LastName: fmt.Sprintf("client%v", i),
		UserName: fmt.Sprintf("userclient%v", i),
		Email: fmt.Sprintf("client%v@client.com", i),
	}
	collection := orm.OrmSession.Client.Database("upvote").Collection("user")
	userIdHex, err := orm.Create(user, collection, context.Background())
	if err != nil {
		return nil, utilities.ReturnInternalError(err)
	}
	userId, err := primitive.ObjectIDFromHex(userIdHex)
	if err != nil {
		return nil, fmt.Errorf("Object ID invalid: %v", err)
	}
	return &userId, nil
}

func SpawnThread(i int, userId *primitive.ObjectID)  (*primitive.ObjectID, error) {
	thread := &thread_service.ThreadModel{
		Title: fmt.Sprintf("Thread %v", i),
		Description: fmt.Sprintf("Thread%v description", i),
		OwnerUserId: userId,
	}
	collection := orm.OrmSession.Client.Database("upvote").Collection("thread")
	threadIdHex, err := orm.Create(thread, collection, context.Background())
	if err != nil {
		return nil, utilities.ReturnInternalError(err)
	}
	threadId, err := primitive.ObjectIDFromHex(threadIdHex)
	if err != nil {
		return nil, fmt.Errorf("Object ID invalid: %v", err)
	}
	return &threadId, nil
}