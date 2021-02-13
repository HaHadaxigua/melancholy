/******
** @date : 2/4/2021 12:41 AM
** @author : zrx
** @description:
******/
package msg

import "time"

type RspFolderListItem struct {
	FileID   int    `json:"fileID"`
	Filename string `json:"filename"`

	FolderID   int    `json:"folderID"`
	FolderName string `json:"folderName"`

	OwnerName string `json:"ownerName"`
	OwnerID   int    `json:"ownerID"`

	CreatedAt time.Time `json:"createdAt"`
}

type RspFileList struct {
	list  []*RspFolderListItem `json:"list"`
	Total int                `json:"total"`
}
