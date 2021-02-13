package middleware

import (
	"github.com/HaHadaxigua/melancholy/internal/basic/msg"
	"github.com/HaHadaxigua/melancholy/internal/basic/tools"
	"github.com/HaHadaxigua/melancholy/internal/global/consts"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type authHeader struct {
	AccessToken string `header:"Access-Token"`
}

func JWT(c *gin.Context) {
	ah := authHeader{}

	if err := c.ShouldBindHeader(&ah); err != nil {
		c.JSON(http.StatusBadRequest, msg.ErrAuthAccessTokenIllegal)
		c.Abort()
		return
	}

	if ah.AccessToken == "" {
		c.JSON(http.StatusBadRequest, msg.ErrAuthAccessTokenIllegal)
		c.Abort()
		return
	}

	claims, err := tools.JwtParseToken(ah.AccessToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, msg.ErrAuthAccessTokenIllegal)
		c.Abort()
		return
	}

	if time.Now().Unix() > claims.ExpiresAt {
		c.JSON(http.StatusBadRequest, msg.ErrAuthCheckTokenTimeout)
		c.Abort()
		return
	}

	uidStr := claims.Id
	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		c.Abort()
		return
	}
	c.Set(consts.UserID, uid)
	c.Next()
}
