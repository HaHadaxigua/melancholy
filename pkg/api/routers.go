package api

import (
	api "github.com/HaHadaxigua/melancholy/pkg/api/v1"
	"github.com/HaHadaxigua/melancholy/pkg/consts"
	"github.com/HaHadaxigua/melancholy/pkg/middleware"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// SetupRouters
func SetupRouters(e *gin.Engine) {
	// cors

	//支持跨域
	e.Use(middleware.Cors)
	// swagger-path
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := e.Group(consts.ApiV1)
	{
		api.SetupAuthRouters(v1)
		api.SetupRoleRouters(v1)
		api.SetupFileRouters(v1)
		api.SetupUserRouters(v1)
	}
}
