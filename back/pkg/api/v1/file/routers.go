package file

import (
	"github.com/HaHadaxigua/melancholy/pkg/middleware"
	"github.com/gin-gonic/gin"
)

// SetupFileRouters 设置file模块的嵌套路由组
func SetupFileRouters(r *gin.RouterGroup) {
	// open

	// secured
	//r.Use(middleware.JWT, middleware.Authorize)
	secured := r.Group("/file", middleware.JWT)
	// 文件夹
	secured.POST("/create", CreateFolder)    // 创建文件夹
	secured.GET("/folders/:pid", GetFolders) // 获取当前文件夹下的子文件夹
	secured.GET("/path", GetDirs)
	// 视频文件
	secured.POST("/upload/video", UploadVideoFile)
	secured.GET("/download/video", DownloadVideoFile)

}
