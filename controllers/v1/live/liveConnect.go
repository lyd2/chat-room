package live

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/lyd2/live/models"
	"github.com/lyd2/live/pkg/code"
	"github.com/lyd2/live/pkg/manager"
	"github.com/lyd2/live/pkg/manager/message"
	service_room "github.com/lyd2/live/services/room"
	service_user "github.com/lyd2/live/services/user"
	"github.com/lyd2/live/util"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
	"strconv"
	"time"
)

func LiveConnect(context *gin.Context) {

	// 查询用户信息
	var user models.User
	username, _ := context.Get("username")
	user.Username = username.(string)

	userInfo := service_user.Info(&user)
	if userInfo.Code != 200 {
		context.JSON(http.StatusUnauthorized, userInfo)
		return
	}

	// 获取并验证直播间信息
	roomId, err := strconv.Atoi(context.Param("roomId"))
	if err != nil {
		context.JSON(http.StatusOK, util.Error(code.INVALID_PARAMS))
		return
	}
	roomInfo := service_room.Info(&models.Room{
		ID: roomId,
	})
	if roomInfo.Code != 200 {
		context.JSON(http.StatusNotFound, roomInfo)
		return
	}

	// 升级为 websocket
	upgrader := websocket.Upgrader{
		ReadBufferSize:  10240,
		WriteBufferSize: 10240,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	ws, err := upgrader.Upgrade(context.Writer, context.Request, nil)
	if err != nil {
		log.Println(err.Error())
		//context.JSON(http.StatusInternalServerError, util.ServerError())
		return
	}

	// 分别开启此连接的读写线程
	connection := manager.NewConnection(user.ID, user.Username, roomId)
	if !manager.LiveRooms.AddConn(roomId, connection) {
		return
	}
	connection.SetConnId(uuid.NewV4().String())
	readThread(ws, connection)
	writeThread(ws, connection)

	util.UserLog(fmt.Sprintf("连接成功，connId=%s", connection.ConnId()))
}

func readThread(ws *websocket.Conn, connection *manager.Connection) {

	go func() {
		for {
			var readMsg message.ReadMsg
			var err error

			// 读取不用超时
			//err = ws.SetReadDeadline(time.Now().Add(30 * time.Second))
			//if err != nil {
			//	connection.Close()
			//	_ = ws.Close()
			//	return
			//}

			//err = ws.ReadJSON(&readMsg)
			t, p, err := ws.ReadMessage()
			if err != nil {
				//fmt.Println(err)
				connection.Close()
				_ = ws.Close()
				return
			}
			if t != websocket.TextMessage {
				continue
			}
			err = json.Unmarshal(p, &readMsg)
			if err != nil {
				util.UserLog(fmt.Sprintf("connId=%s 消息 json 解析失败", connection.ConnId()))
				continue
			}
			readMsg.UserId = connection.UserId()
			readMsg.Username = connection.UserName()

			// 验证客户端发送的消息是否正确
			resp := util.Validator(&readMsg)
			if resp != nil {
				continue
			}

			util.UserLog(fmt.Sprintf("connId=%s, UP: %v", connection.ConnId(), readMsg))

			up := manager.LiveRooms.GetUpChannel(connection.RoomId())
			if up == nil {
				// 直播间不存在
				_ = ws.Close()
				return
			}

			up <- &readMsg
		}
	}()

}

func writeThread(ws *websocket.Conn, connection *manager.Connection) {

	go func() {
		for {
			var err error
			var msg *message.Message

			/*
				假设读协程因网络故障，读取失败，因此退出了
				但假如它没有发送任何消息，而也没有任何其它用户发送消息，此时写协程会被阻塞在这里，而无法退出
				因此在连接的 Close 方法里，会主动往 down 通道写入一个 nil，来使得此协程退出
				当然，如果用户所在直播间有任意一个用户发送消息了，自然也会通知到它，导致它检测到网络已失效而退出
			*/
			down := connection.GetDownChannel()
			msg = <-down
			if msg == nil {
				_ = ws.Close()
				return
			}

			// 写入超时时间
			err = ws.SetWriteDeadline(time.Now().Add(30 * time.Second))
			if err != nil {
				connection.Close()
				_ = ws.Close()
				return
			}

			err = ws.WriteJSON(&msg)
			if err != nil {
				connection.Close()
				_ = ws.Close()
				return
			}
		}
	}()

}
