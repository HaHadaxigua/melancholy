/******
** @date : 2/4/2021 12:41 AM
** @author : zrx
** @description:
******/
package msg

import "time"

// RspFolderListItem 文件列表的item
type RspFolderListItem struct {
	FolderID   string             `json:"folderID"`
	FolderName string             `json:"folderName"`
	FileItems  []*RspFileListItem `json:"fileItems"`

	CreatedAt  time.Time `json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
}

// RspFolderList 文件列表返回体
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

// RspFileDownload 简单文件下载的返回
type RspFileDownload struct {
	Content  []byte `json:"content"`
	FileName string `json:"fileName"`
}

// RspFileMultiCheck 文件分片列表的完成情况的返回
type RspFileMultiCheck struct {
	ChunkList []string `json:"chunkList"` // 文件分片编号列表
	State     int      `json:"state"`     // 文件完成情况  state为0时 说明文件并不完整，为1 说明文件完整
}

type RspFileMultiUpload struct {
	ChunkList []string `json:"chunkList"` // 文件分片编号列表
	State     int      `json:"state"`     // 当前文件的上传完成情况 1:说明文件全部上传完成
}

type RspFileMultiMerge struct {
	Result error `json:"result"`
	Done   chan struct{}
}
