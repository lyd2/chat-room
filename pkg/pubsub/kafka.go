package pubsub

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/lyd2/live/pkg/manager/message"
	"github.com/lyd2/live/pkg/setting"
	"github.com/lyd2/live/util"
	"log"
	"strings"
	"sync"
)

type KafkaClient struct {
	// 多个地址逗号分隔
	Hosts         string
	Topic         string
	ProducerConn  sarama.SyncProducer
	ConsumerConn  sarama.Consumer
	Partition     int32
	ConsumerGroup sarama.PartitionConsumer
	Channels      *sync.Map
}

var KafkaConn KafkaClient

func init() {

	sec, err := setting.Cfg.GetSection("kafka")
	if err != nil {
		log.Fatal(2, "Fail to get section 'kafka': %v", err)
	}

	KafkaConn.Hosts = sec.Key("HOSTS").String()
	KafkaConn.Topic = sec.Key("TOPIC").String()

	// 连接生产者
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err := sarama.NewSyncProducer(
		strings.Split(KafkaConn.Hosts, ","),
		config,
	)
	if err != nil {
		log.Fatal(fmt.Sprintf("producer closed, err:", err))
		return
	}
	KafkaConn.ProducerConn = client

	// 连接消费者
	consumer, err := sarama.NewConsumer(
		strings.Split(KafkaConn.Hosts, ","),
		nil,
	)
	if err != nil {
		log.Fatal(fmt.Sprintf("fail to start consumer, err:%v\n", err))
		return
	}
	KafkaConn.ConsumerConn = consumer

	// 获取 partition
	partitionList, err := KafkaConn.ConsumerConn.Partitions(KafkaConn.Topic)
	if err != nil {
		log.Fatal(fmt.Sprintf("fail to get list of partition:err%v", err))
		return
	}
	KafkaConn.Partition = partitionList[0]

	// 创建一个对应的分区消费者
	cGroup, err := consumer.ConsumePartition(
		KafkaConn.Topic,
		KafkaConn.Partition,
		sarama.OffsetNewest,
	)
	if err != nil {
		log.Fatal(
			fmt.Sprintf(
				"failed to start consumer for partition %d,err:%v",
				KafkaConn.Partition, err,
			),
		)
		return
	}
	KafkaConn.ConsumerGroup = cGroup

	// 推送通道群组
	KafkaConn.Channels = &sync.Map{}

	// 开启消息总线的消费者协程
	KafkaConn.start()
}

func (k KafkaClient) Push(msg *message.Message) {

	util.UserLog(fmt.Sprintf("写入 pubsub: %v", msg))

	value, err := json.Marshal(msg)
	if err != nil {
		util.UserLog("kafka 创建消息失败")
		return
	}

	// 构造一个消息
	kafkaMsg := &sarama.ProducerMessage{}
	kafkaMsg.Topic = k.Topic
	kafkaMsg.Value = sarama.ByteEncoder(value)

	// 发送消息
	pid, offset, err := k.ProducerConn.SendMessage(kafkaMsg)
	if err != nil {
		util.UserLog(fmt.Sprintf("send msg failed, err:", err))
		return
	}
	util.UserLog(fmt.Sprintf("pid:%v offset:%v", pid, offset))
}

func (k KafkaClient) CreateConsumer(in chan<- *message.Message) {
	k.Channels.Store(in, 0)
}

func (k KafkaClient) DelConsumer(in chan<- *message.Message) {
	k.Channels.Delete(in)
}

func (k KafkaClient) start() {

	go func() {

		for {

			kafkaMsg := <-k.ConsumerGroup.Messages()
			value := kafkaMsg.Value

			msg := &message.Message{}
			err := json.Unmarshal(value, msg)
			if err != nil {
				util.UserLog("解析 kafka 消息失败")
				continue
			}

			util.UserLog(fmt.Sprintf("从 pubsub 拉取: %v", msg))

			// 推入通道列表
			k.Channels.Range(func(key, value interface{}) bool {
				c, ok := key.(chan<- *message.Message)
				if !ok || c == nil {
					return true
				}

				c <- msg
				return true
			})

		}

	}()

}
