package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type HomePage struct {
	Name string
}

func HomeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	p
}
