/******
** @date : 2/4/2021 12:41 AM
** @author : zrx
** @description:
******/
package msg

import "time"

type RspFolderListItem struct {
	FolderID   string `json:"folderID"`
	FolderName string `json:"folderName"`

	CreatedAt  time.Time `json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
}

type RspFolderList struct {
	List []*RspFolderListItem `json:"list"`
}
