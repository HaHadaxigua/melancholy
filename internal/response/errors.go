package response

import (
	"fmt"
)

//Err 自定义的错误
type Err struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (e *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", e.Code, e.Message, e.Data)
}

func NewErr(err error) *Err {
	return &Err{
		Code:    UnKnown.Code,
		Message: Unknown,
		Data:    err,
	}
}

func (e *Err) AddCause(err error) *Err {
	e.Data = err.Error()
	return e
}

func Ok(data interface{}) *Err {
	return &Err{
		Code:    6,
		Message: OKStr,
		Data:    data,
	}
}

// 错误码设计： [1/2] [xx] [xx]
// 1 为系统错误 2：普通错误
// 第二个xx 表示出错的模块
// 第三个xx 表示具体的出错类型

var (
	OK                = &Err{Code: 6, Message: OKStr}
	UnKnown           = &Err{Code: 12345, Message: Unknown}
	InternalServerErr = &Err{Code: 10001, Message: InternalServerErrorMsg}
	ErrReq            = &Err{Code: 10002, Message: ErrReqMsg}

	// Request
	BadRequest       = &Err{Code: 20001, Message: BadRequestMsg}
	InvalidParamsErr = &Err{Code: 20002, Message: InvalidParamsErrorMsg}

	// Tools
	GenerateSaltErr    = &Err{Code: 10103, Message: GenerateSaltErrorMsg}
	EncryptPasswordErr = &Err{Code: 10104, Message: EncryptPasswordErrorMsg}
)

const (
	// Success
	OKStr = "Success"

	// Unknown
	Unknown = "UnKnown"

	// Default
	InternalServerErrorMsg = "内部服务器错误"

	// request
	ErrReqMsg             string = "不合法的请求构建"
	BadRequestMsg         string = "请求非法"
	InvalidParamsErrorMsg string = "参数非法"

	BindJsonFailedMsg string = "绑定前端数据失败"

	// Tools
	GenerateSaltErrorMsg    string = "生成盐失败"
	EncryptPasswordErrorMsg string = "加密密码失败"
)
