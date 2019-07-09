package common

import "video_server/api/defs"

func ErrMsg(msg string, code string) defs.Err {
	return defs.Err{Error: msg, ErrorCode: code}
}
