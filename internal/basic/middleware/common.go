/******
** @date : 3/7/2021 12:31 PM
** @author : zrx
** @description: some common middlewares
******/
package middleware

import (
	"github.com/HaHadaxigua/melancholy/internal/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Forbidden() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Abort()
		c.JSON(http.StatusForbidden, response.NewErr(nil))
		return
	}
}
