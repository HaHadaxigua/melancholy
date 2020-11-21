package user

import "github.com/gin-gonic/gin"

// 设置user模块的嵌套路由
func SetupRouters(r *gin.RouterGroup) {
	fr := r.Group("/users")
	{
		fr.GET("/logout", logout)
	}
}
