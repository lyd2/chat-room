package manager

import (
	"fmt"
	"github.com/lyd2/live/pkg/manager/message"
	"github.com/lyd2/live/util"
)

type Connection struct {
	connId string

	userId   int
	userName string
	roomId   int

	closed bool

	down chan *message.Message
}

func NewConnection(userId int, userName string, roomId int) *Connection {

	return &Connection{
		userId:   userId,
		userName: userName,
		roomId:   roomId,

		down: make(chan *message.Message, 2048),
	}

}

func (c *Connection) ConnId() string {
	return c.connId
}

func (c *Connection) SetConnId(connId string) {
	c.connId = connId
}

func (c *Connection) UserId() int {
	return c.userId
}

func (c *Connection) UserName() string {
	return c.userName
}

func (c *Connection) RoomId() int {
	return c.roomId
}

func (c *Connection) Close() {
	c.closed = true
	//close(c.down)
	// 关闭写协程
	c.down <- nil
	util.UserLog(fmt.Sprintf("关闭连接, user_id=%d, room_id=%d", c.userId, c.roomId))
}

func (c *Connection) Closed() bool {
	return c.closed
}

// 用户线程从 down 通道获取待发送的消息，并发送给用户
// 以用户连接分组，每个用户连接都有一个自己的通道
func (c *Connection) GetDownChannel() <-chan *message.Message {
	return c.down
}
