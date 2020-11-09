package file

import (
	"fmt"
	"github.com/HaHadaxigua/melancholy/pkg/model/file"
	"github.com/HaHadaxigua/melancholy/pkg/store"
)

// SaveBaseFile 保存文件
func SaveBaseFile(fb *file.BaseFile) error {
	db := store.GetConn()
	result := db.Create(fb)
	err := result.Error
	if err != nil {
		return err
	}
	return nil
}

// SaveVideFile 保存视频文件
func SaveVideoFile(vf *file.MeVideoFile) error {
	db := store.GetConn()

	result := db.Create(vf)
	fmt.Println(result.Error)
	fmt.Println(db)
	err := result.Error
	if err != nil {
		return err
	}
	return nil
}

//IsDeleted 判断是否已经删除
func IsDeleted(bf *file.BaseFile) (bool, error) {
	return false, nil
}

//FindFolderByNameAndParentId 根据名字和父文件夹查看是否存在相同的文件
func FindFolderByNameAndParentId(name string, parentId int64) (*file.BaseFile, error) {
	db := store.GetConn()
	fb := &file.BaseFile{}
	result := db.Model(&fb).Where("parent_id = ? AND name = ?", parentId, name).Scan(&fb)
	if result.Error != nil {
		return fb, nil
	}else {
		return nil, result.Error
	}
}
