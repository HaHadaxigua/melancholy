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

type ReqFileCreate struct {
	FileName string `json:"fileName"`
	ParentID string `json:"parentID"`

	UserID int
}

type ReqFileUpload struct {
	Data       []byte                `json:"data" `
	FileHeader *multipart.FileHeader `json:"fileHeader"`
	ParentID   string                `form:"parentID" json:"parentID"`

	UserID int
}
