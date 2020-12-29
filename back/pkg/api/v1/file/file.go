package file

import (
	"github.com/HaHadaxigua/melancholy/pkg/consts"
	"github.com/HaHadaxigua/melancholy/pkg/msg"
	service "github.com/HaHadaxigua/melancholy/pkg/service/v1"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//CreateFolder 创建文件夹
func CreateFolder(c *gin.Context) {
	req := &msg.CreateFolderReq{}
	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	uid := c.GetInt(consts.UserID)
	r := &msg.FolderRequest{
		Creator: uid,
		Name: req.Name,
		ParentId: req.ParentID,
	}
	err = service.CreateFolder(r)
	if err != nil {
		c.JSON(http.StatusInternalServerError, msg.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, msg.OK)
}

// GetFolders 获取当前目录下的文件夹
func GetFolders(c *gin.Context) {
	pid, err := strconv.Atoi(c.Query("pid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, msg.InvalidParamsErr)
		return
	}
	uid := c.GetInt(consts.UserID)
	folders, err := service.ListCurrentFolder(uid, pid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, msg.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, folders)
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
