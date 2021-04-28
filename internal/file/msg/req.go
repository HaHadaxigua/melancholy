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

type ReqFolderCreate struct {
	FolderName string `json:"filename"`
	ParentID   string `json:"parentID, omitempty"`

	UserID int
}

type ReqFolderUpdate struct {
	FolderID string `json:"folderID"`
	NewName  string `json:"newName"`

	UserID int
}

type ReqFolderDelete struct {
	FolderID string `json:"folderID"`

	UserID int
}

// ReqFolderInclude 当前文件夹下包含的内容 todo: 整理请求题
type ReqFolderInclude struct {
	FolderID string `json:"folderID"`

	UserID int
}

type ReqFileCreate struct {
	ParentID string `json:"parentID"`
	FileName string `json:"fileName"`
	FileType int    `json:"fileType"` // 创建的文件类型

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
