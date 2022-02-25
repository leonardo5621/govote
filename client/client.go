// package client

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"github.com/gin-gonic/gin"
// 	user "github.com/leonardo5621/govote/user_service"
// 	"google.golang.org/grpc"
// )

// var client user.UserServiceClient

// func main() {
// 	conn, err := grpc.Dial("localhost:5005", grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatalf("Connection to gRPC server failed: %v", err)
// 	}
// 	client = user.NewUserServiceClient(conn)
// 	router = gin.Default()

// 	router.POST("/user", createUser)
// 	router.GET("/user", getUser)

// 	router.Run(":8080")

// }