package api

import (
	"github.com/HaHadaxigua/melancholy/pkg/api/v1/file"
	"github.com/HaHadaxigua/melancholy/pkg/api/v1/roles"
	"github.com/HaHadaxigua/melancholy/pkg/api/v1/user"
	"github.com/HaHadaxigua/melancholy/pkg/middleware"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

//SetupRouters 设置gin的路由
func SetupRouters(e *gin.Engine) {
	// swagger相关
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 注册、登录
	SetupBasicRouters(e)

	v1 := e.Group("/api/v1", middleware.JWT, middleware.Authorize)
	{
		roles.SetupRouters(v1)
		file.SetupRouters(v1)
		user.SetupRouters(v1)
	}

}
