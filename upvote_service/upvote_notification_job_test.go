package upvote_service

import (
	"time"
	"testing"
	"reflect"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestNotifsender(t *testing.T) {
	StartNotificationSender(time.Second)
	var channelValues []string
	i := 0
	for notify := range notificationChannel {
		t.Log(notify)
		channelValues = append(channelValues, notify)
		if i == 2 {
			close(notificationChannel)
			break
		}
		i++
	}
	reflect.DeepEqual([]string{"notify", "notify", "notify"}, channelValues)
}

func TestEmailcache (t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("Email cache test", func(mt *mtest.T) {
		collection := mt.Coll
		oid := primitive.NewObjectID()

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{"_id", primitive.NewObjectID()},
			{"name", "test"},
			{"email", "test@test.com"},
		}))

		err := addUserEmailToCache(oid, collection)
		if err != nil {
			t.Error(err)
		}
		registeredEmails := notificationCache.UserEmailsToNotify
		reflect.DeepEqual(registeredEmails, [...]string{"test@test.com"})
		
	})
}