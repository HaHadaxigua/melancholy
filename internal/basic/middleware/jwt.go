package middleware

import (
	"github.com/HaHadaxigua/melancholy/ent"
	"github.com/HaHadaxigua/melancholy/internal/basic/service"
	"github.com/HaHadaxigua/melancholy/internal/basic/tools"
	"github.com/HaHadaxigua/melancholy/internal/global/response"
	"github.com/HaHadaxigua/melancholy/internal/log/store"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// AuthHeader
type AuthHeader struct {
	AccessToken string `header:"Access-Token"`
}

// JWT
func JWT(c *gin.Context) {
	status := response.OK

	ah := AuthHeader{}

	if err := c.ShouldBindHeader(&ah); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"response": response.AuthAccessTokenIllegalErrorMsg,
		})
		return
	}

	if ah.AccessToken == "" {
		status = response.BadRequest
		status.Data = response.AuthAccessTokenIllegalErrorMsg
	} else {
		claims, err := tools.JwtParseToken(ah.AccessToken)
		if err != nil {
			status = response.AuthCheckTokenErr
			status.Data = response.AuthAccessTokenIllegalErrorMsg
		} else if time.Now().Unix() > claims.ExpiresAt {
			status = response.AuthCheckTokenTimeoutErr
		} else { // 此时token是有效的
			// 判断下token是否已经进入黑名单
			el, errr := store.ExitLogStore.GetExitLog(ah.AccessToken)
			if el != nil {
				c.JSON(http.StatusBadRequest, response.UserExitErr)
				c.Abort()
			}
			if errr != nil {
				if !ent.IsNotFound(errr) {
					c.JSON(http.StatusInternalServerError, response.UnKnown)
					c.Abort()
				}
			}
			userId := service.UserService.CheckUserExist(claims.Email, claims.Password)
			c.Set("user_id", userId)
		}
	}

	if status != response.OK {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":  status.Code,
			"response":   status.Message,
			"cause": status.Data,
		})
		c.Abort()
		return
	}

	c.Next()
}
