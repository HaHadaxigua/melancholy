/******
** @date : 2/4/2021 12:44 AM
** @author : zrx
** @description:
******/
package msg

import (
	"github.com/HaHadaxigua/melancholy/internal/file/consts"
	"github.com/HaHadaxigua/melancholy/internal/response"
)

// File
var (
	ErrFileSave       = &response.Err{Code: 10301, Message: consts.FileSaveFailedErrorMsg}
	ErrFileHasExisted = &response.Err{Code: 10302, Message: consts.FileRepeatedErrorMsg}
	ErrFileNotFound   = &response.Err{Code: 10303, Message: consts.FileNotFoundErrorMsg}
	ErrBadFilename    = &response.Err{Code: 10304, Message: consts.FileBadNameErrorMsg}
	ErrTargetFolderNotExist = &response.Err{Code: 10305, Message: consts.FileTargetFolderNotExistMsg}
	ErrFileUnSupport =  &response.Err{Code: 10306, Message: consts.FileTargetFolderNotExistMsg}
)
