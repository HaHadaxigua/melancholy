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
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

// User 条件过滤器
type ReqUserFilter struct {
	Username string `form:"username"`

	Offset int `form:"offset"`
	Limit  int `form:"limit"`
}

type ReqUserRoleAssociation struct {
	UserID int `json:"userID"`
	RoleID int `json:"roleID"`
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

type ReqPermissionFilter struct {
	Fuzzy string `form:"fuzzy"`

	Offset int `form:"offset"`
	Limit  int `form:"limit"`
}

type ReqRolePermAssociation struct {
	RoleID       int `json:"roleID"`
	PermissionID int `json:"permissionID"`
}
