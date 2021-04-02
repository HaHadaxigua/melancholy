/******
** @date : 2/4/2021 12:41 AM
** @author : zrx
** @description:
******/
package msg

import "time"

type RspFolderListItem struct {
	FolderID   string             `json:"folderID"`
	FolderName string             `json:"folderName"`
	FileItems  []*RspFileListItem `json:"fileItems"`

	CreatedAt  time.Time `json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
}

type RspFolderList struct {
	FolderItems []*RspFolderListItem `json:"subFolders"`
	FileItems   []*RspFileListItem   `json:"files"`
}

// RspFileListItem 文件列表返回体
type RspFileListItem struct {
	ID       string    `json:"id"`       // 文件ID
	ParentID string    `json:"parentID"` // 父文件ID
	Name     string    `json:"name"`     // 文件名
	Suffix   int       `json:"suffix"`   // 文件后缀
	Hash     string    `json:"hash"`     // 文件hash
	Address  string    `json:"address"`  // 云存储地址
	Size     int       `json:"size"`     // 文件大小
	Mode     int       `json:"mode"`     // 是否只读
	Modified time.Time `json:"modified"` // 上次修改时间
}

// RspFileList 文件列表的返回
type RspFileList struct {
	List  []*RspFileListItem `json:"list"`
	Total int                `json:"total"`
}

type RspFileDownload struct {
	Content  []byte `json:"content"`
	FileName string `json:"fileName"`
}
