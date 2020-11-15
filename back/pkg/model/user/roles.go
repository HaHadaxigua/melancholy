package user

import "github.com/HaHadaxigua/melancholy/pkg/model"

type Roles struct {
	model.Model
	RegionId int64  `json:"regionId"` // 域id
	RoleName string `json:"roleName"` // 名称
	State   int    `json:"status"`   // 状态：-20:逻辑删除；10:正常; 20:无效
}

func (a Roles) TableName() string {
	return "roles"
}
