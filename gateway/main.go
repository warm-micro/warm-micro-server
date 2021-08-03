package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	pb "gateway/gateway"
)

const (
	port          = ":50051"
	sprintPort    = ":50052"
	cardPort      = ":50053"
	workspaceProt = ":50054"
	logPort       = ":50060"
)

type server struct {
	pb.UnimplementedGateWayServer
}

func (s *server) Echo(ctx context.Context, in *pb.StringRequest) (*pb.StringResponse, error) {
	return &pb.StringResponse{Value: in.Value}, nil
}

func main() {
	gwmux := runtime.NewServeMux()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln("failed to listen: ", err)
	}

	s := grpc.NewServer()
	pb.RegisterGateWayServer(s, &server{})
	log.Println("Serving gRPC on 127.0.0.1" + port)
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	test_conn, err := grpc.DialContext(
		context.Background(),
		"127.0.0.1"+port,
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)

	if err != nil {
		log.Fatalln("failed to dial server: ", err)
	}

	err = pb.RegisterGateWayHandler(context.Background(), gwmux, test_conn)
	if err != nil {
		log.Fatal("failed to register gateway: ", err)
	}

	sprintConn, err := grpc.DialContext(
		context.Background(),
		"127.0.0.1"+sprintPort,
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("failed to dial server: ", err)
	}
	if err = pb.RegisterSprintManagementHandler(context.Background(), gwmux, sprintConn); err != nil {
		log.Fatal("failed to register gateway: ", err)
	}

	cardConn, err := grpc.DialContext(
		context.Background(),
		"127.0.0.1"+cardPort,
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("failed to dial server: ", err)
	}
	if err = pb.RegisterCardServerHandler(context.Background(), gwmux, cardConn); err != nil {
		log.Fatal("failed to register gateway: ", err)
	}

	workspaceConn, err := grpc.DialContext(
		context.Background(),
		"127.0.0.1"+workspaceProt,
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("failed to dial server: ", err)
	}
	if err = pb.RegisterWorkspaceManagerHandler(context.Background(), gwmux, workspaceConn); err != nil {
		log.Fatal("failed to register gateway: ", err)
	}

	logConn, err := grpc.DialContext(
		context.Background(),
		"127.0.0.1"+logPort,
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("failed to dial server: ", err)
	}
	if err = pb.RegisterApiLogMenagementHandler(context.Background(), gwmux, logConn); err != nil {
		log.Fatal("failed to register gateway: ", err)
	}

	gwServer := &http.Server{
		Addr:    ":8080",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://127.0.0.1:8080")
	log.Fatalln(gwServer.ListenAndServe())
}
