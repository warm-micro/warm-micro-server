package main

import (
	pb "card/card"
	"fmt"

	"context"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const cardBatchSize = 3

var cardMap = make(map[string]pb.Card)

type server struct {
	cardMap map[string]*pb.Card
	pb.UnimplementedCardServerServer
}

func (s *server) AddCard(ctx context.Context, cardReq *pb.Card) (*wrappers.StringValue, error) {
	log.Printf("Card Added. Content: %v", cardReq.SprintId)
	if ok, err := sprintClient.CheckSprint(ctx, &pb.GetSprintRequest{SprintId: cardReq.SprintId}); err != nil || !ok.Value {
		return &wrappers.StringValue{Value: "Wrong Sprint Id: " + cardReq.SprintId}, grpc.Errorf(codes.NotFound, "wrong sprint id")
	}

	out, err := uuid.NewUUID()
	if err != nil {
		log.Fatal(err)
	}
	cardReq.Id = out.String()
	cardMap[cardReq.Id] = *cardReq
	return &wrappers.StringValue{Value: "Card Added: " + cardReq.Id}, nil
}

func (s *server) ListCards(sprintId *pb.SprintId, stream pb.CardServer_ListCardsServer) error {
	for _, card := range cardMap {
		if card.SprintId == sprintId.Value {
			if err := stream.Send(&card); err != nil {
				return fmt.Errorf("error sending message to stream: %v", err)
			}
		}
	}
	return nil
}

func RemoveCard(ctx context.Context, cardId *wrappers.StringValue) (*wrappers.StringValue, error) {
	if _, exists := cardMap[cardId.Value]; exists {
		delete(cardMap, cardId.Value)
		return nil, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "Card does not exists. : %v", cardId.Value)
}

func logUnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("===== [Server Interceptor] ", info.FullMethod)
	startTime := time.Now()
	m, err := handler(ctx, req)
	responseTime := time.Now().Sub(startTime).String()
	log.Println(responseTime)
	if err != nil {
		if e, ok := status.FromError(err); ok {
			log.Println(e.Code())
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
