package controller

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
	"video_server/api/common"
	"video_server/api/defs"
	"video_server/api/model"
	"video_server/api/session"
)

func CreateUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}

	if err := json.Unmarshal(res, ubody); err != nil {
		common.SendErrorRe(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	if err := model.AddUserAuth(ubody.Account, ubody.Password); err != nil  {
		common.SendErrorRe(w, defs.ErrorDb)
		return
	}

	id := session.GenerateNewSessionId(ubody.Account)
	su := &defs.SignedUp{Success:true, SessionId:id}

	if resp, err := json.Marshal(su); err != nil {
		common.SendErrorRe(w, defs.ErrorInternalFaults)
		return
	}else {
		common.SendNormalRe(w, string(resp), 200)
	}
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	username := p.ByName("user_name")
	io.WriteString(w, username)
}