package models

func AddVideoDeleteRe(vid string) error {
	result, err := dbConn.Prepare("INSERT INTO video_del_sc(video_id) VALUE (?)")
	if err != nil {
		return err
	}

	_, err = result.Exec(vid)
	if err != nil {
		return err
	}

	defer result.Close()
	return nil
}
