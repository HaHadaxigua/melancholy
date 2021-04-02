package handler

import (
	"fmt"
	"github.com/HaHadaxigua/melancholy/internal/basic/middleware"
	"github.com/HaHadaxigua/melancholy/internal/consts"
	fConst "github.com/HaHadaxigua/melancholy/internal/file/consts"
	"github.com/HaHadaxigua/melancholy/internal/file/msg"
	"github.com/HaHadaxigua/melancholy/internal/file/service"
	"github.com/HaHadaxigua/melancholy/internal/response"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func SetupFileRouters(r gin.IRouter) {
	// open

	// secured
	secured := r.Group("/f", middleware.Auth)
	// folder's api
	folder := secured.Group("/folder")
	folder.POST("/create", createFolder)
	folder.GET("/space", userSpace)
	folder.PATCH("/modify", modifyFolder)
	folder.DELETE("/:id", deleteFolder)

	// file's api
	file := secured.Group("/file")
	file.GET("/list", listFile)
	file.POST("/create", createFile)
	file.DELETE("/:id", deleteFile)
	file.POST("/upload", uploadFile)
	file.GET("/download", downloadFile)
}

func createFolder(c *gin.Context) {
	var req msg.ReqFolderCreate
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	req.UserID = c.GetInt(consts.UserID)

	if err := service.FileSvc.FolderCreate(&req); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, response.Ok(nil))
}

// userSpace list user's root path file
func userSpace(c *gin.Context) {
	uid := c.GetInt(consts.UserID)
	rsp, err := service.FileSvc.UserRoot(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, response.Ok(rsp))
}

// modifyFolder update folder info
func modifyFolder(c *gin.Context) {
	var req msg.ReqFolderUpdate
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	req.UserID = c.GetInt(consts.UserID)
	if err := service.FileSvc.FolderUpload(&req); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, response.Ok(nil))
}

func deleteFolder(c *gin.Context) {
	folderID := c.Param("id")
	if folderID == "" {
		c.JSON(http.StatusBadRequest, response.NewErr(nil))
		return
	}
	req := &msg.ReqFolderDelete{
		FolderID: folderID, UserID: c.GetInt(consts.UserID),
	}
	if err := service.FileSvc.FolderDelete(req); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, response.Ok(nil))
}

// listFile 列出文件
func listFile(c *gin.Context) {
	var req msg.ReqFileListFilter
	if err := c.BindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	req.UserID = c.GetInt(consts.UserID)
	rsp, err := service.FileSvc.FileList(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, response.Ok(rsp))
}

func createFile(c *gin.Context) {
	var req msg.ReqFileCreate
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	req.UserID = c.GetInt(consts.UserID)
	if err := service.FileSvc.FileCreate(&req); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, response.Ok(nil))
}

func deleteFile(c *gin.Context) {
	fileID := c.Param("id")
	uid := c.GetInt(consts.UserID)
	if err := service.FileSvc.FileDelete(fileID, uid); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, response.Ok(nil))
}

func uploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile(fConst.FileUpload)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	var req msg.ReqFileUpload

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	// fixme: 处理文件上传 需要流进行流式上传
	data, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	req.FileHeader = header
	req.Data = data
	req.UserID = c.GetInt(consts.UserID)

	if err := service.FileSvc.FileUpload(&req); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, response.Ok(nil))
}

func downloadFile(c *gin.Context) {
	var req msg.ReqFileDownload
	if err := c.BindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	req.UserID = c.GetInt(consts.UserID)
	rsp, err := service.FileSvc.FileDownload(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
	c.Header(fConst.ContentDisposition, fmt.Sprintf("attachment; filename=%s", rsp.FileName))
	c.Header(fConst.ContentType, "application/text/plain")
	c.Header(fConst.AcceptLength, fmt.Sprintf("%d", len(rsp.Content)))
	c.Writer.Write(rsp.Content)
}
