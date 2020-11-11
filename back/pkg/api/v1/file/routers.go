package file

import (
	"github.com/gin-gonic/gin"
)

// SetupRouters 设置file模块的嵌套路由组
func SetupRouters(r *gin.RouterGroup) {
	fr := r.Group("/file")
	{
		// Test
		fr.GET("/download", Hello)

		// 文件夹
		fr.POST("/create", CreateFolder) // 创建文件夹
		fr.GET("/folders", GetFolders)   // 获取当前文件夹下的子文件夹
		fr.GET("/path", GetDirs)
		// 视频文件
		fr.POST("/upload/video", UploadVideoFile)
		fr.GET("/download/video", DownloadVideoFile)
	}
}
