package upvote_service

import (
	"context"
	"io"
	"log"
	"fmt"
	"time"
	"github.com/leonardo5621/govote/orm"
	"github.com/leonardo5621/govote/utilities"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type UpvoteServer struct {
	UnimplementedUpvoteServiceServer
}


func (u *UpvoteServer) VoteThread(s UpvoteService_VoteThreadServer) error {
	db := orm.OrmSession.Client.Database("upvote")
	StartNotificationSender(10 * time.Second)
	go NotificationSender(s)
	for {
		votePayload, err := s.Recv()
		// Check if the client is open
		if err == io.EOF {
			log.Println("Client has closed connection")
			break
		}
		// Validate message params
		validationError := votePayload.ValidateAll()
		if validationError != nil {
			return utilities.ReturnValidationError(validationError)
		}

		// Check if thread exists
		threadExistErr := orm.CheckCollectionForDocId(votePayload.GetThreadId(), db.Collection("thread"), context.Background())
		if threadExistErr != nil {
			return utilities.ReturnValidationError(threadExistErr)
		} 
		// Check if user exists
		userExistErr := orm.CheckCollectionForDocId(votePayload.GetUserId(), db.Collection("user"), context.Background())
		if userExistErr != nil {
			return utilities.ReturnValidationError(userExistErr)
		} 

		// Handle voting
		threadVote, err := orm.ConvertStruct(votePayload, UpvoteThreadModel{})
		valid, err := handleVote(threadVote.(*UpvoteThreadModel))
		if err != nil {
			return utilities.ReturnInternalError(err)
		} else if !valid {
			return status.Errorf(codes.NotFound, fmt.Sprintf("Voting operation failed"))
		}
	}

	return nil
}

func handleVote(newVote Upvote) (bool, error) {
	db := orm.OrmSession.Client.Database("upvote")
	voteQuery := newVote.GetSearchQuery()
	voteCollection := db.Collection("votes")
	threadCollection := db.Collection("thread")
	foundVote, err := orm.FindByQuery(voteQuery, voteCollection, context.Background())
	if err != nil {
		return false, utilities.ReturnInternalError(err)
	}
	
	if len(foundVote) == 1 {
		userVote := foundVote[0]
		// Existing vote manager function
		updateThreadVote, err := manageExistingVote(newVote, voteCollection, userVote)
		if err != nil {
			return false, utilities.ReturnInternalError(err)
		}
		// Delete existing vote
		deletedRes, deleteErr := voteCollection.DeleteOne(context.Background(), voteQuery)
		if deleteErr != nil {
			return false, utilities.ReturnInternalError(deleteErr)
		}
		fmt.Printf("Deleted %v Documents!\n", deletedRes.DeletedCount)
		

		// Updating the thread
		res, nil := threadCollection.UpdateOne(context.Background(), bson.M{"_id": newVote.GetResourceId()}, bson.D{
			{"$inc", bson.M{"voteCount": updateThreadVote}},
		})
		if err != nil {
			return false, utilities.ReturnInternalError(err)
		}
		fmt.Printf("Updated %v Documents!\n", res.ModifiedCount)
	} else if len(foundVote) == 0 {
		updateThreadVote := newVote.GetVotedir()
		// Create new vote if it is in the opposite direction
		_, err := orm.Create(newVote, voteCollection, context.Background())
		if err != nil {
			return false, utilities.ReturnInternalError(err)
		}
		// Updating the thread
		res, err := threadCollection.UpdateOne(context.Background(), bson.M{"_id": newVote.GetResourceId()}, bson.D{
			{"$inc", bson.M{"voteCount": updateThreadVote}},
		})
		if err != nil {
			return false, utilities.ReturnInternalError(err)
		}
		fmt.Printf("Updated %v Documents!\n", res.ModifiedCount)
	} else {
		return false, status.Errorf(codes.NotFound, fmt.Sprintf("Multiple votes found fot the same thread of user %v", newVote.GetUserId()))
	}
	notificationErr := addUserEmailToCache(*newVote.GetUserId(), db.Collection("user"))
	if notificationErr != nil {
		return false, utilities.ReturnInternalError(notificationErr)
	}
	return true, nil
}

func manageExistingVote(
	votePayload Upvote,
	collection *mongo.Collection,
	existingVote map[string]interface {},
	) (int32, error) {
		var updateVote int32

		// Check vote account consistency
		lastVoteDir, ok := existingVote["votedir"]
		if !ok {
			return 0, status.Errorf(codes.NotFound, fmt.Sprintf("Vote not set correctly"))
		}

		voteDirection := lastVoteDir.(int32)
		// Determine the update factor of the corresponding resource
		if voteDirection == votePayload.GetVotedir() {
			updateVote = -1 * voteDirection
		}	else {
			updateVote = -2 * voteDirection
			// Create new vote if it is in the opposite direction
			_, err := orm.Create(votePayload, collection, context.Background())
			if err != nil {
				return 0, utilities.ReturnInternalError(err)
			}
		}
		return updateVote, nil
}