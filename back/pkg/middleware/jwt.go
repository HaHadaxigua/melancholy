package middleware

import (
	"github.com/HaHadaxigua/melancholy/pkg/msg"
	"github.com/HaHadaxigua/melancholy/pkg/store"
	"github.com/HaHadaxigua/melancholy/pkg/store/user"
	"github.com/HaHadaxigua/melancholy/pkg/tools"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//AuthHeader 绑定的请求头
type AuthHeader struct {
	AccessToken string `header:"Access-Token"`
}

// JWT中间件
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
		status.Data = msg.AuthAccessTokenIllegalErrorMsg
	} else {
		claims, err := tools.ParseToken(ah.AccessToken)
		if err != nil {
			status = msg.AuthCheckTokenErr
			status.Data = msg.AuthAccessTokenIllegalErrorMsg
		} else if time.Now().Unix() > claims.ExpiresAt {
			status = msg.AuthCheckTokenTimeoutErr
		} else { // 此时token是有效的
			// 判断下token是否已经进入黑名单
			el, e := store.FindExitLog(ah.AccessToken)
			if el != nil || e != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error":msg.UserExitErr,
				})
				c.Abort()
			}

			userId := user.CheckUserExist(claims.Email, claims.Password)
			c.Set("user_id", userId)
		}
	}

	if status != msg.OK {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":  status.Code,
			"msg":   status.Message,
			"cause": status.Data,
		})
		c.Abort()
		return
	}

	c.Next()
}
