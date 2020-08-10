package code

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",

	ERROR_USERNAME:           "用户名错误",
	ERROR_PASSWORD:           "密码错误",
	ERROR_USERNAME_OR_PASSWD: "用户名或密码错误",
	USERNAME_ALREADY_EXISTS:  "用户名已存在",
	ERROR_USER_EMPTY:         "用户不存在",

	ERROR_ROOM_EMPTY: "直播间不存在",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
