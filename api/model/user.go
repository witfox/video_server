package model

import (
	"database/sql"
	"log"
)

func AddUserAuth(account string, pwd string) error {
	model, err := dbConn.Prepare("INSERT INTO user (account, password) VALUES (?,?)")
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	_, err = model.Exec(account, pwd)
	if err != nil {
		return err
	}
	defer model.Close()
	return nil
}

func GetUserAuth(account string) (string, error) {
	model, err := dbConn.Prepare("SELECT password FROM user WHERE account = ?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}
	var pwd string
	err = model.QueryRow(account).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	defer model.Close()
	return pwd, nil
}

func DeleteUser(account string, pwd string) error {
	model, err := dbConn.Prepare("DELETE FROM user where account=? AND password=?")
	if err != nil {
		log.Printf("delete failed: %s", err)
		return err
	}
	_, err = model.Exec(account, pwd)
	if err != nil {
		return err
	}
	defer model.Close()
	return nil
}