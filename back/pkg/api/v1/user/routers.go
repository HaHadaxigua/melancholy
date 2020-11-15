package user

import "github.com/gin-gonic/gin"

// 设置user模块的嵌套路由
func SetupRouters(r *gin.RouterGroup) {
	fr := r.Group("/user")
	{
		// 文件夹
		fr.POST("/register", Register) // 注册用户
	}
}
