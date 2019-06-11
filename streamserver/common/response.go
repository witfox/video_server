package common

import (
	"io"
	"net/http"
)

func SendErrorRe(w http.ResponseWriter, sc int, msg string)  {
	w.WriteHeader(sc)
	io.WriteString(w, msg)
}
