package main

import (
	"context"
	"fmt"
	"log"
	"time"
	pb "workspace/workspace"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

const WorkspaceBatchSize = 3

var workspaceMap = make(map[string]pb.Workspace)

type server struct {
	workspaceMap map[string]*pb.Workspace
	pb.UnimplementedWorkspaceManagerServer
}

func (s *server) AddWorkspace(ctx context.Context, workspaceReq *pb.Workspace) (*wrappers.StringValue, error) {
	out, err := uuid.NewUUID()
	if err != nil {
		log.Fatal(err)
	}
	workspaceReq.Id = out.String()
	workspaceMap[workspaceReq.Id] = *workspaceReq
	return &wrappers.StringValue{Value: "Workspace Added: " + workspaceReq.Id}, nil
}

func (s *server) ListWorkspace(user *wrappers.StringValue, stream pb.WorkspaceManager_ListWorkspaceServer) error {
	for _, workspace := range workspaceMap {
		if err := stream.Send(&workspace); err != nil {
			return fmt.Errorf("error sending message to stream: %v", err)
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
