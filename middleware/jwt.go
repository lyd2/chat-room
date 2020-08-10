package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lyd2/live/pkg/code"
	"github.com/lyd2/live/util"
	"net/http"
	"strings"
	"time"
)

func JWT() gin.HandlerFunc {

	return func(context *gin.Context) {

		token := context.GetHeader("Authorization")
		token = strings.Replace(token, "Bearer ", "", 1)

		if token == "" {
			authorization, ok := context.Get("Authorization")
			if !ok {
				exit(context, code.INVALID_PARAMS)
				return
			}
			token = authorization.(string)
		}

		if token == "" {
			exit(context, code.INVALID_PARAMS)
			return
		}

		claims, err := util.ParseToken(token)
		if err != nil {
			exit(context, code.ERROR_AUTH_CHECK_TOKEN_FAIL)
			return
		}

		if time.Now().Unix() > claims.ExpiresAt {
			exit(context, code.ERROR_AUTH_CHECK_TOKEN_TIMEOUT)
			return
		}

		context.Set("username", claims.Username)
		context.Next()
	}
}

func exit(context *gin.Context, c int) {
	context.JSON(http.StatusUnauthorized, util.Error(c))
	context.Abort()
}
