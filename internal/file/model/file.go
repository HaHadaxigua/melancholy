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
	ID      int    `json:"id"`
	Name    string `json:"name"`
	OwnerID int    `json:"ownerID"` // 拥有者id

	Files   []*File   `json:"files" gorm:"foreignKey:ParentID"`    // 一个文件夹会拥有多个子文件
	Folders []*Folder `json:"folders" gorm:"many2many:folder_sub"` // 一个文件夹会有多个子文件夹

	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

func (f Folder) TableName() string {
	return "folders"
}

type File struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	FolderID int    `json:"folderID"`

	MD5 string `json:"md5"`

	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

func (f File) TableName() string {
	return "files"
}
