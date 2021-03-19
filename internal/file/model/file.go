/******
** @date : 2/13/2021 6:26 PM
** @author : zrx
** @description:
******/
package model

import (
	"gorm.io/gorm"
	"time"
)

type Folder struct {
	ID       string `json:"id"`
	OwnerID  int    `json:"ownerID"`  // 拥有者id
	ParentID string `json:"parentID"` // 父文件夹ID
	Name     string `json:"name"`

	Files []*File   `json:"files" gorm:"foreignKey:ParentID"` // 一个文件夹会拥有多个子文件
	Subs  []*Folder `json:"subs" gorm:"foreignKey:ParentID"`  // 一个文件夹会有多个子文件夹

	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

func (f Folder) TableName() string {
	return "folders"
}

type File struct {
	ID       string `json:"id"`
	OwnerID  int `json:"ownerID"`  // 创建者ID
	ParentID string `json:"parentID"` // 父文件夹ID
	Name     string `json:"name"`     // 文件名
	Suffix   int    `json:"suffix"`   // 文件后缀
	Hash     string `json:"hash"`     // 文件hash
	Address  string `json:"address"`  // 返回的oss地址
	Size     int    `json:"size"`     // 文件大小
	Mode     int    `json:"mode"`     // 文件模式

	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

func (f File) TableName() string {
	return "files"
}
