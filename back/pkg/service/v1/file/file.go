package file

import (
	"fmt"
	"github.com/HaHadaxigua/melancholy/pkg/model/file"
	"github.com/HaHadaxigua/melancholy/pkg/msg"
	sf "github.com/HaHadaxigua/melancholy/pkg/store/file"
	"github.com/HaHadaxigua/melancholy/pkg/tools"
	"time"
)

//CreateFolder 根据请求创建文件夹
func CreateFolder(r *msg.DirRequest) (*file.BaseFile, error) {
	if !VerifyReq(r) {
		return nil, msg.ErrReq
	}
	md5 := tools.MD5(fmt.Sprintf("%d%s%d%s", r.Creator, r.Name, r.ParentId, time.Now().String()))
	tFolder := file.NewFolder(r.Creator, r.ParentId, r.Name, md5)

	if HaveRepeatNameFolder(r.Name, r.ParentId) {
		return nil, msg.FileRepeatErr
	}

	err := sf.SaveBaseFile(tFolder)
	if err != nil {
		return nil, msg.FileSaveErr
	}
	return tFolder, nil
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
	bf, err := sf.FindFolderByNameAndParentId(name, parentId)
	if err != nil {
		return false
	}
	if bf != nil {
		return true
	}
	return false
}
