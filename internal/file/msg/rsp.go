/******
** @date : 2/4/2021 12:41 AM
** @author : zrx
** @description:
******/
package msg

import "time"

// RspFolderListItem 文件列表的item
type RspFolderListItem struct {
	FolderID   string             `json:"id"`
	FolderName string             `json:"name"`
	FileItems  []*RspFileListItem `json:"fileItems"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// RspFolderList 文件列表返回体
type RspFolderList struct {
	FolderItems []*RspFolderListItem `json:"subFolders"`
	FolderTotal int                  `json:"folderTotal"`
	FileItems   []*RspFileListItem   `json:"files"`
	FileTotal   int                  `json:"fileTotal"`
	Total       int                  `json:"total"`
}

// RspFileListItem 文件列表返回体
type RspFileListItem struct {
	ID        string    `json:"id"`        // 文件ID
	ParentID  string    `json:"parentID"`  // 父文件ID
	Name      string    `json:"name"`      // 文件名
	Suffix    int       `json:"suffix"`    // 文件后缀
	Hash      string    `json:"hash"`      // 文件hash
	Address   string    `json:"address"`   // 云存储地址
	Size      int       `json:"size"`      // 文件大小
	Mode      int       `json:"mode"`      // 是否只读
	CreatedAt time.Time `json:"createdAt"` // 创建时间
	UpdatedAt time.Time `json:"updatedAt"` // 上次修改时间
}

// RspFileList 文件列表的返回
type RspFileList struct {
	List  []*RspFileListItem `json:"list"`
	Total int                `json:"total"`
}

// RspFileSearchResult 文件搜索的返回
type RspFileSearchResult struct {
	List     []*RspFileSearchItem `json:"list"`
	ParentID string               `json:"parentID"`
	Total    int                  `json:"total"`
}

func (r RspFileSearchResult) Len() int {
	return len(r.List)
}

func (r RspFileSearchResult) Less(i, j int) bool {
	return RspFileSearchResultSort(&r, i, j)
}

func (r RspFileSearchResult) Swap(i, j int) {
	r.List[i], r.List[j] = r.List[j], r.List[i]
}

// RspFileSearchItem 文件搜索的返回item
type RspFileSearchItem struct {
	ID       string `json:"id"`      // 文件id
	Filename string `json:"name"`    // 文件名
	IsDir    bool   `json:"isDir"`   // 是否是文件夹
	Size     int    `json:"size"`    // 文件大小
	Address  string `json:"address"` // oss地址

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// RspFileDownload 简单文件下载的返回
type RspFileDownload struct {
	Content  []byte `json:"content"`
	FileName string `json:"name"`
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

type RspFIleMultiDownload struct {
}

// RspFindFileByType 根据文件类型查找文件
type RspFindFileByType struct {
	List  []*RspFindFileByTypeItem `json:"list"`
	Total int                      `json:"total"`
}

type RspFindFileByTypeItem struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Size      int       `json:"size"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Hash      string    `json:"hash"`
}

// RspCreateDocFie 创建文档文件的返回体
type RspDocFile struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Size     int       `json:"size"`
	Address  string    `json:"address"`
	Content  string    `json:"content"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
}

// RspVideoFile 关于视频文件的返回
type RspVideoFile struct {

}

// RspMusicFile 关于音频文件的返回
type RspMusicFile struct {

}
