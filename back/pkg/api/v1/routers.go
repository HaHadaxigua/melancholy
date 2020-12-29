package v1

import (
	"github.com/HaHadaxigua/melancholy/pkg/api/v1/auth"
	"github.com/HaHadaxigua/melancholy/pkg/api/v1/file"
	"github.com/HaHadaxigua/melancholy/pkg/api/v1/roles"
	"github.com/HaHadaxigua/melancholy/pkg/api/v1/user"
	"github.com/HaHadaxigua/melancholy/pkg/consts"
	"github.com/HaHadaxigua/melancholy/pkg/middleware"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

//SetupRouters 设置gin的路由
func SetupRouters(e *gin.Engine) {
	// cors

	//支持跨域
	e.Use(middleware.Cors)
	// swagger相关
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 注册、登录、登出

	v1 := e.Group(consts.ApiV1)
	{
		auth.SetupAuthRouters(v1)
		roles.SetupRoleRouters(v1)
		file.SetupFileRouters(v1)
		user.SetupUserRouters(v1)
	}
}
