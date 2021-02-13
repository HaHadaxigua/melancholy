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

type FuncFolderBuildRsp func(r *model.Folder) *msg.RspFolderListItem

func FunctionalFolderMap(folders []*model.Folder, fn interface{}) interface{} {
	var out interface{}
	switch fn.(type) {
	case FuncFolderBuildRsp:
		_temp := make([]*msg.RspFolderListItem, 0)
		for _, e := range folders {
			_temp = append(_temp, (fn).(FuncFolderBuildRsp)(e))
		}
		out = _temp
	}
	return out
}

func FunctionalFolderFilter(folders []*model.Folder, fn func(r *model.Folder) bool) []*model.Folder {
	var out []*model.Folder
	for _, e := range folders {
		if fn(e) {
			out = append(out, e)
		}
	}
	return out
}
