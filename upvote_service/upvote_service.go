package upvote_service

import (
	"context"
	"io"
	"log"
	"fmt"
	"time"
	"github.com/leonardo5621/govote/user_service"
	"github.com/leonardo5621/govote/orm"
	"github.com/leonardo5621/govote/utilities"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/davecgh/go-spew/spew"
)

type UpvoteThreadModel struct {
	Id       *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserId   *primitive.ObjectID `json:"userId" bson:"userId,omitempty"`
	ThreadId *primitive.ObjectID `json:"threadId" bson:"threadId,omitempty"`
	Votedir  int          `json: "votedir" bson:"votedir,omnitempty"`
}

type UpvoteServer struct {
	UnimplementedUpvoteServiceServer
}

type NotificationCache struct {
	UserEmailsToNotify []string
	CacheSize int
}

var notificationCache = NotificationCache{
	UserEmailsToNotify: make([]string, 0),
	CacheSize: 2,
}

var notificationChannel chan string

func (u *UpvoteServer) VoteThread(s UpvoteService_VoteThreadServer) error {
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
		threadId, err := primitive.ObjectIDFromHex(votePayload.GetThreadId())
		spew.Dump(threadId)
		if err != nil {
			return utilities.ReturnInternalError(err)
		}
		threadCollection := orm.OrmSession.Client.Database("upvote").Collection("thread")
		exists, err := orm.CheckDocumentExists(bson.M{"_id": threadId}, threadCollection, context.Background())
		if err != nil {
			return utilities.ReturnInternalError(err)
		} else if !exists {
			return status.Errorf(codes.NotFound, fmt.Sprintf("ThreadId %v does not exist", threadId))
		}
		// Check if user exists
		userId, err := primitive.ObjectIDFromHex(votePayload.GetUserId())
		if err != nil {
			return utilities.ReturnInternalError(err)
		}
		userCollection := orm.OrmSession.Client.Database("upvote").Collection("user")
		userExists, err := orm.CheckDocumentExists(bson.M{"_id": userId}, userCollection, context.Background())
		if err != nil {
			return utilities.ReturnInternalError(err)
		} else if !userExists {
			return status.Errorf(codes.NotFound, fmt.Sprintf("UserId %v does not exist", userId))
		}

		// Handle voting
		threadVote, err := orm.ConvertToEquivalentStruct(votePayload, UpvoteThreadModel{})
		valid, err := handleVote(threadVote.(*UpvoteThreadModel))
		if err != nil {
			return utilities.ReturnInternalError(err)
		} else if !valid {
			return status.Errorf(codes.NotFound, fmt.Sprintf("Voting operation failed"))
		}
	}

	return nil
}

func StartNotificationSender(interval time.Duration) {
	notificationChannel = make(chan string)
	go func() {
		clock := time.NewTicker(interval)
		for {
			select {
			// Trigger a notification each five seconds
			case <-clock.C:
				log.Println("Running notification job")
				notificationChannel <- "notify"
			}
		}
	}()
}

func NotificationSender(stream UpvoteService_VoteThreadServer) {
	for range notificationChannel {
		log.Printf("Sending notification to %v users", len(notificationCache.UserEmailsToNotify))
		for _, e := range notificationCache.UserEmailsToNotify {
			err := stream.Send(&VoteThreadResponse{
				Email: e,
				Notification: "Your thread has been upvoted",
			})
			if err != nil {
				log.Fatalf("A notification has not been sent: %v", err)
			}
		}
		notificationCache.UserEmailsToNotify = make([]string, 0)
	}
}

func addUserEmailToCache(userId primitive.ObjectID, userCollection *mongo.Collection) (error) {
	var threadOwner user_service.UserModel
	user := userCollection.FindOne(context.Background(), bson.M{"_id": userId})
	if err := user.Decode(&threadOwner); err != nil {
		return utilities.ReturnInternalError(err)
	}
	notificationCache.UserEmailsToNotify = append(
		notificationCache.UserEmailsToNotify,
		threadOwner.Email,
	)
	if len(notificationCache.UserEmailsToNotify) >= notificationCache.CacheSize {
		notificationChannel <- "notify"
	}
	return nil
}

func handleVote(votePayload *UpvoteThreadModel) (bool, error) {
	db := orm.OrmSession.Client.Database("upvote")
	voteQuery := bson.M{"threadId": votePayload.ThreadId, "userId": votePayload.UserId }
	voteCollection := db.Collection("votes")
	threadCollection := db.Collection("thread")
	foundVote, err := orm.FindByQuery(voteQuery, voteCollection, context.Background())
	if err != nil {
		return false, utilities.ReturnInternalError(err)
	}

	var updateThreadVote int32
	
	if len(foundVote) == 1 {
		userVote := foundVote[0]
		spew.Dump(userVote)
		lastVoteDir, ok := userVote["votedir"]
		if !ok {
			return false, status.Errorf(codes.NotFound, fmt.Sprintf("Vote not set correctly"))
		}
		// Delete existing vote
		deletedRes, deleteErr := voteCollection.DeleteOne(context.TODO(), voteQuery)
		if deleteErr != nil {
			return false, utilities.ReturnInternalError(deleteErr)
		}
		fmt.Printf("Deleted %v Documents!\n", deletedRes.DeletedCount)
		voteDirection := lastVoteDir.(int32)
		if voteDirection == int32(votePayload.Votedir) {
			updateThreadVote = -1 * voteDirection
		}	else {
			updateThreadVote = -2 * voteDirection
			// Create new vote if it is in the opposite direction
			_, err := orm.Create(votePayload, voteCollection, context.Background())
			if err != nil {
				return false, utilities.ReturnInternalError(err)
			}
		}

		// Updating the thread
		res, nil := threadCollection.UpdateOne(context.Background(), bson.M{"_id": votePayload.ThreadId}, bson.D{
			{"$inc", bson.M{"voteCount": updateThreadVote}},
		})
		if err != nil {
			return false, utilities.ReturnInternalError(err)
		}
		fmt.Printf("Updated %v Documents!\n", res.ModifiedCount)
	} else if len(foundVote) == 0 {
		updateThreadVote = int32(votePayload.Votedir)
		// Create new vote if it is in the opposite direction
		_, err := orm.Create(votePayload, voteCollection, context.Background())
		if err != nil {
			return false, utilities.ReturnInternalError(err)
		}
		// Updating the thread
		res, err := threadCollection.UpdateOne(context.Background(), bson.M{"_id": votePayload.ThreadId}, bson.D{
			{"$inc", bson.M{"voteCount": updateThreadVote}},
		})
		if err != nil {
			return false, utilities.ReturnInternalError(err)
		}
		fmt.Printf("Updated %v Documents!\n", res.ModifiedCount)
	} else {
		return false, status.Errorf(codes.NotFound, fmt.Sprintf("Multiple votes found fot the same thread of user %v", votePayload.UserId))
	}
	notificationErr := addUserEmailToCache(*votePayload.UserId, db.Collection("user"))
	if notificationErr != nil {
		return false, utilities.ReturnInternalError(notificationErr)
	}
	return true, nil
}