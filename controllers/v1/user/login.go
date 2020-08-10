package v1_user

import (
	"github.com/gin-gonic/gin"
	"github.com/lyd2/live/models"
	"github.com/lyd2/live/pkg/code"
	service_user "github.com/lyd2/live/services/user"
	"github.com/lyd2/live/util"
	"net/http"
)

func Login(context *gin.Context) {

	var err error
	var user models.User

	// 获取用户数据
	if err = context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusOK, util.Error(code.ERROR))
		return
	}

	//fmt.Println(user)

	// 验证是否字段是否正确
	resp := util.ValidatorPartial(&user, "Username", "Password")
	if resp != nil {
		context.JSON(http.StatusOK, resp)
		return
	}

	// 执行登录操作
	context.JSON(http.StatusOK, service_user.Login(&user))

}
