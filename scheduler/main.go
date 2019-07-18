package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"video_server/scheduler/taskrunner"
)

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()

	router.GET("/del-video-rc/:video-id", videoDelHandler)

	return router
}

func main() {
	go taskrunner.Start()
	r := RegisterHandler()
	http.ListenAndServe(":9001", r)
}
