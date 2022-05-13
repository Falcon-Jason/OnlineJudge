package util

import (
	"encoding/json"
	"io"
)

type reply struct {
	OK      bool        `json:"ok"`
	ErrInfo string      `json:"err_info,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}

func Reply(writer io.Writer, err error, result interface{}) {
	r := reply{
		OK:      err == nil,
		ErrInfo: "",
		Result:  result,
	}

	if err != nil {
		r.ErrInfo = err.Error()
	}

	_ = json.NewEncoder(writer).Encode(&r)
}
