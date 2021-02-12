/******
** @date : 2/10/2021 8:35 PM
** @author : zrx
** @description:
******/
package msg

import "github.com/HaHadaxigua/melancholy/internal/basic/model"

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
