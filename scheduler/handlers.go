package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"video_server/api/common"
	"video_server/scheduler/models"
)

func videoDelHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("video-id")
	if len(vid) == 0 {
		common.SendNormalRe(w, "参数必须", 400)
		return
	}

	err := models.AddVideoDeleteRe(vid)
	if err != nil {
		common.SendNormalRe(w, "internal error", 500)
		return
	}

	common.SendNormalRe(w, "", 200)
	return
}
