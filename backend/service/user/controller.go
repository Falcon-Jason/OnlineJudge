package user

import (
	"OnlineJudge_Backend/util"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	u := util.Request(w, r, new(user))
	if u == nil {
		return
	}

	if u.Id != "" || u.Username == "" || u.Password == "" {
		util.Reply(w, util.ErrInvalidRequest, nil)
		return
	}

	util.Reply(w, register(r.Context(), u), nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	u := util.Request(w, r, new(user))
	if u == nil {
		return
	}

	reply, err := login(r.Context(), u)
	util.Reply(w, err, reply)
}

func ModifyUserInfo(w http.ResponseWriter, r *http.Request) {
	u := util.Request(w, r, new(user))
	if u == nil {
		return
	}

	u.Id = r.Header.Get("user_id")
	if u.Id == "" {
		util.Reply(w, util.ErrAuthorization, nil)
		return
	}

	err := modifyUserInfo(r.Context(), u)
	util.Reply(w, err, nil)
}

func initController() {
	http.HandleFunc("/api/user/register", Register)
	http.HandleFunc("/api/user/login", Login)
	http.HandleFunc("/api/user/modify", AuthLoggedInMiddleware(ModifyUserInfo))
}
