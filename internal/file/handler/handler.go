package handler

import (
	"errors"
	"fmt"
	"github.com/HaHadaxigua/melancholy/internal/basic/middleware"
	"github.com/HaHadaxigua/melancholy/internal/consts"
	fConst "github.com/HaHadaxigua/melancholy/internal/file/envir"
	"github.com/HaHadaxigua/melancholy/internal/file/msg"
	"github.com/HaHadaxigua/melancholy/internal/file/service"
	"github.com/HaHadaxigua/melancholy/internal/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	folder.GET("/info", getFolderInfo)
	folder.PATCH("/modify", modifyFolder)
	folder.DELETE("/single", deleteFolder) // 删除单个文件
	folder.DELETE("/patch", patchDeleteFolder)
	folder.POST("/include", folderInclude) // 获取给定文件夹下的文件夹和文件

	// file's api
	file := secured.Group("/file")
	file.POST("/search", searchFile)
	file.GET("/list", listFile)
	file.POST("/create", createFile)
	file.DELETE("/single", deleteFile)              // 删除单个文件
	file.DELETE("/patch", patchDeleteFile)          // 批量删除文件
	file.POST("simple/upload", uploadSimpleFile)    // 处理小文件
	file.GET("simple/download", downloadSimpleFile) // 处理小文件
	// 处理分片文件上传
	file.GET("/multi/checkChunk", checkChunk)    // 检查文件的上传情况
	file.POST("/multi/uploadChunk", uploadChunk) // 上传文件分片
	file.POST("/multi/mergeChunk", mergeChunk)   // 合并分片文件

	// 统一处理文件夹和文件的方法
	file.DELETE("/integration", deleteInIntegration) // 通过一个方法来删除文件夹和文件

	// 处理文档型文件相关方法
	file.GET("/findByType", findFileByType) // 获取当前用户的所有图片
	file.POST("/create/doc", createDoc)     // 创建文档类型文件

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

// getFolderInfo 获取文件夹信息
func getFolderInfo(c *gin.Context) {
	var req msg.ReqFolderGetInfo
	if err := c.BindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	req.UserID = c.GetInt(consts.UserID)
	rsp, err := service.FileSvc.FolderGetInfo(&req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, response.NewErr(err))
			return
		}
		c.JSON(http.StatusInternalServerError, response.NewErr)
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

// deleteFolder 递归的删除文件夹
func deleteFolder(c *gin.Context) {
	var req msg.ReqFolderDelete
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	req.UserID = c.GetInt(consts.UserID)
	if err := service.FileSvc.FolderRDelete(&req); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, response.Ok(nil))
}

// patchDeleteFolder 批量删除文件夹
func patchDeleteFolder(c *gin.Context) {
	var req msg.ReqFolderPatchDelete
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	req.UserID = c.GetInt(consts.UserID)
	if err := service.FileSvc.FolderRPatchDelete(&req); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, response.Ok(nil))
}

// folderInclude 获取指定文件夹中包含的内容
func folderInclude(c *gin.Context) {
	var req msg.ReqFolderInclude
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	req.UserID = c.GetInt(consts.UserID)
	if rsp, err := service.FileSvc.FolderInclude(&req); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	} else {
		c.JSON(http.StatusOK, response.Ok(rsp))
	}
}

// searchFile 搜索文件
func searchFile(c *gin.Context) {
	var req msg.ReqFileSearch
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	req.UserID = c.GetInt(consts.UserID)
	rsp, err := service.FileSvc.FileSearch(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, response.Ok(rsp))
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
	rsp, err := service.FileSvc.FileCreate(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, response.Ok(rsp))
}

func deleteFile(c *gin.Context) {
	var req msg.ReqFileDelete
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	req.UserID = c.GetInt(consts.UserID)
	if err := service.FileSvc.FileDelete(&req); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, response.Ok(nil))
}

// patchDeleteFile 批量删除文件
func patchDeleteFile(c *gin.Context) {
	var req msg.ReqFilePatchDelete
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	req.UserID = c.GetInt(consts.UserID)
	if err := service.FileSvc.FilePatchDelete(&req); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, response.Ok(nil))
}

// uploadSimpleFile 上传简单文件
func uploadSimpleFile(c *gin.Context) {
	// 这种方式将文件读入了内存，可能会导致内存爆掉
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
	c.JSON(http.StatusOK, response.OK)
}

// downloadSimpleFile 处理简单文件的下载
func downloadSimpleFile(c *gin.Context) {
	var req msg.ReqFileDownload
	if err := c.BindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	req.UserID = c.GetInt(consts.UserID)
	rsp, err := service.FileSvc.FileSimpleDownload(&req)
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

// checkChunk 检查文件的分片情况
func checkChunk(c *gin.Context) {
	var req msg.ReqFileMultiCheck
	if err := c.BindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	req.UserID = c.GetInt(consts.UserID)
	rsp, err := service.FileSvc.FileMultiCheck(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, response.Ok(rsp))
}

// uploadChunk 上传文件分片,返回已经完成的文件分片列表
func uploadChunk(c *gin.Context) {
	var req msg.ReqFileMultiUpload
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	fileHeader, err := c.FormFile(fConst.FileUpload)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	req.UserID = c.GetInt(consts.UserID)
	req.FileHeader = fileHeader
	rsp, err := service.FileSvc.FileMultiUpload(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, response.Ok(rsp))
}

// mergeChunk 合并文件分片，意味着文件上传完成，需要将文件上传到oss中
func mergeChunk(c *gin.Context) {
	var req msg.ReqFileMultiMerge
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	req.UserID = c.GetInt(consts.UserID)
	rsp, err := service.FileSvc.FileMultiMerge(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
	// 用来判断子协程是否处理完成
	<-rsp.Done
	if rsp.Result != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(rsp.Result))
		return
	}
	c.JSON(http.StatusOK, response.OK)
}

// deleteInIntegration 通过一个方法来同时删除文件夹和文件
func deleteInIntegration(c *gin.Context) {
	var req msg.ReqDeleteInIntegration
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	req.UserID = c.GetInt(consts.UserID)
	if err := service.FileSvc.DeleteInIntegration(&req); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, response.OK)
}

// findFileByType 根据文件类型查找文件
func findFileByType(c *gin.Context) {
	var req msg.ReqFindFileByType
	if err := c.BindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	if req.FileType < 0 || req.Offset < -1 || req.Limit < -1 {
		c.JSON(http.StatusBadRequest, response.UnKnown)
		return
	}
	req.UserID = c.GetInt(consts.UserID)
	rsp, err := service.FileSvc.FindFileByType(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, response.Ok(rsp))
}

// createDoc 创建文档类型文件
func createDoc(c *gin.Context) {
	var req msg.ReqDocFile
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	req.UserID = c.GetInt(consts.UserID)
	rsp, err := service.FileSvc.CreateDoc(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, response.Ok(rsp))
}
