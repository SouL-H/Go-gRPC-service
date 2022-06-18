package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"os"
	"io/ioutil"

	pb "gRPC-Service/usermgmt"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

const (
	port = ":50051"
)

func NewUserManagementServer() *UserManagamentServer {
	return &UserManagamentServer{
		
	}
}

type UserManagamentServer struct {
	pb.UnimplementedUserManagamentServer
	
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
	readBytes, err := ioutil.ReadFile("users.json")
	var users_list *pb.UserList= &pb.UserList{}
	var user_id int = rand.Intn(1000)
	created_user:= &pb.User{
		Name: in.GetName(),
		Age:  in.GetAge(),
		Id:   int32(user_id),
	}
	if err!=nil{
		if os.IsNotExist(err){
			log.Print("File not found")
			users_list.Users=append(users_list.Users,created_user)
			jsonBytes,err := protojson.Marshal(users_list)
			if err!=nil{
				log.Fatalf("Failed to marshal: %v", err)
			}
			if err:= ioutil.WriteFile("users.json",jsonBytes,0664);err!=nil{
				log.Fatalf("Failed to write file: %v", err)

			}
			return created_user,nil
		}else{
			log.Fatalln("Failed to read file: %v", err)
		}
	}
	if err:= protojson.Unmarshal(readBytes,users_list);err!=nil{
		log.Fatalf("Failed to parse user list: %v", err)
	}
	users_list.Users=append(users_list.Users,created_user)
	jsonBytes,err := protojson.Marshal(users_list)
			if err!=nil{
				log.Fatalf("Failed to marshal: %v", err)
			}
			if err:= ioutil.WriteFile("users.json",jsonBytes,0664);err!=nil{
				log.Fatalf("Failed to write file: %v", err)

			}
	return created_user , nil
}
func(s *UserManagamentServer) GetUsers(ctx context.Context,in *pb.GetUsersParams)(*pb.UserList,error){
	jsonBytes,err := ioutil.ReadFile("users.json")
	if err!=nil{
		log.Fatalf("Failed to read file: %v", err)
	}
	var users_list *pb.UserList= &pb.UserList{}
	if err:= protojson.Unmarshal(jsonBytes,users_list);err!=nil{
		log.Fatalf("Unmarshaling failed: %v", err)
	}
	return users_list,nil
}

func main() {
	var user_mgmt_server *UserManagamentServer= NewUserManagementServer()
	if err:=user_mgmt_server.Run();err!=nil{
		log.Fatalf("Failed to run server: %v", err)
	}

}
