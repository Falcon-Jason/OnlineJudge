package problem

import (
	"OnlineJudge_Backend/util"
	"log"
	"net/http"
)

func AddProblem(w http.ResponseWriter, r *http.Request) {
	p := util.Request(w, r, &problem{})
	if p == nil {
		return
	}

	log.Print(p.No)
	p.AuthorId = r.Header.Get("user_id")

	if p.Title == "" || p.No == "" {
		util.Reply(w, util.ErrInvalidRequest, nil)
		return
	}

	util.Reply(w, addProblem(r.Context(), p), nil)
}

func checkAuthorization(w http.ResponseWriter, r *http.Request, pid string) bool {
	if pid == "" {
		util.Reply(w, util.ErrInvalidRequest, nil)
		return false
	}

	req := QueryProblemRequest{pid}
	userId := r.Header.Get("user_id")
	if owner, err := getProblemOwner(r.Context(), req); err != nil || owner != userId {
		util.Reply(w, util.ErrAuthorization, nil)
		return false
	}

	return true
}

func DeleteProblem(w http.ResponseWriter, r *http.Request) {
	req := util.Request(w, r, new(deleteProblemRequest))
	if req == nil {
		return
	}

	if !checkAuthorization(w, r, req.ProblemId) {
		return
	}

	util.Reply(w, deleteProblem(r.Context(), *req), nil)
}

func ModifyProblem(w http.ResponseWriter, r *http.Request) {
	p := util.Request(w, r, new(problem))
	if p == nil {
		return
	}

	if !checkAuthorization(w, r, p.Id) {
		return
	}

	util.Reply(w, modifyProblem(r.Context(), p), nil)
}

func ListProblem(w http.ResponseWriter, r *http.Request) {
	req := util.Request(w, r, new(listProblemRequest))
	if req == nil {
		return
	}

	ret, err := listProblem(r.Context(), *req)
	util.Reply(w, err, ret)
}

func SearchProblem(w http.ResponseWriter, r *http.Request) {
	req := util.Request(w, r, new(searchProblemRequest))
	if req == nil {
		return
	}

	ret, err := searchProblem(r.Context(), *req)
	util.Reply(w, err, ret)
}

func QueryProblemDetail(w http.ResponseWriter, r *http.Request) {
	req := util.Request(w, r, new(QueryProblemRequest))
	if req == nil {
		return
	}

	ret, err := queryProblemDetail(r.Context(), *req)
	util.Reply(w, err, ret)
}

func QueryProblem(w http.ResponseWriter, r *http.Request) {
	req := util.Request(w, r, new(QueryProblemRequest))
	if req == nil {
		return
	}

	ret, err := queryProblem(r.Context(), *req)
	util.Reply(w, err, ret)
}
