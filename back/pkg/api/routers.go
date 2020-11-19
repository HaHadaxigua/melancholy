package api

import (
	"github.com/HaHadaxigua/melancholy/pkg/api/v1/file"
	"github.com/HaHadaxigua/melancholy/pkg/api/v1/user"
	"github.com/HaHadaxigua/melancholy/pkg/middleware"
	"github.com/gin-gonic/gin"
)

//SetupRouters 设置gin的路由
func SetupRouters(e *gin.Engine) {
	// 注册、登录、登出相关
	e.GET("/auth", Login)
	e.POST("/register", Register) // 注册用户

	v1 := e.Group("/api/v1", middleware.JWT)
	{
		file.SetupRouters(v1)
		user.SetupRouters(v1)
	}

}
