package message

const (
	ControlMessage = 0
	TextMessage    = 1
	ImageMessage   = 2
	EmojiMessage   = 3

	ALL        = 1
	AT_USER    = 2
	AT_MESSAGE = 3
)

// 发送给用户的消息格式
type Message struct {
	MsgId   int      `json:"msgId"`
	MsgType int      `json:"msgType"`
	Content string   `json:"content"`
	RoomId  int      `json:"roomId"`
	From    Sender   `json:"from"`
	To      Receiver `json:"to"`
}

type Sender struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type Receiver struct {
	// 接收者类型
	// 若为 ALL 则不使用 ID， AT_USER 则针对一个用户， AT_MESSAGE 则针对某条消息
	RecType int `json:"recType"`
	ID      int `json:"id"`
}

func Build() *Message {
	return &Message{
		MsgType: 0,
		Content: "",
		From:    Sender{},
		To:      Receiver{},
	}
}

func (m *Message) SetMsgType(msgType int) *Message {
	m.MsgType = msgType
	return m
}

func (m *Message) SetContent(content string) *Message {
	m.Content = content
	return m
}

func (m *Message) SetRoomId(roomId int) *Message {
	m.RoomId = roomId
	return m
}

func (m *Message) SetSenderId(id int) *Message {
	m.From.ID = id
	return m
}

func (m *Message) SetSenderName(name string) *Message {
	m.From.Username = name
	return m
}

func (m *Message) SetReceiverType(t int) *Message {
	m.To.RecType = t
	return m
}

func (m *Message) SetReceiverId(id int) *Message {
	m.To.ID = id
	return m
}
