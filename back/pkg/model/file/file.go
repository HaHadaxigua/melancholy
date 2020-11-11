package file

import (
	"fmt"
	"github.com/HaHadaxigua/melancholy/pkg/model"
	"github.com/HaHadaxigua/melancholy/pkg/tools"
	"time"
)

// 云文件
type CloudFile interface {
	Download() (*Folder, error)
	Upload() error
	FileName() string
	FileSize() string
	FilePath() string
}

// BaseFile 基础文件抽象, 所有的文件都基于此
type Folder struct {
	model.Model
	Creator  int64  `json:"creator"`
	Url      string `json:"url"`      // 逻辑路径
	Path     string `json:"path"`     // 底层文件路径
	Name     string `json:"name"`     // 文件名
	Size     int64  `json:"size"`     // 描述文件夹大小， 如果是文件夹， 则描述文件夹内容的大小
	ParentID int64  `json:"parentId"` // 父文件夹id
}

func (f *Folder) TableName() string {
	return "folder"
}

func (f *Folder) String() string {
	return fmt.Sprintf("folder name:%s, url:%s, createdAt:%s", f.Name, f.Url, f.CreatedAt)
}

// NewBaseFile 创建文件
func NewBaseFile(creator, parentDirId int64, name, url string) (*Folder, error) {
	// 需要生成一个md5
	return &Folder{
		Creator:  creator,
		Name:     name,
		Url:      url,
		ParentID: parentDirId,
	}, nil
}

// NewFolder 创建文件夹
func NewFolder(creator, parentDirId int64, name string) *Folder {
	return &Folder{
		Creator:  creator,
		Name:     name,
		Path:     GenFilePath(),
		Url:      GenFileUrl(),
		ParentID: parentDirId,
	}
}

//MeFile 具体文件抽象， 通过MeType 来分辨具体的类型
type MeVideoFile struct {
	MeType  string `json:"meType"`  // 视频类型的分类
	Section string `json:"section"` // 视频标签
}

//TODO: 返回一个path
func GenFilePath() string {
	return tools.MD5(time.Now().String())
}

//TODO: 生成文件路由
func GenFileUrl() string {
	return tools.MD5(time.Now().String())
}
