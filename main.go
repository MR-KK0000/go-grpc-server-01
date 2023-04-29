package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/MR-KK0000/go-grpc-server-01/proto"
	"github.com/gin-gonic/gin"
)

// Server implements the gRPC service interface
type Server struct {
	pb.UnimplementedUserServiceServer
}

// ExampleMethod is a gRPC method that returns an example response
func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	// Your implementation here
	fmt.Println(req.Email)
	fmt.Println(req.Name)
	return &pb.CreateUserResponse{UserId: 123}, nil
}
func (s *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	// Your implementation here
	fmt.Println(req.UserId)
	return &pb.GetUserResponse{Email: "example@ex.com", Name: "test"}, nil
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Create a new gRPC server
	server := grpc.NewServer()

	// Register the service with the server
	pb.RegisterUserServiceServer(server, &Server{})

	// Listen on a network port
	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Start the server
	log.Printf("starting server on port %s", port)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	fmt.Printf("starting server on port %s", port)

	r.Run(":8081")
}
