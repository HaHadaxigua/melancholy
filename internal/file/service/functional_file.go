/******
** @date : 2/13/2021 10:51 PM
** @author : zrx
** @description:
******/
package service

import (
	"github.com/HaHadaxigua/melancholy/internal/file/model"
	"github.com/HaHadaxigua/melancholy/internal/file/msg"
)


// FuncFolderBuildRsp 构造文件夹结构的返回体
type FuncFolderBuildRsp func(folder *model.Folder) *msg.RspFolderListItem

// FuncFileBuildRsp 构造文件结构的返回体
type FuncFileBuildRsp func(file *model.File) *msg.RspFileListItem

var (
	buildFolderItemRsp FuncFolderBuildRsp = func(folder *model.Folder) *msg.RspFolderListItem {
		return &msg.RspFolderListItem{
			FolderID:   folder.ID,
			FolderName: folder.Name,
			CreatedAt:  folder.CreatedAt,
			ModifiedAt: folder.UpdatedAt,
		}
	}

	buildFileItemRsp FuncFileBuildRsp = func(file *model.File) *msg.RspFileListItem {
		return &msg.RspFileListItem{
			ID:       file.ID,
			ParentID: file.ParentID,
			Name:     file.Name,
			Suffix:   file.Suffix,
			Hash:     file.Hash,
			Address:  file.Address,
			Size:     file.Size,
			Mode:     file.Mode,
			Modified: file.UpdatedAt,
		}
	}
)

// FunctionalFolderFilter 过滤文件夹
func FunctionalFolderFilter(folders []*model.Folder, fn func(r *model.Folder) bool) []*model.Folder {
	var out []*model.Folder
	for _, e := range folders {
		if fn(e) {
			out = append(out, e)
		}
	}
	return out
}

// FunctionalFileFilter 过滤文件
func FunctionalFileFilter(files []*model.File, fn func(f *model.File) bool) []*model.File {
	var out []*model.File
	for _, e := range files {
		if fn(e) {
			out = append(out, e)
		}
	}
	return out
}

func FunctionalFolder(folders []*model.Folder, fn interface{}) interface{} {
	var out interface{}
	switch fn.(type) {
	case FuncFolderBuildRsp:
		_temp := make([]*msg.RspFolderListItem, 0)
		for _, e := range folders {
			rsp := fn.(FuncFolderBuildRsp)(e)
			_temp = append(_temp, rsp)
		}
		out = _temp
	}
	return out
}

func FunctionalFile(files []*model.File, fn interface{}) interface{} {
	var out interface{}
	switch fn.(type) {
	case FuncFileBuildRsp:
		_temp := make([]*msg.RspFileListItem, len(files))
		for i, e := range files {
			rsp := fn.(FuncFileBuildRsp)(e)
			_temp[i] = rsp
		}
		out = _temp
	}
	return out
}
