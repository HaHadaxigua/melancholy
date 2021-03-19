/******
** @date : 2/13/2021 11:30 PM
** @author : zrx
** @description:
******/
package consts

const (
	// fi
	FileCreatedFailedMsg        string = "文件创建失败"
	FileSaveFailedErrorMsg      string = "文件保存失败"
	FileRepeatedErrorMsg        string = "文件重名"
	FileNotFoundErrorMsg        string = "文件未找到"
	FileBadNameErrorMsg         string = "文件名非法"
	FileTargetFolderNotExistMsg string = "目标文件不存在"
)

const (
	RootFileID string = "root"
)

const (
	UPLOAD string = "upload"
)


// file types which can be created
const (
	FileType = iota
	FileTypeTxt
)