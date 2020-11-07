package file

import (
	"fmt"
	"github.com/HaHadaxigua/melancholy/pkg/model"
)

// 云文件
type CloudFile interface {
	Download() (*BaseFile, error)
	Upload() error
	FileName() string
	FileSize() string
	FilePath() string
}

// BaseFile 基础文件抽象, 所有的文件都基于此
type BaseFile struct {
	model.Model
	Creator     int64  `json:"creator"`
	Url         string `json:"url"`  // 逻辑路径
	Path        string `json:"path"` // 底层文件路径
	Name        string
	Md5         string
	Size        int64 `json:"size"`        // 描述文件大小， 如果是文件夹， 则描述文件夹内容的大小
	DFlag       bool  `json:"dFlag"`       // 是否是文件夹
	ParentDirID int64 `json:"parentDirId"` // 父文件夹id
	ChildFileID int64 `json:"childFileId"` // 子文件id
}

func(b *BaseFile) String() string{
	return fmt.Sprintf("")
}

// NewBaseFile 创建文件
func NewBaseFile(creator, parentDirId int64, name, url string) (*BaseFile, error) {
	// 需要生成一个md5
	return &BaseFile{
		Creator:     creator,
		Name:        name,
		Url:         url,
		ParentDirID: parentDirId,
	}, nil
}

// NewBaseFileDir 创建文件夹
func NewFolder(creator, parentDirId int64, name, md5 string) *BaseFile {
	return &BaseFile{
		Creator:     creator,
		Name:        name,
		Md5:         md5,
		ParentDirID: parentDirId,
		DFlag:       true,
	}
}

//MeFile 具体文件抽象， 通过MeType 来分辨具体的类型
type MeVideoFile struct {
	BaseFile
	MeType  string `json:"meType"`  // 视频类型的分类
	Section string `json:"section"` // 视频标签
}
