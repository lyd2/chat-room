package util

import (
	"github.com/gin-gonic/gin"
	"github.com/lyd2/live/pkg/setting"
	"strconv"
)

func GetPage(c *gin.Context) int {

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		return 0
	}

	return (page - 1) * setting.PageSize

}
