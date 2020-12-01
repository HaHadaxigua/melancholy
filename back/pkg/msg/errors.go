package msg

import (
	"fmt"
)

//Err 自定义的错误
type Err struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Cause   string `json:"cause"`
}

func (e *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", e.Code, e.Message, e.Cause)
}

func NewErr(code int, msg string, err error) *Err {
	return &Err{
		Code:    code,
		Message: msg,
		Cause:   err.Error(),
	}
}

func (e *Err) AddCause(err error) {
	e.Cause = err.Error()
}

//DecodeErr 解码错误
func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}
	switch typed := err.(type) {
	case *Err:
		if typed.Code == ErrReq.Code {
			typed.Message = ErrReq.Message + "Desc:" + typed.Cause
		}
	}

	return InternalServerErr.Code, err.Error()
}

// 错误码设计： [1/2] [xx] [xx]
// 1 为系统错误 2：普通错误
// 第二个xx 表示出错的模块
// 第三个xx 表示具体的出错类型

var (
	OK                = &Err{Code: 6, Message: Ok}
	InternalServerErr = &Err{Code: 10001, Message: InternalServerErrorMsg}
	ErrReq            = &Err{Code: 10002, Message: ErrReqMsg}

	// 请求相关
	BadRequest       = &Err{Code: 20001, Message: BadRequestMsg}
	InvalidParamsErr = &Err{Code: 20002, Message: InvalidParamsErrorMsg}

	// Tools 模块
	GenerateSaltErr    = &Err{Code: 10103, Message: GenerateSaltErrorMsg}
	EncryptPasswordErr = &Err{Code: 10104, Message: EncryptPasswordErrorMsg}

	// Auth 模块
	AuthCheckTokenErr         = &Err{Code: 10201, Message: AuthCheckTokenErrorMsg}
	AuthCheckTokenTimeoutErr  = &Err{Code: 10202, Message: AuthCheckTokenTimeoutErrorMsg}
	AuthAccessTokenIllegalErr = &Err{Code: 10203, Message: AuthAccessTokenIllegalErrorMsg}
	AuthorizeFailedErr        = &Err{Code: 10204, Message: AuthorizeFailedMsg}

	// File 模块
	FileSaveErr     = &Err{Code: 10301, Message: FileSaveFailedErrorMsg}
	FileRepeatErr   = &Err{Code: 10302, Message: FileRepeatedErrorMsg}
	FileNotFoundErr = &Err{Code: 10303, Message: FileNotFoundErrorMsg}

	// User 模块
	UserCreateErr               = &Err{Code: 10401, Message: UserCreateErrorMsg}
	UserNameIllegalErr          = &Err{Code: 10402, Message: UserNameIllegalErrorMsg} // 用户名非法
	UserPwdIllegalErr           = &Err{Code: 10403, Message: UserPwdIllegalErrorMsg}
	UserEmailIllegalErr         = &Err{Code: 10404, Message: UserEmailIllegalErrorMsg}
	UserNameOrPwdIncorrectlyErr = &Err{Code: 10405, Message: UserNameOrPwdIncorrectlyErrorMsg}
	UserHasExistedErr           = &Err{Code: 10406, Message: UserHasExistedErrorMsg}
	UserExitErr                 = &Err{Code: 10407, Message: UserExitErrorMsg}

	// Role 模块
	RepeatedRoleErr = &Err{Code: 10501, Message: RepeatedRoleMsg}
	RoleNotFoundErr = &Err{Code: 10502, Message: RoleNotExistedMsg}
)

const (
	// Success
	Ok = "Success"

	// Default
	InternalServerErrorMsg = "内部服务器错误"

	// request 模块
	ErrReqMsg             string = "不合法的请求构建"
	BadRequestMsg         string = "请求非法"
	InvalidParamsErrorMsg string = "参数非法"

	BindJsonFailedMsg string = "绑定前端数据失败"

	// Tools 模块
	GenerateSaltErrorMsg    string = "生成盐失败"
	EncryptPasswordErrorMsg string = "加密密码失败"

	// Auth 模块
	AuthCheckTokenErrorMsg         string = "认证Token失败"
	AuthCheckTokenTimeoutErrorMsg  string = "Token超时"
	AuthAccessTokenIllegalErrorMsg string = "非法Token"
	AuthorizeFailedMsg             string = "认证失败"

	// file 模块
	FileCreatedFailedMsg   string = "文件创建失败"
	FileSaveFailedErrorMsg string = "文件保存失败"
	FileRepeatedErrorMsg   string = "文件重名"
	FileNotFoundErrorMsg   string = "文件未找到"

	// User 模块
	UserCreateErrorMsg               string = "用户创建失败"
	UserNameIllegalErrorMsg          string = "名称非法"
	UserPwdIllegalErrorMsg           string = "密码非法"
	UserEmailIllegalErrorMsg         string = "邮箱非法"
	UserNameOrPwdIncorrectlyErrorMsg string = "用户名或密码不正确"
	UserHasExistedErrorMsg           string = "用户已存在"
	UserExitErrorMsg                 string = "已退出"

	// Role
	RepeatedRoleMsg string = "重复的角色"
	RoleNotExistedMsg string = "角色不存在"
)
