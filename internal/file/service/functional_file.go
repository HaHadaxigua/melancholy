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
	// 构造文件夹列表的返回体
	buildFolderItemRsp FuncFolderBuildRsp = func(folder *model.Folder) *msg.RspFolderListItem {
		fileItems := FunctionalFile(folder.Files, buildFileItemRsp).([]*msg.RspFileListItem)
		return &msg.RspFolderListItem{
			FolderID:   folder.ID,
			FolderName: folder.Name,
			FileItems:  fileItems,
			CreatedAt:  folder.CreatedAt,
			UpdatedAt:  folder.UpdatedAt,
		}
	}

	// 构建文件列表的返回体
	buildFileItemRsp FuncFileBuildRsp = func(file *model.File) *msg.RspFileListItem {
		return &msg.RspFileListItem{
			ID:        file.ID,
			ParentID:  file.ParentID,
			Name:      file.Name,
			Suffix:    file.Suffix,
			Hash:      file.Hash,
			Address:   file.Address,
			Size:      file.Size,
			Mode:      file.Mode,
			CreatedAt: file.CreatedAt,
			UpdatedAt: file.UpdatedAt,
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

// FunctionalFolder 用于文件夹的通用方法
func FunctionalFolder(folders []*model.Folder, fn interface{}) interface{} {
	var out interface{}
	switch fn.(type) {
	case FuncFolderBuildRsp:
		items := make([]*msg.RspFolderListItem, len(folders))
		for i, folder := range folders {
			items[i] = fn.(FuncFolderBuildRsp)(folder)
		}
		out = items
	}
	return out
}

// FunctionalFile 用于文件的同用方法
func FunctionalFile(files []*model.File, fn interface{}) interface{} {
	var out interface{}
	switch fn.(type) {
	case FuncFileBuildRsp:
		items := make([]*msg.RspFileListItem, len(files))
		for i, file := range files {
			items[i] = fn.(FuncFileBuildRsp)(file)
		}
		out = items
	}
	return out
}

// ======================自定义结构的方法======================
