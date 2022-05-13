package submission

import (
	ss "OnlineJudge_Backend/pb/submission"
	"OnlineJudge_Backend/service/problem"
	"OnlineJudge_Backend/service/user"
	"OnlineJudge_Backend/util"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

var judgerConn *grpc.ClientConn
var judger ss.JudgeClient

func Submit(w http.ResponseWriter, r *http.Request) {
	s := util.Request(w, r, new(submission))
	if s == nil {
		return
	}

	if s.Id != "" {
		util.Reply(w, util.ErrInvalidRequest, nil)
		return
	}

	s.AuthorId = r.Header.Get("user_id")
	s.Status = "pending"

	if err := submit(r.Context(), s); err != nil {
		util.Reply(w, err, nil)
		return
	}

	id, err := primitive.ObjectIDFromHex(s.ProblemId)
	if err != nil {
		util.Reply(w, err, nil)
		return
	}

	tc, err := problem.QueryTestCase(r.Context(), id)
	if err != nil {
		util.Reply(w, err, nil)
		log.Print(id)
		return
	}

	fs := ss.Submission{
		Id:           s.Id,
		CodeLanguage: s.CodeLanguage,
		CodeText:     s.CodeText,
		TestCases:    []*ss.TestCase{},
	}

	for _, i := range tc.Cases {
		fs.TestCases = append(
			fs.TestCases,
			&ss.TestCase{
				Input:  i.Input,
				Output: i.Output,
			})
	}

	_, _ = judger.Judge(context.TODO(), &fs)

	//fs := fullSubmission{submission: s, TestCases: tc}
	//m, err := json.Marshal(fs)
	//log.Printf("submit: %v", string(m))
	util.Reply(w, err, nil)
}

func SearchSubmissions(w http.ResponseWriter, r *http.Request) {
	req := util.Request(w, r, new(searchSubmissionRequest))
	if req == nil {
		return
	}

	idStr := r.Header.Get("user_id")
	id, err := primitive.ObjectIDFromHex(idStr)

	if err != nil {
		util.Reply(w, err, nil)
		return
	}

	auth, err := user.Authorization(r.Context(), &user.AuthRequest{id})
	if !auth.IsTeacher && !auth.IsAdmin {
		req.AuthorId = idStr
	}

	data, err := searchSubmissions(r.Context(), req)
	util.Reply(w, err, data)
}

func QuerySubmission(w http.ResponseWriter, r *http.Request) {
	req := util.Request(w, r, new(querySubmissionRequest))
	if req == nil {
		return
	}

	data, err := querySubmission(r.Context(), req)
	if data.AuthorId != r.Header.Get("user_id") {
		util.Reply(w, util.ErrAuthorization, nil)
		return
	}

	util.Reply(w, err, data)
}
