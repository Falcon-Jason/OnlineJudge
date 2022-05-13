package submission

import (
	"OnlineJudge_Backend/dao"
	s "OnlineJudge_Backend/pb/submission"
	"OnlineJudge_Backend/service/user"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

var judgerAddr = flag.String("judger_addr", "localhost:7101", "the default judger address")

func Init() {
	submissions = dao.Mongo.Collection("submission")

	conn, err := grpc.Dial(*judgerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	judgerConn = conn
	judger = s.NewJudgeClient(conn)

	http.HandleFunc("/api/submission/submit", user.AuthLoggedInMiddleware(Submit))
	http.HandleFunc("/api/submission/search_submission", user.AuthLoggedInMiddleware(SearchSubmissions))
	http.HandleFunc("/api/submission/query_submission", user.AuthLoggedInMiddleware(QuerySubmission))
}

func Close() {
	judgerConn.Close()
}
