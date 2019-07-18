package model

import (
	"log"
	"video_server/api/defs"
	"video_server/api/utils"
)

func AddComment(aid int, content string, videoId string) (*defs.Comment, error) {
	cid, _ := utils.NewUUID()

	model, err := dbConn.Prepare("INSERT INTO comments (id, video_id, author_id, content) VALUES (?,?,?,?)")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	_, err = model.Exec(cid, videoId, aid, content)
	if err != nil {
		return nil, err
	}
	defer model.Close()
	return &defs.Comment{
		Id: cid,
		VideoId: videoId,
		Author: "",
		Content: content,
	}, err
}

func DeleteComment(cid string) error {
	model, err := dbConn.Prepare("DELETE FROM comments where id=?")
	if err != nil {
		log.Printf("delete failed: %s", err)
		return err
	}
	_, err = model.Exec(cid)
	if err != nil {
		return err
	}
	defer model.Close()
	return nil
}

func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
	sql := "SELECT c.id,u.account,c.content INNER JOIN users ON c.author_id = u.id WHERE c.video_id = ? and c.create_at > FROM_UNIXTIME(?) AND c.create_at <= FROM_UNIXTIME(?)"
	model, err := dbConn.Prepare(sql)

	var res []*defs.Comment

	rows, err := model.Query()
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var id, account, content string
		if err := rows.Scan(&id, &account, &content); err != nil {
			return res, err
		}

		c := &defs.Comment{Id: id, VideoId: vid, Author: account, Content: content}
		res = append(res, c)
	}

	return res, nil

}