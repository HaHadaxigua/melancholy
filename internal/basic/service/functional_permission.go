/******
** @date : 2/11/2021 2:16 PM
** @author : zrx
** @description:
******/
package service

import (
	"github.com/HaHadaxigua/melancholy/internal/basic/model"
	"github.com/HaHadaxigua/melancholy/internal/basic/msg"
)

var (
	buildPermRsp FuncPermsBuildRsp = func(r *model.Permission) *msg.RspPermListItem {
		return &msg.RspPermListItem{
			PermissionID:   r.ID,
			PermissionName: r.Name,
		}
	}
)

type FuncPermsBuildRsp func(r *model.Permission) *msg.RspPermListItem

func FunctionalPermissionMap(perms []*model.Permission, fn interface{}) interface{} {
	var out interface{}
	switch fn.(type) {
	case FuncPermsBuildRsp:
		_temp := make([]*msg.RspPermListItem, 0)
		for _, e := range perms {
			_temp = append(_temp, (fn).(FuncPermsBuildRsp)(e))
		}
		out = _temp
	}
	return out
}

func FunctionalPermissionFilter(perms []*model.Permission, fn func(r *model.Permission) bool) []*model.Permission {
	var out []*model.Permission
	for _, e := range perms {
		if fn(e) {
			out = append(out, e)
		}
	}
	return out
}
