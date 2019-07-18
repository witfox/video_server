package model

import (
	"testing"
)

func clearTables()  {
	dbConn.Exec("truncate user")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M)  {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T)  {
	t.Run("Add", TestAddUserAuth)
	t.Run("Get", TestGetUserAuth)
	t.Run("Delete", TestDeleteUser)
}

func TestAddUserAuth(t *testing.T) {
	err := AddUserAuth("chenye", "123456")
	if err != nil {
		t.Errorf("add user failed: %v", err)
	}
}
func TestGetUserAuth(t *testing.T) {
	pwd, err := GetUserAuth("chenye")
	if err != nil {
		t.Errorf("get user failed: %v", err)
	}
	if pwd != "123456" {
		t.Errorf("user not exist: %v", err)
	}
}
func TestDeleteUser(t *testing.T) {
	err := DeleteUser("chenye", "123456")
	if err != nil {
		t.Errorf("delete user failed: %v", err)
	}
}