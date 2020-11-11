package api

import (
	"github.com/HaHadaxigua/melancholy/pkg/api/v1/file"
	"github.com/gin-gonic/gin"
	"net/http"
)

//SetupRouters 设置gin的路由
func SetupRouters(e *gin.Engine) {
	v1 := e.Group("v1")
	{
		v1.GET("/hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"hello": "world",
			})
		})
	}
	file.SetupRouters(v1)
}
