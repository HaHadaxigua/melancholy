/******
** @date : 2/11/2021 2:17 PM
** @author : zrx
** @description:
******/
package service

import (
	"github.com/HaHadaxigua/melancholy/internal/basic/model"
	"github.com/HaHadaxigua/melancholy/internal/basic/msg"
)

var (
	buildUserRsp FuncUserBuildRsp = func(u *model.User) *msg.RspUserListItem {
		return &msg.RspUserListItem{
			UserID:     u.ID,
			UserName:   u.Username,
			UserEmail:  u.Email,
			UserMobile: u.Mobile,
			Roles:      (FunctionalRoleMap(u.Roles, buildRoleRsp)).([]*msg.RspRoleListItem),
		}
	}
)

type FuncUserBuildRsp func(u *model.User) *msg.RspUserListItem

func FunctionalUserMap(users []*model.User, fn interface{}) interface{} {
	var out interface{}
	switch fn.(type) {
	case FuncUserBuildRsp:
		_temp := make([]*msg.RspUserListItem, 0)
		for _, e := range users {
			_temp = append(_temp, (fn).(FuncUserBuildRsp)(e))
		}
		out = _temp
	}
	return out
}
