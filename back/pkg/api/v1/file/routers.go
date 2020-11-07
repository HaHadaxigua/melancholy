package file

import (
	"github.com/HaHadaxigua/melancholy/pkg/msg"
	service "github.com/HaHadaxigua/melancholy/pkg/service/v1/file"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SetupFileRouters 设置file模块的嵌套路由组
func SetupFileRouters(r *gin.RouterGroup) {
	fr := r.Group("/file")
	{
		// Test
		fr.GET("/download", Hello)

		// 创建文件夹
		fr.POST("/create", CreateFolder)
		fr.GET("/path", GetDirs)

		// 视频文件
		fr.POST("/upload/video", UploadVideoFile)
		fr.GET("/download/video", DownloadVideoFile)
	}
}

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "hello world",
	})
}

//CreateFolder 创建文件夹
func CreateFolder(c *gin.Context) {
	dReq := &msg.DirRequest{}
	err := c.BindJSON(dReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": msg.BadRequest,
		})
		return
	}
	bf, err := service.CreateFolder(dReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": err,
		})
	}
	name,_ := bf.CreatedAt.Zone()
	c.JSON(http.StatusOK, gin.H{
		"data": bf,
		"time/zone": name,
	})
}

// GetDirs 根据路径获取文件夹
func GetDirs(c *gin.Context) {

}

// UploadFile 上传文件
func UploadVideoFile(c *gin.Context) {

}

//DownloadVideoFile 下载文件
func DownloadVideoFile(c *gin.Context) {

}
