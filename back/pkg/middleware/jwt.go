package middleware

import (
	"github.com/HaHadaxigua/melancholy/pkg/msg"
	"github.com/HaHadaxigua/melancholy/pkg/tools"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//AuthHeader 绑定的请求头
type AuthHeader struct {
	AccessToken string `header:"Access-Token"`
}

// JWT中间件 fixme: 重新生成token后 原来的token仍然有效， 即：修复token的过期机制
func JWT(c *gin.Context) {
	status := msg.OK

	ah := AuthHeader{}

	if err := c.ShouldBindHeader(&ah); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": msg.AuthAccessTokenIllegalErrorMsg,
		})
		return
	}

	if ah.AccessToken == "" {
		status = msg.BadRequest
		status.Cause = msg.AuthAccessTokenIllegalErrorMsg
	} else {
		claims, err := tools.ParseToken(ah.AccessToken)
		if err != nil {
			status = msg.AuthCheckTokenErr
			status.Cause = msg.AuthAccessTokenIllegalErrorMsg
		} else if time.Now().Unix() > claims.ExpiresAt {
			status = msg.AuthCheckTokenTimeoutErr
		}
	}

	if status != msg.OK {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":  status.Code,
			"msg":   status.Message,
			"cause": status.Cause,
		})
		c.Abort()
		return
	}
	c.Next()
}
