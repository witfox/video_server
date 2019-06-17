package controller

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
)

type HomePage struct {
	Name string
}
type UserHome struct {
	Name string
}

func HomeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cname, err1 := r.Cookie("account")
	sid, err2 := r.Cookie("session")

	if err1 != nil && err2 != nil {
		p := HomePage{Name: "chenye"}
		t, err := template.ParseFiles("./template/home.html")

		if err != nil {
			log.Printf("Parsing temolate error: %s", err)
			return
		}

		t.Execute(w, p)
		return
	}

	if len(cname.Value) != 0 && len(sid.Value) !=0 {
		http.Redirect(w, r, "userhome", http.StatusFound)
		return
	}
}

func UserHomeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params)  {
	cname, err1 := r.Cookie("account")
	_, err2 := r.Cookie("session")

	if err1 != nil && err2!= nil{
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	fname := r.FormValue("account")

	var p *UserHome
	if len(cname.Value) !=0 {
		p = &UserHome{Name: cname.Value}
	}else if len(fname) != 0 {
		p = &UserHome{Name: fname}
	}

	t, err := template.ParseFiles("./template/userhome.html")

	if err != nil {
		return
	}

	t.Execute(w, p)
	return
}

func ApiHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params)  {
	return
}
