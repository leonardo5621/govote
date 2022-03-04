package main

import (
	"fmt"
	"log"
	"io"
	"time"
	"context"
	"google.golang.org/grpc"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/leonardo5621/govote/orm"
	"github.com/leonardo5621/govote/upvote_service"
)

type Essentials struct {
	UserIds []*primitive.ObjectID
	ThreadId *primitive.ObjectID
	Votes []*upvote_service.VoteThreadRequest
}

var waiter = make(chan struct{})


var EssentialsRun Essentials

func (e *Essentials) spawnEssentials() error{
	for i := 0; i <= 4; i++ {
		userId, err := SpawnUser(0)
		if err != nil {
			return fmt.Errorf("User spawn error: %v", err)
		}
		e.UserIds = append(e.UserIds, userId)

	}
	userId := e.UserIds[0]
	threadId, err := SpawnThread(0, userId)
	if err != nil {
		return fmt.Errorf("Thread spawn error: %v", err)
	}

	e.ThreadId = threadId
	return nil
}

func (e *Essentials) spawnVotes() {
	for _, uId := range e.UserIds {
		newVote := &upvote_service.VoteThreadRequest{
			UserId: uId.Hex(),
			ThreadId: e.ThreadId.Hex(),
			Votedir: -1,
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
	_ = orm.OpenMongoDBconnection()
	upvoteClient := upvote_service.NewUpvoteServiceClient(connect)
	spawnErr := EssentialsRun.spawnEssentials()
	if spawnErr != nil {
		log.Fatalf("could not spawn essentials: %v", spawnErr)
	}

	EssentialsRun.spawnVotes()
	runVoteStream(upvoteClient)
	
	<-waiter
}

func runVoteStream(upvoteClient upvote_service.UpvoteServiceClient) {
	stream, err := upvoteClient.VoteThread(context.Background())
	if err != nil {
		log.Fatalf("Could not start the voting stream: %v", err)
	}
	go func() {
		for {
			result, err := stream.Recv()
			if err == io.EOF {
				log.Println("EOF")
				close(waiter)
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
			time.Sleep(2000 * time.Millisecond)
		}
		stream.CloseSend()
	}()
}