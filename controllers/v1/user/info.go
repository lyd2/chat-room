package v1_user

import (
	"github.com/gin-gonic/gin"
	"github.com/lyd2/live/models"
	service_user "github.com/lyd2/live/services/user"
	"net/http"
)

func Info(context *gin.Context) {

	var user models.User
	username, _ := context.Get("username")
	user.Username = username.(string)

	context.JSON(http.StatusOK, service_user.Info(&user))

}
