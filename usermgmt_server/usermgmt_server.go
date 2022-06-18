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

func NewUserManagementServer() *UserManagamentServer {
	return &UserManagamentServer{
		user_list: &pb.UserList{},
	}
}

type UserManagamentServer struct {
	pb.UnimplementedUserManagamentServer
	user_list *pb.UserList
}

func (server *UserManagamentServer) Run() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserManagamentServer(s, server)
	log.Printf("server listening at %v", lis.Addr())
	return s.Serve(lis)
}
func (s *UserManagamentServer) CreateNewUSer(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Received: %v", in.GetName())
	var user_id int = rand.Intn(1000)
	created_user:= &pb.User{
		Name: in.GetName(),
		Age:  in.GetAge(),
		Id:   int32(user_id),
	}
	s.user_list.Users=append(s.user_list.Users, created_user)
	return created_user , nil
}
func(s *UserManagamentServer) GetUsers(ctx context.Context,in *pb.GetUsersParams)(*pb.UserList,error){
	return s.user_list,nil
}

func main() {
	var user_mgmt_server *UserManagamentServer= NewUserManagementServer()
	if err:=user_mgmt_server.Run();err!=nil{
		log.Fatalf("Failed to run server: %v", err)
	}

}
