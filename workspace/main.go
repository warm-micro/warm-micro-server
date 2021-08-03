package main

import (
	"log"
	"net"
	pb "workspace/workspace"

	"google.golang.org/grpc"
)

const (
	port       = ":50054"
	logAddress = "127.0.0.1:50060"
)

var logClient pb.ApiLogMenagementClient

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

	pb.RegisterWorkspaceManagerServer(s, &server{})
	log.Println("Starting gRPC listener on port " + port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func initSampleData() {
	workspaceMap["1001"] = pb.Workspace{Id: "1001", Name: "WM"}
	workspaceMap["1002"] = pb.Workspace{Id: "1002", Name: "PNU"}
}
