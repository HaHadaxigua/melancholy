/******
** @date : 2/4/2021 12:40 AM
** @author : zrx
** @description:
******/
package msg

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
