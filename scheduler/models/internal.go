package models

func ReadDeleteVideoR(count int) ([]string, error) {
	var ids []string
	result, err := dbConn.Prepare(`SELECT video_id FROM video_del_sc limit ?`)
	if err != nil {
		return ids, err
	}

	rows, err := result.Query(count)
	if err != nil{
		return ids, err
	}

	for rows.Next() {
		var id string
		err = rows.Scan(&id)
		if err != nil {
			return ids, err
		}
		ids = append(ids, id)
	}

	defer result.Close()

	return ids, nil
}

func DelDeleteVideoR(vid string) error {
	result, err := dbConn.Prepare("DELETE FROM video_del_sc WHERE video_id=?")
	defer result.Close()
	if err != nil {
		return err
	}

	_, err = result.Exec(vid)
	if err != nil {
		return err
	}

	return nil
}
