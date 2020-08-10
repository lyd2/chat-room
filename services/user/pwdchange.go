package service_user

import (
	"github.com/lyd2/live/models"
	"github.com/lyd2/live/pkg/code"
	"github.com/lyd2/live/util"
)

func PwdChange(u, nu *models.User) *util.Response {
	if u.PasswdChange(nu) {
		return util.Success("")
	}

	return util.Error(code.ERROR_USERNAME_OR_PASSWD)
}
