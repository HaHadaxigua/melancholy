/******
** @date : 2/4/2021 12:40 AM
** @author : zrx
** @description:
******/
package msg

import "mime/multipart"

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
	UserID   int
}

type ReqFileUpload struct {
	Data       []byte                `json:"data"`
	FileHeader *multipart.FileHeader `json:"fileHeader"`
	ParentID   string                `form:"parentID" json:"parentID"`

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
