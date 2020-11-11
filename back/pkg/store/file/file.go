package file

import (
	"github.com/HaHadaxigua/melancholy/pkg/model/file"
	"github.com/HaHadaxigua/melancholy/pkg/store"
)

// SaveBaseFile 保存文件
func SaveBaseFile(fb *file.Folder) error {
	db := store.GetConn()
	result := db.Create(fb)
	err := result.Error
	if err != nil {
		return err
	}
	return nil
}

//GetChildFolders 获取指定文件夹id下的子文件夹
func GetSubFolders(folderId int) ([]*file.Folder, error) {
	db := store.GetConn()
	folder := &file.Folder{}
	var folders []*file.Folder
	if err := db.Model(folder).Where("parent_id = ? AND deleted_at is not null", folderId).Scan(&folders).Error; err != nil {
		return nil, err
	}
	return folders, nil
}

//FindFolderByNameAndParentId 根据名字和父文件夹查看是否存在相同的文件
func FindFolderByNameAndParentId(name string, parentId int64) (*file.Folder, error) {
	db := store.GetConn()
	fb := &file.Folder{}
	result := db.Model(&fb).Where("parent_id = ? AND name = ? AND deleted_at is not null ", parentId, name).Scan(&fb)
	if result.Error != nil {
		return fb, nil
	} else {
		return nil, result.Error
	}
}

// SaveVideFile 保存视频文件
func SaveVideoFile(vf *file.MeVideoFile) error {
	db := store.GetConn()
	result := db.Create(vf)
	err := result.Error
	if err != nil {
		return err
	}
	return nil
}
