package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func UserLog(info string) {
	_, _ = fmt.Fprintf(gin.DefaultWriter, "user info: %s\n", info)
}
