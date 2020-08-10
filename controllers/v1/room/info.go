package v1_room

import (
	"github.com/gin-gonic/gin"
	"github.com/lyd2/live/models"
	"github.com/lyd2/live/pkg/code"
	service_room "github.com/lyd2/live/services/room"
	"github.com/lyd2/live/util"
	"net/http"
	"strconv"
)

func Info(context *gin.Context) {

	var err error
	var room models.Room

	// 获取用户数据
	if room.ID, err = strconv.Atoi(context.Param("roomId")); err != nil {
		context.JSON(http.StatusOK, util.Error(code.ERROR))
		return
	}

	//fmt.Println(user)

	// 验证是否字段是否正确
	resp := util.ValidatorPartial(&room, "ID")
	if resp != nil {
		context.JSON(http.StatusOK, resp)
		return
	}

	// 获取直播间信息
	context.JSON(http.StatusOK, service_room.Info(&room))

}
