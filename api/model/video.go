package model

import (
	"database/sql"
	"log"
	"time"
	"video_server/api/defs"
	"video_server/api/utils"
)

func AddVideo(aid int, name string) (*defs.VideoInfo, error) {
	vid, _ := utils.NewUUID()
	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05")

	model, err := dbConn.Prepare("INSERT INTO video_info (id, author_id, name, display_ctime) VALUES (?,?,?,?)")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	_, err = model.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err
	}
	defer model.Close()
	return &defs.VideoInfo{
		Id: vid,
		AuthorId: aid,
		Name: name,
		DisplayCtime: ctime,
	}, err
}

func GetVideo(vid string) (*defs.VideoInfo, error) {
	model, err := dbConn.Prepare("SELECT * FROM video_info WHERE id = ?")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	var videoInfo *defs.VideoInfo
	err = model.QueryRow(vid).Scan(&videoInfo)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	defer model.Close()
	return videoInfo, nil
}

func DeleteVideo(vid string) error {
	model, err := dbConn.Prepare("DELETE FROM video_info where id=?")
	if err != nil {
		log.Printf("delete failed: %s", err)
		return err
	}
	_, err = model.Exec(vid)
	if err != nil {
		return err
	}
	defer model.Close()
	return nil
}