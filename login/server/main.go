package main

import (
	"context"
	"log"
	"net"

	pb "github.com/lyazii22/grpc-login/login/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type loginServer struct {
	pb.UnimplementedLoginServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLoginServer(s, &loginServer{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *loginServer) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	res := &pb.CreateUserResponse{
		Id:   "id@12323",
		Name: "Tom",
	}

	return res, nil
}

func (s *loginServer) GetUser(context.Context, *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return &pb.GetUserResponse{
		Id:   "idsfs3",
		Name: "Harry",
	}, nil
}

func (s *loginServer) Login(context.Context, *pb.LoginRequest) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{
		Token: "sadfsfasdfsf.asfasdfasfasdfasdfasfasd.asdfasdf",
	}, nil
}
