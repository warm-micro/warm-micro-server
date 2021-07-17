package main

import (
	"context"
	"fmt"
	"log"
	pb "sprint/sprint"

	"github.com/golang/protobuf/ptypes/wrappers"
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
	fmt.Println(user)
	for key, sprint := range sprintMap {
		log.Print(key, sprint)
		if err := stream.Send(&sprint); err != nil {
			return fmt.Errorf("error sending message to stream : %v", err)
		}
	}
	return nil
}
