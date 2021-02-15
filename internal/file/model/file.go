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
	ID      string `json:"id"`
	OwnerID int    `json:"ownerID"` // 拥有者id
	Name    string `json:"name"`

	Files []*File   `json:"files" gorm:"foreignKey:ParentID"` // 一个文件夹会拥有多个子文件
	Subs  []*Folder `json:"subs" gorm:"many2many:folder_sub"` // 一个文件夹会有多个子文件夹

	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

func (f Folder) TableName() string {
	return "folders"
}

type File struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ParentID string `json:"parentFolderID"`

	MD5      string `json:"md5"`
	//FileType int    `json:"fileType"` // 文件类型
	//Address  string `json:"address"`  // oss地址

	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

func (f File) TableName() string {
	return "files"
}
