package service_room

import (
	"github.com/gin-gonic/gin"
	"github.com/lyd2/live/models"
	"github.com/lyd2/live/pkg/setting"
	"github.com/lyd2/live/util"
)

func List(c *gin.Context, r *models.Room) *util.Response {

	var where models.Room
	if r.Name != "" {
		where.Name = r.Name + "%"
	}

	data := r.Rooms(util.GetPage(c), setting.PageSize, where)
	total := r.Total(where)

	return util.Success(map[string]interface{}{
		"list":  data,
		"total": total,
	})
}
