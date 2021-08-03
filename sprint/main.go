package main

import (
	"log"
	"net"

	pb "sprint/sprint"

	"google.golang.org/grpc"
)

const (
	port       = ":50052"
	logAddress = "127.0.0.1:50060"
)

var logClient pb.ApiLogMenagementClient

func main() {
	initSampleData()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to liesten: %v", err)
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(logUnaryServerInterceptor),
		grpc.StreamInterceptor(logStreamServerInterceptor),
	)

	conn, err := grpc.Dial(logAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	logClient = pb.NewApiLogMenagementClient(conn)

	pb.RegisterSprintManagementServer(s, &server{})
	log.Printf("Starting gRPC listener on port " + port)
	// reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func initSampleData() {
	sprintMap["101"] = pb.Sprint{Id: "101", WorkspaceId: "1", Description: "test101"}
	sprintMap["102"] = pb.Sprint{Id: "102", WorkspaceId: "1", Description: "test102"}
	sprintMap["103"] = pb.Sprint{Id: "103", WorkspaceId: "1", Description: "test103"}
	sprintMap["104"] = pb.Sprint{Id: "104", WorkspaceId: "2", Description: "test104"}
}
