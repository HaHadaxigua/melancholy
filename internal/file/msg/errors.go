/******
** @date : 2/4/2021 12:44 AM
** @author : zrx
** @description:
******/
package msg

import (
	"github.com/HaHadaxigua/melancholy/internal/file/envir"
	"github.com/HaHadaxigua/melancholy/internal/response"
)

// File
var (
	ErrFileSave             = &response.Err{Code: 10301, Message: envir.FileSaveFailedErrorMsg}
	ErrFileHasExisted       = &response.Err{Code: 10302, Message: envir.FileRepeatedErrorMsg}
	ErrFileNotFound         = &response.Err{Code: 10303, Message: envir.FileNotFoundErrorMsg}
	ErrBadFilename          = &response.Err{Code: 10304, Message: envir.FileBadNameErrorMsg}
	ErrTargetFolderNotExist = &response.Err{Code: 10305, Message: envir.FileTargetFolderNotExistMsg}
	ErrFileUnSupport        = &response.Err{Code: 10306, Message: envir.FileTargetFolderNotExistMsg}
	ErrMergeFileFailed      = &response.Err{Code: 10307, Message: envir.FileMergeFileFailedMsg}
	ErrMergeFileHasExist    = &response.Err{Code: 10308, Message: envir.FileMergeFileHasExistedMsg}
)
