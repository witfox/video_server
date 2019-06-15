package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", homeHander)
	router.POST("/", homeHander)
	router.GET("/userhome", userHomeHander)
	router.POST("/userhome", userHomeHandler)
	router.POST("/api", apiHandler)

	router.ServeFiles("/statics/*filepath", http.Dir("./template"))

	return router
}

func main() {
	r := RegisterHandler()
	http.ListenAndServe(":9002", r)
}
