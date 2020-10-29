package file

import (
	"github.com/HaHadaxigua/melancholy/pkg/model"
	"github.com/HaHadaxigua/melancholy/pkg/tools"
)

// 云文件
type CloudFile interface {
	Download() (*BaseFile, error)
	Upload() error
	FileName() string
	FileSize() string
	FilePath() string
}

// 基础文件描述, 所有的文件都基于此
type BaseFile struct {
	Name     string
	Path     string
	FileType uint32 `json:"file_type"`
	Size     int64
	Md5      string
	model.Model
}

func (f *BaseFile) FileName() string {
	return f.Name
}

func (f *BaseFile) FileSize() string {
	return tools.FormatBytes(f.Size)
}

func (f *BaseFile) FilePath() string {
	return ""
}

func (f *BaseFile) Download() (*BaseFile, error) {
	return nil, nil
}

func (f *BaseFile) Upload() error {
	return nil
}
