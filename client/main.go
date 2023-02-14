package main

import (
	"context"
	"gRPC/pb"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewUserClient(conn)

	reqInsert := &pb.UserRequest{
		Name: "Insert",
	}

	resIn, err := client.Insert(context.Background(), reqInsert)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(resIn)

	reqUpdate := &pb.UserRequest{
		Name: "Updade",
	}

	resUp, err := client.Update(context.Background(), reqUpdate)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(resUp)

	reqDelete := &pb.UserRequest{
		Name: "Delete",
	}

	resDelete, err := client.Delete(context.Background(), reqDelete)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(resDelete)

	reqFind := &pb.UserRequest{
		Name: "Find",
	}

	resFind, err := client.Find(context.Background(), reqFind)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(resFind)

}
