/******
** @date : 2/10/2021 8:35 PM
** @author : zrx
** @description:
******/
package msg

import "mime/multipart"

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ReqRegister struct {
	Username string `json:"username"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

// User 条件过滤器
type ReqUserFilter struct {
	Username string `json:"username"form:"username"`

	Offset int `json:"offset"form:"offset"`
	Limit  int `json:"limit"form:"limit"`
}

type ReqUserRoleAssociation struct {
	UserID int `json:"userID"`
	RoleID int `json:"roleID"`
}

type ReqRoleCreate struct {
	RoleName string `json:"name" form:"name"`
}

type ReqRoleDelete struct {
	RoleID int `json:"id" form:"id"`
}

type ReqRoleFilter struct {
	Fuzzy string `json:"fuzzy" form:"fuzzy"`

	Offset int `json:"offset" form:"offset"`
	Limit  int `json:"limit" form:"limit"`
}

type ReqPermissionCreate struct {
	PermissionName string `json:"name"`
}

type ReqPermissionDelete struct {
	PermissionID int `json:"id"`
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

// ReqSetUserInfo 设置用户信息
type ReqSetUserInfo struct {
	Username string `form:"username" json:"username" gorm:"username"`
	Mobile   string `form:"mobile" json:"mobile" gorm:"mobile"`
	//Password          string `form:"password" json:"password" gorm:"password"`
	OssEndPoint       string `form:"ossEndPoint" json:"ossEndPoint" gorm:"oss_end_point"`
	OssAccessKey      string `form:"ossAccessKey" json:"ossAccessKey" gorm:"oss_access_key"`
	OssAccessSecret   string `form:"ossAccessSecret" json:"ossAccessSecret" gorm:"oss_access_secret"`
	CloudAccessKey    string `form:"cloudAccessKey" json:"cloudAccessKey" gorm:"cloud_access_key"`
	CloudAccessSecret string `form:"cloudAccessSecret" json:"cloudAccessSecret" gorm:"cloud_access_secret"`

	UserID int
}

// ReqUpdateAvatar 处理用户头像上传
type ReqUpdateAvatar struct {
	Data       []byte                `json:"data"`
	FileHeader *multipart.FileHeader `json:"fileHeader"`

	UserID int
}
