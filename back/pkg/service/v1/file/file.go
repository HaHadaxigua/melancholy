package file

import (
	model "github.com/HaHadaxigua/melancholy/pkg/model/file"
	"github.com/HaHadaxigua/melancholy/pkg/msg"
	sf "github.com/HaHadaxigua/melancholy/pkg/store/file"
)

//CreateFolder 根据请求创建文件夹
func CreateFolder(r *msg.DirRequest) (*model.Folder, error) {
	if !VerifyReq(r) {
		return nil, msg.ErrReq
	}
	tFolder := model.NewFolder(r.Creator, r.ParentId, r.Name)

	//判断是否有重复的文件夹
	if HaveRepeatNameFolder(r.Name, r.ParentId) {
		return nil, msg.FileRepeatErr
	}

	err := sf.SaveBaseFile(tFolder)
	if err != nil {
		return nil, msg.FileSaveErr
	}
	return tFolder, nil
}

// GetChildFolders 根据当前路径的文件夹id 获取api
func GetChildFolders(folderId int) ([]*model.Folder, error) {
	if folderId < 0 {
		return nil, msg.ErrReq
	}
	subFolders, err := sf.GetSubFolders(folderId)
	if err != nil {
		return nil, msg.FileNotFoundErr
	}
	return subFolders, nil
}

// VerifyReq 验证请求合法性
func VerifyReq(r *msg.DirRequest) bool {
	if r.Creator <= 0 || r.Name == "" || r.Name == " " || r.ParentId < 0 {
		return false
	}
	return true
}

//HaveRepeatNameFolder 判断是否有相同的文件名在同一个文件夹下
func HaveRepeatNameFolder(name string, parentId int64) bool {
	rows, err := sf.FindFolderByNameAndParentId(name, parentId)
	if err != nil {
		return false
	}
	if rows > 0 {
		return true
	}
	return false
}
