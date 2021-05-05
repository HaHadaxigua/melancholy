/******
** @date : 2/4/2021 12:40 AM
** @author : zrx
** @description:
******/
package msg

import (
	"github.com/HaHadaxigua/melancholy/internal/file/envir"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"path"
	"strings"
	"time"
)

type ReqFolderGetInfo struct {
	FolderID string `form:"folderID" json:"folderID" binding:"required"`
	UserID   int
}

type ReqFolderCreate struct {
	FolderName string `json:"filename"`
	ParentID   string `json:"parentID, omitempty"`
	
	ID string `json:"id"`
	UserID int
}

type ReqFolderUpdate struct {
	FolderID string `json:"folderID" binding:"required"`
	NewName  string `json:"newName"`

	UserID int
}

type ReqFolderDelete struct {
	FolderID string `json:"folderID" binding:"required"`

	UserID int
}

type ReqFolderPatchDelete struct {
	FolderIDs []string `json:"folderIDs" binding:"required"`

	UserID int
}

// ReqFolderInclude 当前文件夹下包含的内容 todo: 整理请求体
type ReqFolderInclude struct {
	FolderID   string `json:"folderID" binding:"required"`
	SortedBy   string `json:"sortedBy"`
	Descending bool   `json:"descending"` // 是否是降序

	UserID int
}

type ReqFileCreate struct {
	ParentID string `json:"parentID"`
	FileName string `json:"fileName"`
	FileType int    `json:"fileType"` // 创建的文件类型
	Size     int    `json:"size"`     // 文件大小
	Address  string `json:"address"`  // 文件的oss地址

	UserID int
}

//  Verify 验证请求是否合法
func (req ReqFileCreate) Verify() bool {
	if req.FileType == 0 {
		// 判断上传的文件id是否支持
		if ftid, ok := envir.MapFileTypeToID[path.Ext(req.FileName)]; !ok {
			return false
		} else {
			req.FileType = ftid
		}
	} else {
		if ftstr, ok := envir.MapFileTypeToStr[req.FileType]; !ok {
			return false
		} else {
			if !strings.HasSuffix(req.FileName, ftstr) {
				return false
			}
		}
	}
	return true
}

type ReqFileDelete struct {
	FileID string `json:"fileID" binding:"required"`
	UserID int
}

type ReqFilePatchDelete struct {
	FileIDs []string `json:"fileIDs" binding:"required"`

	UserID int
}

// ReqFileUpload 简单文件的上传请求
type ReqFileUpload struct {
	Data       []byte                `json:"data"`
	FileHeader *multipart.FileHeader `json:"fileHeader"`
	ParentID   string                `form:"parentID" json:"parentID"`
	FileType   int                   `json:"fileType"`

	UserID int
}

// ReqFileSearch 文件搜索
type ReqFileSearch struct {
	Fuzzy string     `json:"fuzzy"`                          // 通过name来搜索文件或者是文件夹
	Start *time.Time `json:"start" time_format:"2006-01-02"` // 文件的最早更新时间
	End   *time.Time `json:"end" time_format:"2006-01-02"`   // 文件的最后更新时间

	Offset int `json:"offset"`
	Limit  int `json:"limit"`
	UserID int
}

// ReqFolderListFilter 列出文件夹
type ReqFolderListFilter struct {
	FolderID string `json:"folderID"`

	UserID int
}

// ReqFileListFilter 列出指定文件夹中的文件
type ReqFileListFilter struct {
	FolderID string `form:"folderID" json:"folderID"`

	UserID int
}

// ReqFileDownload 下载文件的请求
type ReqFileDownload struct {
	FileID string `form:"fileID" json:"fileID"`

	UserID int
}

// ReqDeleteInIntegration 文件夹和文件的整合删除方法
type ReqDeleteInIntegration struct {
	FolderIDs []string `json:"folderIDs"`
	FileIDs   []string `json:"fileIDs"`

	UserID int
}

//=====================分片上传方法====================================

// ReqFileMultiCheck 检查文件上传情况
type ReqFileMultiCheck struct {
	Hash     string `form:"hash" json:"hash"` // 要上传的文件hash
	Filename string `form:"filename" json:"filename"`

	UserID int
}

// ReqFileMultiUpload 上传文件分片
type ReqFileMultiUpload struct {
	Filename  string `json:"filename"`  // 文件名称
	Hash      string `json:"hash"`      // 文件hash， 根据文件hash找到文件
	ChunkID   string `json:"chunkID"`   // 文件的分片id
	ChunkHash string `json:"ChunkHash"` // 分片的hash
	Total     int    `json:"total"`     // 总共有多少个分片

	C          *gin.Context          // gin的上下文
	FileHeader *multipart.FileHeader // describes a file part of a multipart request.
	UserID     int
}

// ReqFileMultiMerge 请求将分片文件合并
type ReqFileMultiMerge struct {
	Filename string `json:"filename"`
	Hash     string `json:"hash"`

	UserID int
}

// ReqFileMultiDownload 文件的分片下载
type ReqFileMultiDownload struct {
	Hash string `json:"hash"`

	UserID int
}

// ReqFindFileByType 找出图片
type ReqFindFileByType struct {
	FileType int `form:"fileType" json:"fileType" binding:"required"`
	Offset   int `form:"offset" json:"offset"`
	Limit    int `form:"limit" json:"limit"`

	UserID int
}

// ReqDocCreate 创建文本文件的请求
type ReqDocCreate struct {
	Name    string `json:"name"`    // 文件名，需要后缀
	Content string `json:"content"` // 文本内容

	UserID int
}
