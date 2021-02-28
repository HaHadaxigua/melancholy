package handler

import (
	"github.com/HaHadaxigua/melancholy/internal/basic/middleware"
	"github.com/HaHadaxigua/melancholy/internal/file/msg"
	"github.com/HaHadaxigua/melancholy/internal/file/service"
	"github.com/HaHadaxigua/melancholy/internal/global/consts"
	"github.com/HaHadaxigua/melancholy/internal/global/response"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func SetupFileRouters(r gin.IRouter) {
	// open

	// secured
	secured := r.Group("/f", middleware.JWT)
	// 文件夹
	folder := secured.Group("/folder")
	folder.POST("/create", createFolder)
	folder.GET("/space", fileSpace)
	folder.PATCH("/info", modifyFolder)
	folder.DELETE("/:id", deleteFolder)

	file := secured.Group("/file")
	file.POST("/create", createFile)
	file.DELETE("/:id", deleteFile)
	file.POST("/upload", uploadFile)
}

func createFolder(c *gin.Context) {
	req := &msg.ReqFolderCreate{}
	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	} else {
		uid := c.GetInt(consts.UserID)
		req.UserID = uid
	}
	if err := service.FileSvc.CreateFolder(req); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, response.Ok(nil))
}

func fileSpace(c *gin.Context) {
	uid := c.GetInt(consts.UserID)
	if rsp, err := service.FileSvc.ListFileSpace(uid); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	} else {
		c.JSON(http.StatusOK, response.Ok(rsp))
	}
}

func modifyFolder(c *gin.Context) {
	req := &msg.ReqFolderUpdate{}
	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	} else {
		req.UserID = c.GetInt(consts.UserID)
	}
	if err := service.FileSvc.UpdateFolder(req); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	} else {
		c.JSON(http.StatusOK, response.Ok(nil))
	}
}
func deleteFolder(c *gin.Context) {
	folderID := c.Param("id")
	if folderID == "" {
		c.JSON(http.StatusBadRequest, response.NewErr(nil))
		return
	}
	if err := service.FileSvc.DeleteFolder(folderID, c.GetInt(consts.UserID)); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, response.Ok(nil))
}

func createFile(c *gin.Context) {
	req := &msg.ReqFileCreate{}
	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	} else {
		req.UserID = c.GetInt(consts.UserID)
	}
	if err := service.FileSvc.CreateFile(req); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, response.Ok(nil))
}

func deleteFile(c *gin.Context) {
	fileID := c.Param("id")
	uid := c.GetInt(consts.UserID)
	if err := service.FileSvc.DeleteFile(fileID, uid); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, response.Ok(nil))
}

func uploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	var req msg.ReqFileUpload

	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	data, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	req.FileHeader = header
	req.Data = data
	req.UserID = c.GetInt(consts.UserID)

	if err := service.FileSvc.UploadFile(&req); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, response.Ok(nil))
}
