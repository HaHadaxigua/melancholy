package user

import "github.com/HaHadaxigua/melancholy/pkg/model"

type Role struct {
	model.Model
	Name   string `json:"roleName"` // 名称
	Status int    `json:"status"`   // 状态：-20:逻辑删除；10:正常; 20:无效
}

func (a *Role) TableName() string {
	return "roles"
}

type XUserRoles struct {
	ID  int `json:"ID"`
	UserID int `json:"userID"`
	RoleID int `json:"roleID"`
}

func (x *XUserRoles) TableName() string {
	return "user_roles"
}
