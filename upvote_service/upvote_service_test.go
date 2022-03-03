package upvote_service

import (
	"testing"
	"log"
	"io"
	"context"
	"time"
	"net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"github.com/leonardo5621/govote/orm"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	
    lis = bufconn.Listen(bufSize)
    s := grpc.NewServer()
    RegisterUpvoteServiceServer(s, &UpvoteServer{})
    go func() {
        if err := s.Serve(lis); err != nil {
            log.Fatalf("Server exited with error: %v", err)
        }
    }()
}

func bufDialer(context.Context, string) (net.Conn, error) {
    return lis.Dial()
}

func TestUpvote(t *testing.T) {
    ctx := context.Background()
    conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
    if err != nil {
        t.Fatalf("Failed to dial bufnet: %v", err)
    }
    defer conn.Close()
		closer := make(chan struct{})
    client := NewUpvoteServiceClient(conn)
    mongoClient := orm.OpenMongoDBconnection()
    requests := []*VoteThreadRequest{
			{UserId: "621d9471dca46778d6b66486", ThreadId: "621eca5ff0636d9bad2826bd", Votedir: 1 },
			{UserId: "621d9471dca46778d6b66486", ThreadId: "621eca5ff0636d9bad2826bd", Votedir: 1 },
			{UserId: "621d9471dca46778d6b66486", ThreadId: "621eca5ff0636d9bad2826bd", Votedir: -1 },
		}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		stream, err := client.VoteThread(context.Background())

		go func() {
			for {
				result, err := stream.Recv()
				if err == io.EOF {
					log.Println("EOF")
					close(closer)
					return
				}
				if err != nil {
					log.Printf("Err: %v", err)
				}
				log.Printf("output: %v", result.GetNotification())
				}
		} ()
		
		for _, instruction := range requests {
			t.Log(instruction)
			if err := stream.Send(instruction); err != nil {
				log.Fatalf("%v.Send(%v) = %v: ", stream, instruction, err)
				t.Errorf("%v.Send(%v) = %v: ", stream, instruction, err)
			}
		}
		<- closer
		if err := stream.CloseSend(); err != nil {
			log.Fatalf("%v.CloseSend() got error %v, want %v", stream, err, nil)
			t.Errorf("%v.CloseSend() got error %v, want %v", stream, err, nil)
		}
		if errDisconnect := mongoClient.Disconnect(context.Background()); errDisconnect != nil {
			log.Fatalf("Mongo DB disconnection failed: %v", err)
			t.Errorf("Mongo DB disconnection failed: %v", err)
		}
}