package defs

type UserCredential struct {
	Account string `json:"account"`
	Password string `json:"password"`
}

type SignedUp struct {
	Success bool `json:"success"`
	SessionId string `json:"session_id"`
} 

type VideoInfo struct {
	Id string
	AuthorId int
	Name string
	DisplayCtime string
}


type Comment struct {
	Id string
	VideoId string
	Author string
	Content string
}

type SimpleSession struct {
	Account string
	TTL int64
}