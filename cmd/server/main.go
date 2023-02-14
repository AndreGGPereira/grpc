package main

import (
	"context"
	"gRPC/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedUserServer
}

func (s *Server) Insert(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{
		Name: "Insert new User",
	}, nil
}

func (s *Server) Update(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{
		Name: "Updade User",
	}, nil
}

func (s *Server) Delete(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{
		Name: "Delete  User",
	}, nil
}

func (s *Server) Find(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{
		Name: "Find users",
	}, nil
}

func main() {
	println("Running gRPC server")

	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServer(s, &Server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to server %v", err)
	}

}
