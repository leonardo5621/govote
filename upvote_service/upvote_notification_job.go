package upvote_service

import (
	"log"
	"fmt"
	"time"
	"context"
	"github.com/leonardo5621/govote/utilities"
	"github.com/leonardo5621/govote/user_service"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NotificationCache struct {
	UserEmailsToNotify []string
	CacheSize int
}

var notificationCache = NotificationCache{
	UserEmailsToNotify: make([]string, 0),
	CacheSize: 2,
}

var notificationChannel chan string

var closeNotificationJob chan string

// Running a timer which after a certain interval
// Emits a trigger in order to send a notification
// To the client
// Such a notification is meant to inform about the upvoting
// Of a thread to a certain user(which could be sent to its e-mail
// as it is done in social media)
func StartNotificationSender(interval time.Duration) {
	notificationChannel = make(chan string)
	closeNotificationJob = make(chan string)
	go func() {
		clock := time.NewTicker(interval)
		for {
			select {
			// Trigger a notification each five seconds
			case <-clock.C:
				log.Println("Running notification job")
				notificationChannel <- "notify"
			case <- closeNotificationJob:
				log.Println("Closing notification job")
				return
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
				Notification: fmt.Sprintf(`
				To: %v
				Your thread has been upvoted!
				`, e),
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
