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
	Size       int    `json:"size"`       // 文件大小
	Mode       int    `json:"mode"`       // 文件模式

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
		CreatedAt: f.CreatedAt,
		UpdatedAt: f.UpdatedAt,
	}
}
