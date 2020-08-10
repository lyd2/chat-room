package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lyd2/live/controllers/v1/live"
	v1_room "github.com/lyd2/live/controllers/v1/room"
	v1_user "github.com/lyd2/live/controllers/v1/user"
	"github.com/lyd2/live/middleware"
	"github.com/lyd2/live/pkg/setting"
)

func InitRouter() *gin.Engine {
	engine := gin.Default()

	gin.SetMode(setting.RunMode)

	engine.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	// 用户接口
	v1UserGroup := engine.Group("/v1/user")
	v1UserGroup.POST("/login", v1_user.Login)
	v1UserGroup.POST("/register", v1_user.Register)
	v1UserGroup.PUT("/pwdchange", v1_user.PwdChange)
	v1UserGroup.Use(middleware.JWT()).GET("/info", v1_user.Info)

	// 直播间接口
	v1RoomGroup := engine.Group("/v1/room")
	v1RoomGroup.Use(middleware.JWT())
	v1RoomGroup.GET("/info/:roomId", v1_room.Info)
	v1RoomGroup.GET("/list", v1_room.List)

	// 直播websocket接口
	v1LiveGroup := engine.Group("/v1/live")
	v1LiveGroup.Use(middleware.WebsocketAuth()).Use(middleware.JWT())
	v1LiveGroup.GET("/liveConnect/:roomId", live.LiveConnect)

	return engine
}
