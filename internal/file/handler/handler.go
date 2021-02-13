package handler

import (
	"github.com/HaHadaxigua/melancholy/internal/basic/middleware"
	"github.com/gin-gonic/gin"
)

func SetupFileRouters(r gin.IRouter) {
	// open

	// secured
	secured := r.Group("/file", middleware.JWT)
	// 文件夹
	secured.POST("/create")
	secured.GET("/folders/:pid")
}
