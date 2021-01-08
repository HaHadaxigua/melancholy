package v1

import (
	"github.com/HaHadaxigua/melancholy/pkg/consts"
	"github.com/HaHadaxigua/melancholy/pkg/middleware"
	"github.com/HaHadaxigua/melancholy/pkg/msg"
	service "github.com/HaHadaxigua/melancholy/pkg/service/v1"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// SetupFileRouters 设置file模块的嵌套路由组
func SetupFileRouters(r *gin.RouterGroup) {
	// open

	// secured
	//r.Use(middleware.JWT, middleware.Authorize)
	secured := r.Group("/file", middleware.JWT)
	// 文件夹
	secured.POST("/create", CreateFolder)     // 创建文件夹
	secured.GET("/folders/:pid", ListFolders) // 获取当前文件夹下的子文件夹
}

// CreateFolder
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
	err = service.FolderService.CreateFolder(r)
	if err != nil {
		c.JSON(http.StatusInternalServerError, msg.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, msg.OK)
}

// ListFolders
func ListFolders(c *gin.Context) {
	pid, err := strconv.Atoi(c.Param("pid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, msg.InvalidParamsErr)
		return
	}
	uid := c.GetInt(consts.UserID)
	folders, err := service.FolderService.ListCurrentFolder(uid, pid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, msg.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, msg.OkResp(folders))
}
