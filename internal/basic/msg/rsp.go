/******
** @date : 2/10/2021 8:35 PM
** @author : zrx
** @description:
******/
package msg

import (
	"github.com/HaHadaxigua/melancholy/internal/basic/model"
	"time"
)

type RspUserList struct {
	List  []*RspUserListItem `json:"list"`
	Total int                `json:"total"`
}

type RspUserListItem struct {
	UserID     int                `json:"userID"`
	UserName   string             `json:"userName"`
	UserEmail  string             `json:"email"`
	UserMobile string             `json:"mobile"`
	Roles      []*RspRoleListItem `json:"roles"`
}

type RspRoleList struct {
	List  []*RspRoleListItem `json:"list"`
	Total int                `json:"total"`
}

type RspRoleListItem struct {
	RoleID      int                 `json:"roleID"`
	RoleName    string              `json:"roleName"`
	Permissions []*model.Permission `json:"permissions"`
}

type RspPermList struct {
	List  []*RspPermListItem `json:"list"`
	Total int                `json:"total"`
}

type RspPermListItem struct {
	PermissionID   int    `json:"permissionID"`
	PermissionName string `json:"permissionName"`
}

type UserInfo struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	Status   int    `json:"status"`
	Avatar   string `json:"avatar"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// RspLogin 登陆后的返回体
type RspLogin struct {
	Token string    `json:"token"`
	User  *UserInfo `json:"user"`
}
