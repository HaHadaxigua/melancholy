package api

import (
	"github.com/HaHadaxigua/melancholy/pkg/api/v1/file"
	"github.com/HaHadaxigua/melancholy/pkg/api/v1/user"
	"github.com/HaHadaxigua/melancholy/pkg/conf"
	"github.com/gin-gonic/gin"
)

//SetupRouters 设置gin的路由
func SetupRouters(e *gin.Engine) {
	e.Use(gin.Logger())
	e.Use(gin.Recovery())
	gin.SetMode(conf.C.Mode)
	e.GET("/auth", GetAuth)

	v1 := e.Group("/api/v1")
	file.SetupRouters(v1)
	user.SetupRouters(v1)

	// 注册、登录、登出相关


}
