/******
** @date : 2/4/2021 12:40 AM
** @author : zrx
** @description:
******/
package msg

type ReqFolderCreate struct {
	Filename string `json:"filename"`
	ParentID int    `json:"parentID"`
	UserID   int    `json:"-"`
}

type ReqFileListChildFilter struct {
}
