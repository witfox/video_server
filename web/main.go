package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"video_server/web/controller"
)

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", controller.HomeHandler)
	router.POST("/", controller.HomeHandler)
	router.GET("/userhome", controller.UserHomeHandler)
	router.POST("/userhome", controller.UserHomeHandler)
	router.POST("/api", controller.ApiHandler)

	router.ServeFiles("/statics/*filepath", http.Dir("./template"))

	return router
}

func main() {
	r := RegisterHandler()
	http.ListenAndServe(":9002", r)
}
