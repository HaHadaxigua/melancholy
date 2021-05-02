/******
** @date : 2/13/2021 11:30 PM
** @author : zrx
** @description:
******/
package envir

const (
	FileUpload = "file"
)

// 文件下载
const (
	ContentType        = "Content-Type"
	ContentDisposition = "Content-Disposition"
	AcceptLength       = "Accept-Length"
)

const (
	// fi
	FileCreatedFailedMsg        string = "文件创建失败"
	FileSaveFailedErrorMsg      string = "文件保存失败"
	FileRepeatedErrorMsg        string = "文件重名"
	FileNotFoundErrorMsg        string = "文件未找到"
	FileBadNameErrorMsg         string = "文件名非法"
	FileTargetFolderNotExistMsg string = "目标文件不存在"
	FileFileUnSupport           string = "文件格式不支持"
	FileMergeFileFailedMsg      string = "合并文件错误"
	FileMergeFileHasExistedMsg  string = "合并文件已存在"

	FilePatchDeleteFailedMsg string = "文件并发删除失败"
)

const (
	RootFileID string = "root"
)

// file types which can be created
const (
	FileTypeTxtID = iota + 1
	FileTypeShID
	FileTypePngID
)

const (
	FileTypeTxtStr = ".txt"
	FileTypeShell  = ".sh"
	FileTypePng    = ".png"
)

// 需要在本地排除的文件
const (
	ExcludeFile1 = ".DS_Store"
)
