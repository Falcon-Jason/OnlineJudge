package util

import (
	"encoding/json"
	"net/http"
)

func Request[T interface{}](w http.ResponseWriter, r *http.Request, t *T) *T {
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		Reply(w, err, nil)
		return nil
	}

	return t
}
