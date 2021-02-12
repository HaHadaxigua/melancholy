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

// User 条件过滤器
type ReqUserFilter struct {
	Username string `json:"username"`

	Offset int `json:"offset, omitempty"`
	Limit  int `json:"limit, omitempty"`
}

type ReqRoleCreate struct {
	RoleName string `json:"roleName"`
}

type ReqRoleListFilter struct {
	Fuzzy string `json:"fuzzy"`

	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type ReqPermissionCreate struct {
	PermissionName string `json:"permissionName"`
}
