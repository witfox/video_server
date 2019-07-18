package controller

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"video_server/web/common"
	"encoding/json"
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
	if r.Method != http.MethodPost {
		m := common.ErrMsg("Bad Request", "001")
		re, _ := json.Marshal(m)
		io.WriteString(w, string(re))
		return
	}

	res, _ := ioutil.ReadAll(r.Body)
	apibody := &common.ApiBody{}
	if err := json.Unmarshal(res, apibody); err != nil {
		re, _ := json.Marshal(common.ErrMsg("read body error", "002"))
		io.WriteString(w, string(re))
		return
	}

	common.Request(apibody, w, r)

	defer r.Body.Close()

}
