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

var (
	buildFolderRsp FuncFolderBuildRsp = func(r *model.Folder) (*msg.RspFolderListItem, error) {
		return &msg.RspFolderListItem{
			FolderID:   r.ID,
			FolderName: r.Name,
			CreatedAt:  r.CreatedAt,
			ModifiedAt: r.UpdatedAt,
		}, nil
	}
)

type FuncFolderBuildRsp func(r *model.Folder) (*msg.RspFolderListItem, error)

func FunctionalFolderMap(folders []*model.Folder, fn interface{}) interface{} {
	var out interface{}
	switch fn.(type) {
	case FuncFolderBuildRsp:
		_temp := make([]*msg.RspFolderListItem, 0)
		for _, e := range folders {
			rsp, err := (fn).(FuncFolderBuildRsp)(e)
			if err != nil {
				return err
			}
			_temp = append(_temp, rsp)
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
