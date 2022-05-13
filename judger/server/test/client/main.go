package main

import (
	judge "OnlineJudge_JudgerServer/pb/submission"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:7101", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	judger := judge.NewJudgeClient(conn)

	_, _ = judger.Judge(context.TODO(), &judge.Submission{
		Id:           "50",
		CodeLanguage: "cpp14",
		CodeText:     `#include <stdio.h>` + "\n" + `int main() {for(;;);}`,
		TestCases: []*judge.TestCase{
			{Input: "2 3\n", Output: "5\n"},
			{Input: "1 1\n", Output: "2\n"},
		},
	})
}
