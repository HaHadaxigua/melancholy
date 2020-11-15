package middleware

import (
	"github.com/HaHadaxigua/melancholy/pkg/msg"
	"github.com/HaHadaxigua/melancholy/pkg/tools"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// JWT中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var status *msg.Err
		token := c.Query("token")
		if token == "" {
			status = msg.BadRequest
			return
		} else {
			claims, err := tools.ParseToken(token)
			if err != nil {
				status = msg.AuthCheckTokenErr
			} else if time.Now().Unix() > claims.ExpiresAt {
				status = msg.AuthCheckTokenTimeoutErr
			} else {
				status = msg.InternalServerErr
			}
		}
		if status != msg.OK {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": status.Code,
				"msg":  status.Message,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
