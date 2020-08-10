package service_user

import (
	"github.com/lyd2/live/models"
	"github.com/lyd2/live/pkg/code"
	"github.com/lyd2/live/util"
)

func Login(u *models.User) *util.Response {

	if !u.CheckAuth() {
		return util.Error(code.ERROR_USERNAME_OR_PASSWD)
	}

	// 生成jwt
	token, err := util.GenerateToken(u.Username)
	if err != nil {
		return util.ServerError()
	}

	return util.Success(map[string]interface{}{
		"token": token,
	})

}
