/******
** @date : 2/10/2021 8:35 PM
** @author : zrx
** @description:
******/
package msg

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ReqRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type ReqRoleCreate struct {
	RoleName string `json:"roleName"`
}

type ReqRoleFilter struct {
	Fuzzy string `form:"fuzzy"`

	Offset int `form:"offset"`
	Limit  int `form:"limit"`
}

type ReqPermissionCreate struct {
	PermissionName string `json:"permissionName"`
}

// User 条件过滤器
type ReqUserFilter struct {
	Username string `json:"username"`

	Offset int `json:"offset, omitempty"`
	Limit  int `json:"limit, omitempty"`
}
