package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	pb "gRPC-Service/usermgmt"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type UserManagamentServer struct {
	pb.UnimplementedUserManagamentServer
}

func (s *UserManagamentServer) CreateNewUSer(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Received: %v", in.GetName())
	var user_id int = rand.Intn(1000)

	return &pb.User{
		Name: in.GetName(),
		Age:  in.GetAge(),
		Id:   int32(user_id),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserManagamentServer(s, &UserManagamentServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
