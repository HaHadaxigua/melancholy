package middleware

import (
	"github.com/HaHadaxigua/melancholy/internal/basic/msg"
	"github.com/HaHadaxigua/melancholy/internal/basic/tools"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type authHeader struct {
	AccessToken string `header:"Access-Token"`
}

func JWT(c *gin.Context) {
	ah := authHeader{}

	if err := c.ShouldBindHeader(&ah); err != nil {
		c.JSON(http.StatusBadRequest, msg.ErrAuthAccessTokenIllegal)
		return
	}

	if ah.AccessToken == "" {
		c.JSON(http.StatusBadRequest, msg.ErrAuthAccessTokenIllegal)
		return
	}

	claims, err := tools.JwtParseToken(ah.AccessToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, msg.ErrAuthAccessTokenIllegal)
		return
	}

	if time.Now().Unix() > claims.ExpiresAt {
		c.JSON(http.StatusBadRequest, msg.ErrAuthCheckTokenTimeout)
		return
	}

	c.Next()
}
