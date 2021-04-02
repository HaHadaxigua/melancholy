/******
** @date : 2/13/2021 11:30 PM
** @author : zrx
** @description:
******/
package consts

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
)

const (
	RootFileID string = "root"
)

// file types which can be created
const (
	FileTypeTxtID = iota + 1
	FileTypeShID
)

const (
	FileTypeTxtStr = ".txt"
	FileTypeShell  = ".sh"
)

var (
	MapFileTypeToID map[string]int = map[string]int{
		FileTypeTxtStr: FileTypeTxtID,
		FileTypeShell:  FileTypeShID,
	}
	MapFileTypeToStr map[int]string = map[int]string{
		FileTypeTxtID: FileTypeTxtStr,
		FileTypeShID:  FileTypeShell,
	}
)
