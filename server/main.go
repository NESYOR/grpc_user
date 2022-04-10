// Package main implements a server for Greeter service.
package main

import (
	"context"
	"log"
	"net"

	pb "github.com/gouthamreddy09/totality-corp/assignment/userrpc"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type user struct {
	id      int64
	fname   string
	city    string
	phone   int64
	height  float32
	Married bool
}

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGetUserServer
}

var users = []user{
	user{1, "Hervey", "Bayanbaraat", 4355014679, 5.9, true},
	user{2, "Jonathan", "Mýto", 6959756970, 6.8, true},
	user{3, "Andros", "Leidian", 7182871166, 5.8, true},
	user{4, "Aleda", "Pittsburgh", 4122686890, 4.9, true},
	user{5, "Viole", "Boudinar", 6539944888, 4.9, true},
	user{6, "Alicia", "Nanshe", 3984985555, 5.0, false},
	user{7, "Bethany", "Panyam", 1719996540, 6.6, false},
	user{8, "Dyanne", "Neos Voutzás", 8629979173, 5.1, false},
	user{9, "Karia", "Ryazhsk", 8819883852, 6.1, true},
	user{10, "Eward", "Oulmes", 2375006597, 4.4, true},
}

// SayHello implements helloworld.GreeterServer
func (s *server) GetUser(ctx context.Context, in *pb.FetchUserByIdRequest) (*pb.User, error) {
	log.Printf("Received: %v", in.GetId())

	tmpuser := user{}
	n := in.Id

	for i := range users {
		if users[i].id == n {
			tmpuser = users[i]
		}
	}

	return &pb.User{
		Id:      tmpuser.id,
		Fname:   tmpuser.fname,
		City:    tmpuser.city,
		Phone:   tmpuser.phone,
		Height:  tmpuser.height,
		Married: tmpuser.Married}, nil //HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server) GetUserList(ctx context.Context, in *pb.FetchUserListByIdRequest) (*pb.UsersList, error) {
	log.Printf("Received: %v", in.GetId())

	tmpUsers := []*pb.User{}

	for i := range in.Id {
		j := 0

		for j < len(users) {
			if users[j].id == in.Id[i] {
				tmpUsers = append(tmpUsers, &pb.User{
					Id:      users[j].id,
					Fname:   users[j].fname,
					City:    users[j].city,
					Phone:   users[j].phone,
					Height:  users[j].height,
					Married: users[j].Married,
				})
			}
			j = j + 1

		}

	}

	return &pb.UsersList{Users: tmpUsers}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGetUserServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
