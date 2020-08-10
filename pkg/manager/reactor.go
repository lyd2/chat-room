package manager

import (
	"fmt"
	"github.com/lyd2/live/models"
	"github.com/lyd2/live/pkg/manager/message"
	"github.com/lyd2/live/pkg/pubsub"
	"github.com/lyd2/live/util"
	"sync"
)

// 一个连接组对应一个直播间
type ConnGroup struct {
	clist  []*Connection
	up     chan *message.ReadMsg
	rwlock *sync.RWMutex
	down   chan *message.Message
}

type RoomManager struct {
	// 直播间对应一个连接列表
	Rooms *sync.Map
}

var LiveRooms = &RoomManager{
	Rooms: &sync.Map{},
}

func initConnGroup() *ConnGroup {
	return &ConnGroup{
		clist:  make([]*Connection, 0),
		up:     make(chan *message.ReadMsg, 20480),
		rwlock: &sync.RWMutex{},
		down:   make(chan *message.Message, 20480),
	}
}

// 用户线程从 websocket 读取到消息，将消息写入 up 通道
// 以直播间分组，每个直播间有一个自己的通道
func (rm *RoomManager) GetUpChannel(roomId int) chan<- *message.ReadMsg {
	room, ok := rm.Rooms.Load(roomId)
	if !ok || room == nil {
		return nil
	}
	return room.(*ConnGroup).up
}

func (rm *RoomManager) AddConn(roomId int, conn *Connection) bool {
	room, ok := rm.Rooms.Load(roomId)
	if !ok || room == nil {
		return false
	}

	nRoom := room.(*ConnGroup)
	nRoom.rwlock.Lock()
	nRoom.clist = append(nRoom.clist, conn)
	nRoom.rwlock.Unlock()

	rm.Rooms.LoadOrStore(roomId, nRoom)
	return true
}

func (rm *RoomManager) AddRoom(roomId int, q pubsub.PubSub) bool {
	room, ok := rm.Rooms.Load(roomId)
	if ok && room != nil {
		return false
	}

	g := initConnGroup()
	rm.Rooms.LoadOrStore(roomId, g)

	util.UserLog(fmt.Sprintf("创建直播间 room_id=%d", roomId))

	// 为此直播间设置消费通道（下行通道）
	q.CreateConsumer(g.down)

	// 开启此 Room 的 up 消息协程
	// 将消息格式化为 *Message 类型
	// 它不断从 up （上行通道）读取消息，然后上传到全局消息总线
	go func() {
		util.UserLog(fmt.Sprintf("直播间 room_id=%d 的 up 线程启动", roomId))

		for {
			readMsg := <-g.up
			if readMsg == nil {
				return
			}

			util.UserLog(fmt.Sprintf("------"))
			util.UserLog(fmt.Sprintf("  直播间 room_id=%d 读取到消息", roomId))
			util.UserLog(fmt.Sprintf("  消息内容：%v", readMsg))

			msg := message.Build().
				SetMsgType(readMsg.MsgType).
				SetContent(readMsg.Content).
				SetRoomId(roomId).
				SetSenderId(readMsg.UserId).
				SetSenderName(readMsg.Username).
				SetReceiverType(readMsg.ReceiverType).
				SetReceiverId(readMsg.ReceiverID)

			util.UserLog(fmt.Sprintf("  转换后的消息：%v", msg))
			util.UserLog(fmt.Sprintf("------"))

			// 写入数据库
			msg.MsgId = models.Message{}.Insert(msg)

			q.Push(msg)
		}
	}()

	// 从全局消息总线获取消息，消息总线在获取到消息后，会推送给直播间的 down 通道
	// 然后分发给特定的直播间的连接
	go func() {

		util.UserLog(fmt.Sprintf("直播间 room_id=%d 的 down 线程启动", roomId))

		for {
			// 从消息总线获取消息
			msg := <-g.down
			if roomId != msg.RoomId {
				continue
			}

			util.UserLog(fmt.Sprintf("------"))
			util.UserLog(fmt.Sprintf("  直播间 room_id=%d 从 pubsub 拉取消息", roomId))
			util.UserLog(fmt.Sprintf("  消息内容：%v", msg))

			// 开始推送给直播间下所有连接
			g.rwlock.RLock()

			for _, conn := range g.clist {
				if !conn.Closed() {

					conn.down <- msg

					util.UserLog(fmt.Sprintf("  推送消息给用户：%d", conn.userId))
					util.UserLog(fmt.Sprintf("------"))

				}
			}

			g.rwlock.RUnlock()
		}

	}()

	return true
}

func (rm *RoomManager) DelRoom(roomId int) {
	room, ok := rm.Rooms.Load(roomId)
	if !ok || room == nil {
		return
	}

	close(room.(*ConnGroup).up)
	rm.Rooms.Delete(roomId)
}
