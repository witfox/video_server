package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"video_server/streamserver/common"
)

type middleWareHandler struct {
	r *httprouter.Router
	l *common.ConnLimiter
}

func NewMiddleWareHandler(r *httprouter.Router, cc int) http.Handler {
	m := middleWareHandler{}
	m.r = r
	m.l = common.NewConnLimiter(cc)
	return  m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	if !m.l.GetConn() {
		common.SendErrorRe(w, http.StatusTooManyRequests, "Too many request")
		return
	}

	m.r.ServeHTTP(w, r)
	defer m.l.ReleaseConn()
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/videos/:video-id", common.StreamHandler)
	router.POST("/upload/:video-id", common.UploadHandler)

	return router
}

func main() {
	r := RegisterHandlers()
	mr := NewMiddleWareHandler(r, 2)
	http.ListenAndServe(":9000", mr)
}
