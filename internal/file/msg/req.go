/******
** @date : 2/4/2021 12:40 AM
** @author : zrx
** @description:
******/
package msg

type ReqFolderCreate struct {
	FolderName string `json:"filename"`
	ParentID   int    `json:"parentID"`

	UserID int
}

type ReqFileListChildFilter struct {
}
