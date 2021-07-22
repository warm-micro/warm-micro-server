package main

import (
	"context"
	"fmt"
	"log"
	pb "logManager/logManager"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/google/uuid"
)

const logBatchSize = 3

var logMap = make(map[string]pb.ApiLog)
var ApiCount = make(map[string]pb.ApiCount)

type server struct {
	logMap map[string]*pb.ApiLog
	pb.UnimplementedApiLogMenagementServer
}

func (s *server) AddLog(ctx context.Context, logReq *pb.ApiLog) (*wrappers.StringValue, error) {
	log.Printf("Log Added: [%v] %v", logReq.Status, logReq.Api)
	out, err := uuid.NewUUID()
	if err != nil {
		log.Fatal(err)
	}
	logReq.Id = out.String()
	logMap[logReq.Id] = *logReq
	if apiCount, ok := ApiCount[logReq.Api]; ok {
		apiCount.Count += 1
	} else {
		apiCount := pb.ApiCount{Api: logReq.Api, Count: 1}
		ApiCount[logReq.Api] = apiCount
	}
	return &wrappers.StringValue{Value: "Log Added: " + logReq.Id}, nil
}

func (s *server) ListLogs(user *wrappers.StringValue, stream pb.ApiLogMenagement_ListLogsServer) error {
	log.Println(user)
	for key, ApiLog := range logMap {
		log.Print(key, ApiLog)
		if err := stream.Send(&ApiLog); err != nil {
			return fmt.Errorf("error sending message to stream: %v", err)
		}
	}
	return nil
}

func (s *server) ListCounts(user *wrappers.StringValue, stream pb.ApiLogMenagement_ListCountsServer) error {
	log.Println(user)
	for key, apiCount := range ApiCount {
		log.Print(key, apiCount)
		if err := stream.Send(&apiCount); err != nil {
			return fmt.Errorf("eeror sending meessage to stream: %v", err)
		}
	}
	return nil
}
