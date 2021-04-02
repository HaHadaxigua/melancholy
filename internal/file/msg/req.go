/******
** @date : 2/4/2021 12:40 AM
** @author : zrx
** @description:
******/
package msg

import (
	"github.com/HaHadaxigua/melancholy/internal/file/consts"
	"mime/multipart"
	"path"
	"strings"
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

type ReqFileCreate struct {
	ParentID string `json:"parentID"`
	FileName string `json:"fileName"`
	FileType int    `json:"fileType"` // 创建的文件类型

	UserID int
}

// 验证请求是否合法
func (req ReqFileCreate) Verify() bool {
	if req.FileType == 0 {
		// 判断上传的文件id是否支持
		if ftid, ok := consts.MapFileTypeToID[path.Ext(req.FileName)]; !ok {
			return false
		} else {
			req.FileType = ftid
		}
	} else {
		if ftstr, ok := consts.MapFileTypeToStr[req.FileType]; !ok {
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
