package upvote_service

import (
	"testing"
	//"log"
	//"io"
	//"context"
	//"time"
	//"net"
	"gotest.tools/v3/assert"
	// "google.golang.org/grpc"
	// "google.golang.org/grpc/test/bufconn"
	//"github.com/leonardo5621/govote/orm"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

const bufSize = 1024 * 1024

//var lis *bufconn.Listener

func TestVotehandler(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("Vote handler test", func(mt *mtest.T) {
		collection := mt.Coll
		oid := primitive.NewObjectID()
		//Down vote again
		newVote := &UpvoteThreadModel{
			Id: &oid,
			Votedir: int32(-1),
		}
		existingVote := bson.M{
			"_id": oid,
			"votedir": int32(-1),
		}
		mt.AddMockResponses(mtest.CreateSuccessResponse())
		downVoteUpdateFactor, err := manageExistingVote(newVote, collection, existingVote)
		if err != nil {
			t.Error(err)
		}
		var expectedDownvoteFactor int32 = 1
		assert.Equal(t, downVoteUpdateFactor, expectedDownvoteFactor)

		//Upvote again
		newVote.Votedir = int32(1)
		existingVote["votedir"] = int32(1)
		upVoteUpdateFactor, err := manageExistingVote(newVote, collection, existingVote)
		if err != nil {
			t.Error(err)
		}
		var expectedUpvoteFactor int32 = -1
		assert.Equal(t, upVoteUpdateFactor, expectedUpvoteFactor)

		//Revert downvote
		mt.AddMockResponses(mtest.CreateSuccessResponse())
		newVote.Votedir = int32(1)
		existingVote["votedir"] = int32(-1)
		revertDownVoteUpdateFactor, err := manageExistingVote(newVote, collection, existingVote)
		if err != nil {
			t.Error(err)
		}
		var expectedRevertDownvoteFactor int32 = 2
		assert.Equal(t, revertDownVoteUpdateFactor, expectedRevertDownvoteFactor)

		//Revert upvote
		mt.AddMockResponses(mtest.CreateSuccessResponse())
		newVote.Votedir = int32(-1)
		existingVote["votedir"] = int32(1)
		revertUpVoteUpdateFactor, err := manageExistingVote(newVote, collection, existingVote)
		if err != nil {
			t.Error(err)
		}
		var expectedRevertUpvoteFactor int32 = -2
		assert.Equal(t, revertUpVoteUpdateFactor, expectedRevertUpvoteFactor)
	})
}

// func init() {
	
//     lis = bufconn.Listen(bufSize)
//     s := grpc.NewServer()
//     RegisterUpvoteServiceServer(s, &UpvoteServer{})
//     go func() {
//         if err := s.Serve(lis); err != nil {
//             log.Fatalf("Server exited with error: %v", err)
//         }
//     }()
// }

// func bufDialer(context.Context, string) (net.Conn, error) {
//     return lis.Dial()
//}

// func TestUpvote(t *testing.T) {
//     ctx := context.Background()
//     conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
//     if err != nil {
//         t.Fatalf("Failed to dial bufnet: %v", err)
//     }
//     defer conn.Close()
// 		closer := make(chan struct{})
//     client := NewUpvoteServiceClient(conn)
//     mongoClient := orm.OpenMongoDBconnection()
//     requests := []*VoteThreadRequest{
// 			{UserId: "621d9471dca46778d6b66486", ThreadId: "621eca5ff0636d9bad2826bd", Votedir: 1 },
// 			{UserId: "621d9471dca46778d6b66486", ThreadId: "621eca5ff0636d9bad2826bd", Votedir: 1 },
// 			{UserId: "621d9471dca46778d6b66486", ThreadId: "621eca5ff0636d9bad2826bd", Votedir: -1 },
// 		}
// 		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 		defer cancel()
// 		stream, err := client.VoteThread(context.Background())

// 		go func() {
// 			for {
// 				result, err := stream.Recv()
// 				if err == io.EOF {
// 					log.Println("EOF")
// 					close(closer)
// 					return
// 				}
// 				if err != nil {
// 					log.Printf("Err: %v", err)
// 				}
// 				log.Printf("output: %v", result.GetNotification())
// 				}
// 		} ()
		
// 		for _, instruction := range requests {
// 			t.Log(instruction)
// 			if err := stream.Send(instruction); err != nil {
// 				log.Fatalf("%v.Send(%v) = %v: ", stream, instruction, err)
// 				t.Errorf("%v.Send(%v) = %v: ", stream, instruction, err)
// 			}
// 		}
// 		if err := stream.CloseSend(); err != nil {
// 			log.Fatalf("%v.CloseSend() got error %v, want %v", stream, err, nil)
// 			t.Errorf("%v.CloseSend() got error %v, want %v", stream, err, nil)
// 		}
// 		if errDisconnect := mongoClient.Disconnect(context.Background()); errDisconnect != nil {
// 			log.Fatalf("Mongo DB disconnection failed: %v", err)
// 			t.Errorf("Mongo DB disconnection failed: %v", err)
// 		}
// }