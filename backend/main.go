package main

import (
	"OnlineJudge_Backend/dao"
	"OnlineJudge_Backend/service/problem"
	"OnlineJudge_Backend/service/submission"
	"OnlineJudge_Backend/service/user"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	port = flag.Int("port", 8080, "The port to be listened")
)

func main() {
	flag.Parse()

	dao.Init()
	defer dao.Close()

	user.Init()
	problem.Init()
	submission.Init()
	defer submission.Close()

	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", *port), nil)
	if err != nil {
		log.Fatalf("failed to listen and serve http service: %v", err)
	}
}
