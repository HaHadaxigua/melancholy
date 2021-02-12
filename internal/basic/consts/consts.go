/******
** @date : 2/10/2021 9:06 PM
** @author : zrx
** @description:
******/
package consts

const (
	// User
	UserCreateErrorMsg               string = "用户创建失败"
	UserNameIllegalErrorMsg          string = "名称非法"
	UserPwdIllegalErrorMsg           string = "密码非法"
	UserEmailIllegalErrorMsg         string = "邮箱非法"
	UserNameOrPwdIncorrectlyErrorMsg string = "用户名或密码不正确"
	UserHasExistedErrorMsg           string = "用户已存在"
	UserExitErrorMsg                 string = "已退出"
	UserNotFoundMsg                  string = "找不到用户"

	// Role
	RepeatedRoleMsg   string = "重复的角色"
	RoleNotExistedMsg string = "角色不存在"

	// Auth
	MsgAuthCheckTokenError    string = "认证Token失败"
	MsgAuthCheckTokenTimeout  string = "Token超时"
	MsgAuthAccessTokenIllegal string = "非法Token"
	MsgAuthFailed             string = "认证失败"
)

const (
	UserStatusOK      int = 13 // 用户状态
	UserStatusBlocked int = 93 // 用户被封禁
)

const (
	Admin string = "Admin"
)
