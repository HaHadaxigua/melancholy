package file

import (
	"github.com/HaHadaxigua/melancholy/pkg/model/file"
	"github.com/HaHadaxigua/melancholy/pkg/store"
)

// SaveBaseFile 保存文件
func CreateBaseFile(fb *file.Folder) error {
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
	if err := db.Model(folder).Where("parent_id = ? AND deleted_at is null", folderId).Scan(&folders).Error; err != nil {
		return nil, err
	}
	return folders, nil
}

//FindFolderByNameAndParentId 判断当前父文件夹下是否有同名文件夹
func FindFolderByNameAndParentId(name string, parentId int64) (int64, error) {
	db := store.GetConn()
	fb := &file.Folder{}
	rows := db.Where("parent_id = ? AND name = ? AND deleted_at is null ", parentId, name).Find(&fb)
	if rows.Error != nil {
		return 0, rows.Error
	} else {
		return rows.RowsAffected, nil
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
