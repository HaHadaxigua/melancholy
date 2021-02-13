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
	buildRoleRsp FuncRoleBuildRsp = func(r *model.Role) *msg.RspRoleListItem {
		return &msg.RspRoleListItem{
			RoleID:      r.ID,
			RoleName:    r.Name,
			Permissions: r.Permissions,
		}
	}
)

type FuncRoleBuildRsp func(r *model.Role) *msg.RspRoleListItem


func FunctionalRoleMap(roles []*model.Role, fn interface{}) interface{} {
	var out interface{}
	switch fn.(type) {
	case FuncRoleBuildRsp:
		_temp := make([]*msg.RspRoleListItem, 0)
		for _, e := range roles {
			_temp = append(_temp, (fn).(FuncRoleBuildRsp)(e))
		}
		out = _temp
	}
	return out
}

func FunctionalRoleFilter(roles []*model.Role, fn func(r *model.Role) bool) []*model.Role {
	var out []*model.Role
	for _, e := range roles {
		if fn(e) {
			out = append(out, e)
		}
	}
	return out
}
