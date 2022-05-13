package main

import (
	"OnlineJudge_JudgerServer/pb/judge_result"
	"OnlineJudge_JudgerServer/pb/submission"
	"OnlineJudge_JudgerServer/service"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
)

var (
	port   = flag.Int("port", 7101, "The port listened by this service")
	wbAddr = flag.String("wb_addr", "localhost:7102", "The port of write back service")
)

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Launch Error: failed to listen to %v", listener.Addr())
	}

	conn, err := grpc.Dial(*wbAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Launch Error: failed to connect to write back service %s", *wbAddr)
	}

	defer conn.Close()

	server := grpc.NewServer()
	submission.RegisterJudgeServer(server, service.NewJudgeService(
		judge_result.NewResultWriteBackClient(conn)))

	log.Printf("service listening at %v", listener.Addr())

	err = server.Serve(listener)
	if err != nil {
		log.Fatalf("Launch Error: failed to start service: %v", err)
	}
}
