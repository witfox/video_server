package common

import (
	"encoding/json"
	"io"
	"net/http"
	"video_server/api/defs"
)

func SendErrorRe(w http.ResponseWriter, errResp defs.ErrorResponse)  {
	w.WriteHeader(errResp.HttpSC)
	resStr,_ := json.Marshal(&errResp.Error)
	io.WriteString(w, string(resStr))
}

func SendNormalRe(w http.ResponseWriter, resp string, sc int)  {
	w.WriteHeader(sc)
	io.WriteString(w, resp)

}
