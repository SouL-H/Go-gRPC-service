package main

import (
	pb "gRPC-Service/usermgmt"

	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewUserManagamentClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var new_users = make(map[string]int32)
	new_users["user1"] = 1
	new_users["user2"] = 2
	new_users["user3"] = 3
	new_users["user4"] = 4

	for name, age := range new_users {
		r, err := c.CreateNewUSer(ctx, &pb.NewUser{Name: name, Age: age})
		if err != nil {
			log.Fatalf("could not create user: %v", err)
		}
		log.Printf(`
		User Details: 
		Name: %s
		Age: %d
		ID: %d`, r.GetName(), r.GetAge(), r.GetId())
	}
}
