package file

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// SetupFileRouters 设置file模块的嵌套路由组
func SetupFileRouters(r *gin.RouterGroup) {
	fr := r.Group("/file")
	{
		fr.GET("/download", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "hello world",
			})
		})
	}
}
