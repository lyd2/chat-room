package message

// 用户发送消息的格式
type ReadMsg struct {
	MsgType      int    `json:"msgType" validate:"oneof=1 2 3"`
	Content      string `json:"content" validate:"max=200"`
	ReceiverType int    `json:"receiverType" validate:"oneof=1 2 3"`
	ReceiverID   int    `json:"receiverID"`
	UserId       int    `json:"-"`
	Username     string `json:"-"`
}
