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
	port        = ":50051"
	sprint_port = ":50052"
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
	log.Println("Serving gRPC on 0.0.0.0" + port)
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	test_conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0"+port,
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

	sprint_conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0"+sprint_port,
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)

	if err != nil {
		log.Fatalln("failed to dial server: ", err)
	}

	err = pb.RegisterSprintManagementHandler(context.Background(), gwmux, sprint_conn)

	if err != nil {
		log.Fatal("failed to register gateway: ", err)
	}

	gwServer := &http.Server{
		Addr:    ":8080",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8080")
	log.Fatalln(gwServer.ListenAndServe())
}