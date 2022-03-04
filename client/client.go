package client

import (
	"fmt"
	"log"
	"io"
	"time"
	"context"
	"google.golang.org/grpc"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/leonardo5621/govote/orm"
	"github.com/leonardo5621/govote/thread_service"
	"github.com/leonardo5621/govote/user_service"
	"github.com/leonardo5621/govote/upvote_service"
	"github.com/leonardo5621/govote/utilities"
)

type Essentials struct {
	UserIds []*primitive.ObjectID
	ThreadId *primitive.ObjectID
	Votes []*upvote_service.UpvoteThreadModel
}

var EssentialsRun Essentials

func spawnUser(i int)  (*primitive.ObjectID, error) {
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

func spawnThread(i int, userId *primitive.ObjectID)  (*primitive.ObjectID, error) {
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

func (e *Essentials) spawnEssentials() error{
	for i := 0; i <= 4; i++ {
		userId, err := spawnUser(0)
		if err != nil {
			return fmt.Errorf("User spawn error: %v", err)
		}
		e.UserIds = append(e.UserIds, userId)

	}
	userId := e.UserIds[0]
	threadId, err := spawnThread(0, userId)
	if err != nil {
		return fmt.Errorf("Thread spawn error: %v", err)
	}

	e.ThreadId = threadId
	return nil
}

func (e *Essentials) spawnVotes() {
	for _, uId := range e.UserIds {
		newVote := &upvote_service.UpvoteThreadModel{
			UserId: uId,
			ThreadId: e.ThreadId,
			Votedir: 1,
		}
		e.Votes = append(e.Votes, newVote)
	}
}

func main() {
	connect, err := grpc.Dial("localhost:5005", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer connect.Close()
	client := upvote_service.NewUpvoteServiceClient(connect)
	mongoClient := orm.OpenMongoDBconnection()
	spawnErr := EssentialsRun.spawnEssentials()
	if spawnErr != nil {
		log.Fatalf("could not spawn essentials: %v", spawnErr)
	}
	runVoteStream(client)
	if errDisconnect := mongoClient.Disconnect(context.Background()); errDisconnect != nil {
		log.Fatalf("Mongo DB disconnection failed: %v", err)
	}
}

func runVoteStream(upvoteClient upvote_service.UpvoteServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	stream, err := upvoteClient.VoteThread(ctx)
	if err != nil {
		log.Fatalf("Could not start the voting stream: %v", err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			result, err := stream.Recv()
			if err == io.EOF {
				log.Println("EOF")
				close(waitc)
				return
			}
			if err != nil {
				log.Printf("Err: %v", err)
			}
			log.Printf("Notification received: %v", result.GetNotification())
			}
	} ()
	go func () {
		requests := EssentialsRun.Votes
		for _, instruction := range requests {
			if err := stream.Send(instruction); err != nil {
				log.Fatalf("%v.Send(%v) = %v: ", stream, instruction, err)
			}
		}

	}()
}