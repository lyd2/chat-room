package pubsub

import (
	"github.com/lyd2/live/pkg/manager/message"
)

// 全局消息总线需要实现的接口
type PubSub interface {
	Push(message *message.Message)
	CreateConsumer(in chan<- *message.Message)
	DelConsumer(in chan<- *message.Message)
}
