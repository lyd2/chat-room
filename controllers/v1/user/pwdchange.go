package v1_user

import (
	"github.com/gin-gonic/gin"
	"github.com/lyd2/live/models"
	"github.com/lyd2/live/pkg/code"
	service_user "github.com/lyd2/live/services/user"
	"github.com/lyd2/live/util"
	"net/http"
)

func PwdChange(context *gin.Context) {

	var err error
	var user models.User
	var newUser models.User

	// 获取用户数据和新密码
	info := struct {
		models.User
		NewPassword string `json:"newPassword"`
	}{}
	if err = context.ShouldBindJSON(&info); err != nil {
		context.JSON(http.StatusOK, util.Error(code.ERROR))
		return
	}

	user.Username = info.Username
	user.Password = info.Password
	user.Phone = info.Phone
	newUser.Password = info.NewPassword

	//fmt.Println(user)

	// 验证是否字段是否正确
	resp := util.ValidatorExcept(&user, "ID")
	if resp != nil {
		context.JSON(http.StatusOK, resp)
		return
	}
	resp = util.ValidatorPartial(&newUser, "Password")
	if resp != nil {
		context.JSON(http.StatusOK, resp)
		return
	}

	// 执行修改密码操作
	context.JSON(http.StatusOK, service_user.PwdChange(&user, &newUser))

}
