package user

import (
	"OnlineJudge_Backend/util"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

//type AuthLoggedInMiddleware struct {
//	Next func(w http.ResponseWriter, r *http.Request)
//}

type MiddlewareFunc func(w http.ResponseWriter, r *http.Request)

func AuthLoggedInMiddleware(next MiddlewareFunc) MiddlewareFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if next == nil {
			next = http.DefaultServeMux.ServeHTTP
		}

		token := r.Header.Get("token")
		if token == "" {
			util.Reply(w, errors.New("not yet login"), nil)
			return
		}

		id, err := queryUserIdByToken(r.Context(), token)
		if err != nil {
			util.Reply(w, err, nil)
			return
		}

		r.Header.Set("user_id", id)
		next(w, r)
	}
}

func AuthTeacherMiddleware(next MiddlewareFunc) MiddlewareFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if next == nil {
			next = http.DefaultServeMux.ServeHTTP
		}

		id, err := primitive.ObjectIDFromHex(r.Header.Get("user_id"))
		if err != nil {
			util.Reply(w, err, nil)
			return
		}

		reply, err := Authorization(r.Context(), &AuthRequest{Id: id})

		if err != nil {
			util.Reply(w, err, nil)
			return
		}

		if !reply.IsTeacher && !reply.IsAdmin {
			util.Reply(w, util.ErrAuthorization, nil)
			return
		}

		next(w, r)
	}
}
