package models

import (
	"github.com/lyd2/live/pkg/manager/message"
	"time"
)

type Message struct {
	ID           int    `gorm:"primary_key" json:"id" validate:"min=1"`
	MsgType      int    `json:"msgType" validate:"oneof=1 2 3"`
	Content      string `json:"content" validate:"max=200"`
	UserId       int    `json:"userId" validate:"min=1"`
	ReceiverType int    `json:"receiverType" validate:"oneof=1 2 3"`
	ReceiverId   int    `json:"receiverId"`
	RoomId       int    `json:"roomId" validate:"min=1"`
	CreatedAt    int64  `json:"createdAt"`
}

func (Message) Insert(msg *message.Message) int {
	var m = &Message{
		MsgType:      msg.MsgType,
		Content:      msg.Content,
		UserId:       msg.From.ID,
		ReceiverType: msg.To.RecType,
		ReceiverId:   msg.To.ID,
		RoomId:       msg.RoomId,
		CreatedAt:    time.Now().Unix(),
	}

	db.Create(m)
	return m.ID
}
