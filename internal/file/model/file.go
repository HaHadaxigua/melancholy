/******
** @date : 2/13/2021 6:26 PM
** @author : zrx
** @description:
******/
package model

import (
	"github.com/HaHadaxigua/melancholy/internal/file/msg"
	"gorm.io/gorm"
	"time"
)

type Folder struct {
	ID       string `json:"id"`
	OwnerID  int    `json:"ownerID"`  // 拥有者id
	ParentID string `json:"parentID"` // 父文件夹ID
	Name     string `json:"name"`     // 文件夹名

	Files []*File   `json:"files" gorm:"foreignKey:ParentID"` // 一个文件夹会拥有多个子文件
	Subs  []*Folder `json:"subs" gorm:"foreignKey:ParentID"`  // 一个文件夹会有多个子文件夹

	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

func (f Folder) TableName() string {
	return "folders"
}

func (f Folder) ToFileSearchItem() *msg.RspFileSearchItem {
	return &msg.RspFileSearchItem{
		ID:       f.ID,
		Filename: f.Name,
		IsDir:    true,
		// todo, 这里是否需要去找一下id
		Size:      0,
		CreatedAt: f.CreatedAt,
		UpdatedAt: f.UpdatedAt,
	}
}

func (f Folder) ToFolderListItem() *msg.RspFolderListItem {
	var files Files
	files = f.Files
	fileItems := files.ToRspFileList()
	return &msg.RspFolderListItem{
		FolderID:   f.ID,
		FolderName: f.Name,
		FileItems:  fileItems,
		CreatedAt:  f.CreatedAt,
		UpdatedAt:  f.UpdatedAt,
	}
}

type Folders []*Folder

func (folders Folders) GetIDs() []string {
	var ids []string
	for i := 0; i < len(folders); i++ {
		ids = append(ids, folders[i].ID)
	}
	return ids
}

type File struct {
	ID         string `json:"id"`         // 文件ID
	OwnerID    int    `json:"ownerID"`    // 创建者ID
	ParentID   string `json:"parentID"`   // 父文件夹ID
	Name       string `json:"name"`       // 文件名
	Suffix     int    `json:"suffix"`     // 文件后缀
	Hash       string `json:"hash"`       // 文件hash
	Address    string `json:"address"`    // 返回的oss地址
	BucketName string `json:"bucketName"` // oss中的存储桶名字
	ObjectName string `json:"objectName"` // oss中的存储对象名字
	Endpoint   string `json:"endpoint"`
	Size       int    `json:"size"`  // 文件大小
	Mode       int    `json:"mode"`  // 文件模式
	Ftype      int    `json:"ftype"` // 文件类型

	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

func (f File) TableName() string {
	return "files"
}

// ToFileSearchItem 转换结构体
func (f File) ToFileSearchItem() *msg.RspFileSearchItem {
	return &msg.RspFileSearchItem{
		ID:        f.ID,
		Filename:  f.Name,
		IsDir:     false,
		Size:      f.Size,
		Address:   f.Address,
		CreatedAt: f.CreatedAt,
		UpdatedAt: f.UpdatedAt,
	}
}

// ToFindFileByTypeItem 转换结构体
func (f File) ToFindFileByTypeItem() *msg.RspFindFileByTypeItem {
	return &msg.RspFindFileByTypeItem{
		ID:        f.ID,
		Name:      f.Name,
		Address:   f.Address,
		Size:      f.Size,
		CreatedAt: f.CreatedAt,
		UpdatedAt: f.UpdatedAt,
		Hash:      f.Hash,
	}
}

func (f File) ToFileListItem() *msg.RspFileListItem {
	return &msg.RspFileListItem{
		ID:        f.ID,
		ParentID:  f.ParentID,
		Name:      f.Name,
		Suffix:    f.Suffix,
		Hash:      f.Hash,
		Address:   f.Address,
		Size:      f.Size,
		Mode:      f.Mode,
		CreatedAt: f.CreatedAt,
		UpdatedAt: f.UpdatedAt,
	}
}

type Files []*File

func (files Files) GetLen() int {
	return len(files)
}

// GetIDs 获取文件list的idList
func (files Files) GetIDs() []string {
	var ids []string
	for i := 0; i < len(files); i++ {
		ids = append(ids, files[i].ID)
	}
	return ids
}

// ToRspFindFileItemByType 用于构造（根据文件类型查找时的返回值）
func (files Files) ToRspFindFileItemByType() []*msg.RspFindFileByTypeItem {
	var items []*msg.RspFindFileByTypeItem
	for i := 0; i < files.GetLen(); i++ {
		items = append(items, files[i].ToFindFileByTypeItem())
	}
	return items
}

// ToRspFileList 构造（查询文件信息时候的返回体）
func (files Files) ToRspFileList() []*msg.RspFileListItem {
	var items []*msg.RspFileListItem
	for i := 0; i < files.GetLen(); i++ {
		items = append(items, files[i].ToFileListItem())
	}
	return items
}

// DocFile 文本类型文件
type DocFile struct {
	ID      string `json:"id"`      // 文件ID
	Content string `json:"Content"` // 内容
}

func (DocFile) TableName() string {
	return "doc_files"
}

// VideoFile 视频类型文件
type VideoFile struct {
	ID                string   `json:"id"`                // 对应的文件id
	Title             string   `json:"title"`             // 视频标题
	Description       string   `json:"description"`       // 视频描述
	CoverUrl          string   `json:"coverUrl"`          // 视频封面地址
	Area              string   `json:"area"`              // 地区
	Species           string   `json:"species"`           // 视频类型
	ProductionCompany string   `json:"productionCompany"` // 制作公司
	Years             int      `json:"years"`             // 年份
	Duration          int      `json:"duration"`          // 时长
	Finished          bool     `json:"finished"`          // 是否上传完成
	VideoID           string   `json:"videoID"`           // 视频点播的视频id
	Region            string   `json:"region"`            // 存储地区
}

func (VideoFile) TableName() string {
	return "video_files"
}

// MusicFile 视频类型文件
type MusicFile struct {
	ID       string   `json:"id"`       // 对应的文件id
	Name     string   `json:"name"`     // 歌名
	CoverUrl string   `json:"coverUrl"` // 封面地址
	Duration int      `json:"duration"` // 时长
	Singer   string   `json:"singer"`   // 歌手
	Album    string   `json:"album"`    // 专辑
	Years    int      `json:"years"`    // 年份
	Species  string   `json:"species"`  // 类型
	Finished bool     `json:"finished"` // 是否上传完成
	MusicID  string   `json:"musicID"`  // 视频点播的音频id
	Region   string   `json:"region"`   // 存储地区
}

func (MusicFile) TableName() string {
	return "music_files"
}
