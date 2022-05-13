package problem

import (
	"OnlineJudge_Backend/dao"
	"OnlineJudge_Backend/service/user"
	"net/http"
)

func Init() {
	problems = dao.Mongo.Collection("problem")

	http.HandleFunc("/api/problem/add", user.AuthLoggedInMiddleware(user.AuthTeacherMiddleware(AddProblem)))
	http.HandleFunc("/api/problem/delete", user.AuthLoggedInMiddleware(user.AuthTeacherMiddleware(DeleteProblem)))
	http.HandleFunc("/api/problem/modify", user.AuthLoggedInMiddleware(user.AuthTeacherMiddleware(ModifyProblem)))
	http.HandleFunc("/api/problem/list", ListProblem)
	http.HandleFunc("/api/problem/search", SearchProblem)
	http.HandleFunc("/api/problem/query_detail", QueryProblemDetail)
	http.HandleFunc("/api/problem/query", user.AuthLoggedInMiddleware(user.AuthTeacherMiddleware(QueryProblem)))
}
