package main

import (
	pb "card/card"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port          = ":50053"
	logAddress    = "localhost:50060"
	sprintAddress = "localhost:50052"
)

var logClient pb.ApiLogMenagementClient
var sprintClient pb.SprintManagementClient

func init() {
	initSampleData()
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(logUnaryServerInterceptor),
		grpc.StreamInterceptor(logStreamServerInterceptor),
	)
	logConn, err := grpc.Dial(logAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer logConn.Close()
	logClient = pb.NewApiLogMenagementClient(logConn)

	sprintConn, err := grpc.Dial(sprintAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer sprintConn.Close()
	sprintClient = pb.NewSprintManagementClient(sprintConn)
	pb.RegisterCardServerServer(s, &server{})
	log.Println("Starting gRPC listenr on port " + port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func initSampleData() {
	cardMap["101"] = pb.Card{Id: "101", SprintId: "101", Name: "i wanna zol-up", Content: "help me", UserId: []string{"1", "2"}}
	cardMap["102"] = pb.Card{Id: "102", SprintId: "101", Content: "hello grpc", UserId: []string{"1"}}
	cardMap["103"] = pb.Card{Id: "103", SprintId: "102", Content: "hungry"}
}
