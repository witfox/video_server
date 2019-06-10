package common

import (
	"net/http"
	"video_server/api/defs"
	"video_server/api/session"
)

var HEADER_FIELD_SESSION = "X-Session-Id"
var HEADER_FIELD_ACCOUNT = "X-Session-Account"

func ValidateUserSession(r *http.Request) bool {
	sid := r.Header.Get(HEADER_FIELD_SESSION)
	if len(sid) == 0 {
		return false
	}

	account, ok := session.IsSessionExpired(sid)
	if ok {
		return false
	}

	r.Header.Set(HEADER_FIELD_ACCOUNT, account)
	return true
}

func ValidateUser(w http.ResponseWriter, r *http.Request) bool {
	
	account := r.Header.Get(HEADER_FIELD_ACCOUNT)
	if len(account) == 0 {
		SendErrorRe(w, defs.ErrorNotAuthUser)
		return false
	}

	return true
}


