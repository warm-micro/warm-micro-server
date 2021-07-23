package main

import (
	"context"
	"fmt"
	"log"
	pb "sprint/sprint"
	"time"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const sprintBatchSize = 3

var sprintMap = make(map[string]pb.Sprint)

type server struct {
	sprintMap map[string]*pb.Sprint
	pb.UnimplementedSprintManagementServer
}

func (s *server) AddSprint(ctx context.Context, sprintReq *pb.Sprint) (*wrappers.StringValue, error) {
	log.Printf("Sprint Added. ID: %v", sprintReq.Id)
	out, err := uuid.NewUUID()
	if err != nil {
		log.Fatal(err)
	}
	sprintReq.Id = out.String()
	sprintMap[sprintReq.Id] = *sprintReq
	return &wrappers.StringValue{Value: "Sprint Added: " + sprintReq.Id}, nil
}

func (s *server) GetSprint(ctx context.Context, sprintId *wrappers.StringValue) (*pb.Sprint, error) {
	ord, exists := sprintMap[sprintId.Value]
	if exists {
		return &ord, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "Sprint does not exists. : ", sprintId)
}

func (s *server) ListSprints(user *wrappers.StringValue, stream pb.SprintManagement_ListSprintsServer) error {
	for _, sprint := range sprintMap {
		if err := stream.Send(&sprint); err != nil {
			return fmt.Errorf("error sending message to stream : %v", err)
		}
	}
	return nil
}

func logUnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("===== [Server Interceptor] ", info.FullMethod)
	startTime := time.Now()
	m, err := handler(ctx, req)
	responseTime := time.Now().Sub(startTime).String()
	log.Println(responseTime)
	if err != nil {
		if e, ok := status.FromError(err); ok {
			logClient.AddLog(ctx, &pb.ApiLog{
				Api:    info.FullMethod,
				Status: e.Code().String(),
				Time:   responseTime,
			})
		} else {
			log.Printf("not able to parse error returned %v", e)
		}
	} else {
		logClient.AddLog(ctx, &pb.ApiLog{
			Api:    info.FullMethod,
			Status: "SUCCESS",
			Time:   responseTime,
		})
	}
	log.Printf("===== [Server Interceptor] : %s", responseTime)
	return m, err
}

func logStreamServerInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Println("===== [Server Interceptor] ", info.FullMethod)
	startTime := time.Now()
	err := handler(srv, stream)
	responseTime := time.Now().Sub(startTime).String()
	if err != nil {
		if e, ok := status.FromError(err); ok {
			logClient.AddLog(stream.Context(), &pb.ApiLog{
				Api:    info.FullMethod,
				Status: e.Code().String(),
				Time:   responseTime,
			})
		} else {
			log.Println("not able to parse error returned %v", e)
		}
	} else {
		logClient.AddLog(stream.Context(), &pb.ApiLog{
			Api:    info.FullMethod,
			Status: "SUCCESS",
			Time:   responseTime,
		})
	}
	log.Printf("===== [Server Interceptor] : %s", time.Now().Sub(startTime))
	return err
}
