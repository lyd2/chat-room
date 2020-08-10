package manager

import (
	"github.com/lyd2/live/models"
	"github.com/lyd2/live/pkg/pubsub"
	"github.com/lyd2/live/util"
)

func init() {

	rooms := new(models.Room).All()

	util.UserLog("初始化，创建所有直播间...")

	for _, room := range rooms {
		LiveRooms.AddRoom(room.ID, pubsub.KafkaConn)
	}

}
