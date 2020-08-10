package v1_room

import (
	"github.com/gin-gonic/gin"
	"github.com/lyd2/live/models"
	service_room "github.com/lyd2/live/services/room"
	"net/http"
)

func List(context *gin.Context) {

	var room models.Room

	// 获取用户数据
	room.Name = context.DefaultQuery("name", "")

	// 获取直播间列表
	context.JSON(http.StatusOK, service_room.List(context, &room))

}
