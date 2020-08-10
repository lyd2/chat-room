package code

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004

	ERROR_USERNAME           = 30001
	ERROR_PASSWORD           = 30002
	ERROR_USERNAME_OR_PASSWD = 30003
	USERNAME_ALREADY_EXISTS  = 30004
	ERROR_USER_EMPTY         = 30005

	ERROR_ROOM_EMPTY = 40001
)