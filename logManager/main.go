package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"logManager/db"
	pb "logManager/logManager"
)

const (
	port = ":50060"
)

func init() {
	db.Connect()
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterApiLogMenagementServer(s, &server{})
	log.Printf("Starting gRPC listener on port " + port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
