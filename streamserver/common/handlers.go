package common

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func StreamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	vid := p.ByName("video-id")
	vl := VIDEO_DIR + vid

	video, err := os.Open(vl)
	if err != nil {
		log.Printf("open file error: %v", err)
		SendErrorRe(w, http.StatusInternalServerError, "Internal error")
		return
	}
	w.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(w, r, "", time.Now(), video)

	defer video.Close()
}

func UploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	r.Body = http.MaxBytesReader(w, r.Body,MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE);err != nil {
		SendErrorRe(w, http.StatusBadRequest, "File is too larger")
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		SendErrorRe(w, http.StatusInternalServerError, "Internal error")
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil{
		log.Printf("read file error: %v", err)
		SendErrorRe(w, http.StatusInternalServerError, "Internal error")
		return
	}

	fn := p.ByName("video-id")
	err = ioutil.WriteFile(VIDEO_DIR + fn, data, 0666)
	if err != nil {
		log.Printf("write file error: %v", err)
		SendErrorRe(w, http.StatusInternalServerError, "Internal error")
		return
	}

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Uploaded successfully")
}