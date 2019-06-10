package session

import (
	"sync"
	"time"
	"video_server/api/defs"
	"video_server/api/model"
	"video_server/api/utils"
)

var sessionMap *sync.Map
func init()  {
	sessionMap = &sync.Map{}
}

func LoadSessionsFromDB(){
	r, err := model.RetrieveAllSessions()
	if err != nil {
		return
	}
	r.Range(func(k, v interface{}) bool {
		ss := v.(*sync.Map)
		sessionMap.Store(k, ss)
		return true
	})

}

func deleteExpiredSession(sid string)  {
	sessionMap.Delete(sid)
	model.DeleteSession(sid)
}

func GenerateNewSessionId(uname string) string {
	id, _ := utils.NewUUID()
	ct := time.Now().UnixNano()/1000000
	ttl := ct + 30 * 60 * 1000

	ss := &defs.SimpleSession{Account:uname, TTL:ttl}
	sessionMap.Store(id, ss)
	model.InsertSession(id, ttl, uname)

	return id
}

func IsSessionExpired(sid string) (string, bool) {
	if ss, ok := sessionMap.Load(sid); ok {
		ct := time.Now().UnixNano()/1000000
		if ss.(*defs.SimpleSession).TTL < ct {
			deleteExpiredSession(sid)
			return "", true
		}
		return ss.(*defs.SimpleSession).Account, false
	}
	deleteExpiredSession(sid)
	return "", true
}