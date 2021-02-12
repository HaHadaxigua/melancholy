/******
** @date : 2/10/2021 9:06 PM
** @author : zrx
** @description:
******/
package msg

import (
	"github.com/HaHadaxigua/melancholy/internal/basic/consts"
	"github.com/HaHadaxigua/melancholy/internal/global/response"
)

var (
	// Auth
	ErrAuthCheckToken         = &response.Err{Code: 10201, Message: consts.MsgAuthCheckTokenError}
	ErrAuthCheckTokenTimeout  = &response.Err{Code: 10202, Message: consts.MsgAuthCheckTokenTimeout}
	ErrAuthAccessTokenIllegal = &response.Err{Code: 10203, Message: consts.MsgAuthAccessTokenIllegal}
	ErrAuthFailed             = &response.Err{Code: 10204, Message: consts.MsgAuthFailed}

	// User
	ErrUserCreate         = &response.Err{Code: 10401, Message: consts.UserCreateErrorMsg}
	ErrUserNameIllegal    = &response.Err{Code: 10402, Message: consts.UserNameIllegalErrorMsg} // 用户名非法
	ErrUserPwdIllegal     = &response.Err{Code: 10403, Message: consts.UserPwdIllegalErrorMsg}
	ErrUserEmailIllegal   = &response.Err{Code: 10404, Message: consts.UserEmailIllegalErrorMsg}
	ErrUserNameOrPwdWrong = &response.Err{Code: 10405, Message: consts.UserNameOrPwdIncorrectlyErrorMsg}
	ErrUserHasExisted     = &response.Err{Code: 10406, Message: consts.UserHasExistedErrorMsg}
	ErrUserExit           = &response.Err{Code: 10407, Message: consts.UserExitErrorMsg}
	ErrUserNotFound       = &response.Err{Code: 10408, Message: consts.UserNotFoundMsg}

	// Role
	ErrRoleRepeated = &response.Err{Code: 10501, Message: consts.RepeatedRoleMsg}
	ErrRoleNotFound = &response.Err{Code: 10502, Message: consts.RoleNotExistedMsg}
)
