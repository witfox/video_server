package model

import (
	"database/sql"
	"log"
	"strconv"
	"sync"
	"video_server/api/defs"
)

func InsertSession(sid string, ttl int64, uname string) error {
	ttlstr := strconv.FormatInt(ttl, 10)

	model, err := dbConn.Prepare("INSERT INTO sessions (session_id, account, TTL) values (?,?,?)")
	if err != nil {
		return err
	}
	_, err = model.Exec(sid, uname, ttlstr)
	if err != nil {
		return err
	}

	defer model.Close()

	return nil
}

func RetrieveSession(sid string) (*defs.SimpleSession, error) {

	model, err := dbConn.Prepare("SELECT account, TTL from sessions WHERE session_id = ?")
	if err != nil {
		return nil, err
	}
	var ttl string
	var account string
	err = model.QueryRow(sid).Scan(&account, &ttl)

	if err != nil && err != sql.ErrNoRows{
		return nil, err
	}

	defer model.Close()

	var ttlStr, _ = strconv.ParseInt(ttl, 10, 64)
	return &defs.SimpleSession{Account: account, TTL:ttlStr}, nil

}

func RetrieveAllSessions() (*sync.Map, error) {
	m := &sync.Map{}
	model, err := dbConn.Prepare("SELECT * from sessions")
	if err != nil {
		return nil, err
	}
	rows, err := model.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id string
		var ttlstr string
		var account string
		if err := rows.Scan(&id, &account, &ttlstr);err != nil {
			log.Printf("retrive sessions error: %s", err)
			break
		}
		if ttl, err1 := strconv.ParseInt(ttlstr, 10, 64); err1 == nil {
			ss := &defs.SimpleSession{Account: account, TTL:ttl}
			m.Store(id, ss)
			log.Printf(" session id: %s, ttl: %d", id, ss.TTL)
		}
	}
	return m, nil
}

func DeleteSession(sid string) error {
	model, err := dbConn.Prepare("DELETE FROM sessions where session_id=?")
	if err != nil {
		log.Printf("delete failed: %s", err)
		return err
	}
	_, err = model.Exec(sid)
	if err != nil {
		return err
	}
	defer model.Close()
	return nil
}