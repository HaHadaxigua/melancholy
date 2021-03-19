package middleware

import (
	"github.com/HaHadaxigua/melancholy/internal/basic/msg"
	"github.com/HaHadaxigua/melancholy/internal/basic/service"
	"github.com/HaHadaxigua/melancholy/internal/basic/tools"
	"github.com/HaHadaxigua/melancholy/internal/consts"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

type authHeader struct {
	AccessToken string `header:"Access-Token"`
}

func Auth(c *gin.Context) {
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
	_user, err := service.User.GetUserByID(uid, true)
	if err != nil {
		c.Abort()
		logrus.Info("get user failed")
		return
	}
	c.Set(consts.User, _user)

	c.Next()
}
