package main

import (
	"OnlineJudge_JudgerServer/pb/judge_result"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type writeBackService struct {
	judge_result.UnimplementedResultWriteBackServer
}

func (s *writeBackService) WriteBack(
	ctx context.Context, req *judge_result.ResultWriteBackRequest) (
	*judge_result.ResultWriteBackReply, error) {

	log.Printf("id-%s:%s", req.Id, req.Status.String())

	return &judge_result.ResultWriteBackReply{Ok: true}, nil
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", 7102))
	if err != nil {
		log.Fatalf("Launch Error: failed to listen to %v", listener.Addr())
	}

	server := grpc.NewServer()
	judge_result.RegisterResultWriteBackServer(server, new(writeBackService))
	log.Printf("service listening at %v", listener.Addr())

	err = server.Serve(listener)
	if err != nil {
		log.Fatalf("Launch Error: failed to start service: %v", err)
	}
}
