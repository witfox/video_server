package defs

type Err struct {
	Error string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrorResponse struct {
	HttpSC int
	Error Err
}

var (
	ErrorRequestBodyParseFailed = ErrorResponse{
		HttpSC: 400,
		Error: Err{
			Error: "Request body is not current", ErrorCode: "001",
		},
	}
	ErrorNotAuthUser = ErrorResponse{
		HttpSC: 401,
		Error: Err{
			Error: "No Auth", ErrorCode: "002",
		},
	}
	ErrorDb = ErrorResponse{
		HttpSC: 500,
		Error: Err{
			Error: "Db error", ErrorCode: "003",
		},
	}
	ErrorInternalFaults = ErrorResponse{
		HttpSC: 500,
		Error: Err{
			Error: "Internal error", ErrorCode: "004",
		},
	}
)