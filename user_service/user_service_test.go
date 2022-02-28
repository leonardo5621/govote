package user_service

import (
	"testing"
	"fmt"
	"log"
	"context"
	"net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	RegisterUserServiceServer(s, &UserServer{})
	go func() {
			if err := s.Serve(lis); err != nil {
					log.Fatalf("Server exited with error: %v", err)
			}
	}()
}

func TestUser(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
			t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	req := &CreateUserRequest{
		User: &User{
			FirstName: "Dummy",
			LastName: "Grande Dummy",
			Email: "dummy@dummy.com",
		},
	}
	client := NewUserServiceClient(conn)
  res, err := client.CreateUser(ctx, req)
	if err != nil {
		t.Errorf("User creation failed %v", err)
	}
	t.Log(res.Id)
	fmt.Println("Closing connection")
}


// func TestSayHello(t *testing.T) {
//     ctx := context.Background()
//     conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
//     if err != nil {
//         t.Fatalf("Failed to dial bufnet: %v", err)
//     }
//     defer conn.Close()
//     client := pb.NewGreeterClient(conn)
//     resp, err := client.SayHello(ctx, &pb.HelloRequest{"Dr. Seuss"})
//     if err != nil {
//         t.Fatalf("SayHello failed: %v", err)
//     }
//     log.Printf("Response: %+v", resp)
//     // Test for output here.
// }