package errmsg

// 处理前端之后如果有错误，返回给前端的错误码
const (
	SUCCESS = 202
	ERR     = 505

	// 1000为user的错误码
	ERROR_NAME_USED = 1000 + iota - 1
	ERROR_PASSWORD_WRONG
	ERROR_USER_NOT_EXIST
	ERROR_TOKEN_EXIST
	ERROR_TOKEN_RUNTIME
	ERROR_TOKEN_WRONG
	ERROR_TOKEN_TYPE
	USER_ROLE_COM
)

var codeMsg = map[int]string{
	SUCCESS: "OK",
	ERR:     "FAIL",

	ERROR_NAME_USED:      "用户名已经存在",
	ERROR_PASSWORD_WRONG: "密码错误",
	ERROR_USER_NOT_EXIST: "用户不存在",
	ERROR_TOKEN_EXIST:    "TOKEN不存在",
	ERROR_TOKEN_RUNTIME:  "TOKEN已过期",
	ERROR_TOKEN_WRONG:    "TOKEN错误",
	ERROR_TOKEN_TYPE:     "TOKEN格式错误",
	USER_ROLE_COM:        "无登陆权限",
}

func GetErrorMsg(code int) string {
	return codeMsg[code]
}
