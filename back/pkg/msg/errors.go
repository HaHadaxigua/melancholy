package msg

import (
	"fmt"
)

//Err 自定义的错误
type Err struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Cause   error  `json:"cause"`
}

func (e *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", e.Code, e.Message, e.Cause)
}

func NewErr(code int, msg string, err error) *Err {
	return &Err{
		Code:    code,
		Message: msg,
		Cause:   err,
	}
}

//DecodeErr 解码错误
func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}
	switch typed := err.(type) {
	case *Err:
		if typed.Code == ErrReq.Code {
			typed.Message = ErrReq.Message + "Desc:" + typed.Cause.Error()
		}
	}

	return InternalServerErr.Code, err.Error()
}

// 错误码设计： [1/2] [xx] [xx]
// 1 为系统错误 2：普通错误
// 第二个xx 表示出错的模块
// 第三个xx 表示具体的出错类型

var (
	OK                = &Err{Code: 0, Message: Ok}
	InternalServerErr = &Err{Code: 10001, Message: InternalServerErrorMsg}
	ErrReq            = &Err{Code: 10002, Message: ErrReqMsg}

	// File 模块
	SaveBaseFileErr = &Err{Code: 10201, Message: FileSaveFailedErrorMsg}

	BadRequest = &Err{Code: 2001, Message: BadRequestMsg}
)

const (
	// Success
	Ok = "Success"

	// Default
	InternalServerErrorMsg = "内部服务器错误"

	// request 模块
	ErrReqMsg     string = "不合法的请求构建"
	BadRequestMsg string = "请求非法"

	BindJsonFailedMsg string = "绑定前端数据失败"

	// file 模块
	FileCreatedFailedMsg   string = "文件创建失败"
	FileSaveFailedErrorMsg string = "文件保存失败"
)
