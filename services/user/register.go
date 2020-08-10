package service_user

import (
	"github.com/lyd2/live/models"
	"github.com/lyd2/live/pkg/code"
	"github.com/lyd2/live/util"
)

func Register(u *models.User) *util.Response {
	if u.Insert() {
		return util.Success("")
	}

	return util.Error(code.USERNAME_ALREADY_EXISTS)
}
