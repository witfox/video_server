package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"video_server/api/common"
	"video_server/api/controller"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	common.ValidateUserSession(r)
	m.r.ServeHTTP(w, r)
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	//用户路由
	router.POST("/user", controller.CreateUsers)
	router.POST("/user/:user_name", controller.Login)

	return router
}

func main()  {
	r := RegisterHandlers()
	mr := NewMiddleWareHandler(r)
	http.ListenAndServe(":8000", mr)
}