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

		fr.POST("/create", CreateFolder) // 创建文件夹
		fr.GET("/folders", GetFolders)   // 获取当前文件夹下的子文件夹
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
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": bf,
	})
}

// GetFolders 获取当前目录下的文件夹
func GetFolders(c *gin.Context){
	curPath := c.GetInt("curPath")
	folders, err := service.GetChildFolders(curPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": folders,
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
