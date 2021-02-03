package handler

import (
	"github.com/HaHadaxigua/melancholy/internal/basic/middleware"
	"github.com/HaHadaxigua/melancholy/internal/file/msg"
	"github.com/HaHadaxigua/melancholy/internal/file/service"
	"github.com/HaHadaxigua/melancholy/internal/global/consts"
	"github.com/HaHadaxigua/melancholy/internal/global/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// SetupFileRouters 设置file模块的嵌套路由组
func SetupFileRouters(r gin.IRouter) {
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
	req := &msg.ReqFolderCreate{}
	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	req.UserID = c.GetInt(consts.UserID)
	if err := service.FolderService.Create(req); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, response.OK)
}

/**
ListFolders cid represents the current path
文件表需要用字符ID以及用户ID作为主键
*/
func ListFolders(c *gin.Context) {
	pid, err := strconv.Atoi(c.Param("cid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.InvalidParamsErr)
		return
	}
	uid := c.GetInt(consts.UserID)
	folders, err := service.FolderService.ListFolder(uid, pid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, response.OkResp(folders))
}
