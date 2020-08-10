package middleware

import (
	"github.com/gin-gonic/gin"
)

func WebsocketAuth() gin.HandlerFunc {

	return func(context *gin.Context) {

		// TODO: TOKEN做加密/解密操作，具体流程如下：
		/*
			在发起websocket请求前，先拿jwt令牌去调一个接口生成token
			token存储到redis中，以满足集群环境的需求
			这里再将token还原成jwt
		*/

		token := context.Query("token")
		context.Set("Authorization", token)

		context.Next()

	}
}
