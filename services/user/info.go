package service_user

import (
	"github.com/lyd2/live/models"
	"github.com/lyd2/live/pkg/code"
	"github.com/lyd2/live/util"
)

func Info(r *models.User) *util.Response {
	if r.Info() {
		return util.Success(r)
	}
	return util.Error(code.ERROR_USER_EMPTY)
}
