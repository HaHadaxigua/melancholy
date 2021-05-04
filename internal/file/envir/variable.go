package envir

// 文件类型的ID
var (
	MapFileTypeToID map[string]int = map[string]int{
		FileTypeTxtStr: FileTypeTxtID,
		FileTypeShell:  FileTypeShID,
		FileTypePng:    FileTypePngID,
	}
	MapFileTypeToStr map[int]string = map[int]string{
		FileTypeTxtID: FileTypeTxtStr,
		FileTypeShID:  FileTypeShell,
		FileTypePngID: FileTypePng,
	}
)

// getFileSuffixID 获取文件的后缀id
func getFileSuffixID(suffix string) int {
	switch suffix {
	case ".txt":
		return FileTypeTxtID
	case ".sh":
		return FileTypeShID
	case ".png":
		return FileTypePngID
	case ".jpg":
		return FileTypeJpgID
	case ".mp3":
		return FileTypeMp3ID
	default:
		return 0
	}
}

// 需要在本地排除的文件
var (
	ExcludeFiles map[string]struct{} = map[string]struct{}{
		ExcludeFile1: {},
	}
)

// GetFileType 用于支持的数据结构转换
func GetFileType(suffix int) int {
	switch suffix {
	case FileTypeTxtID, FileTypeShID:
		return TypeTxt
	case FileTypePngID:
		return TypePictures
	default:
		return TypeOther
	}
}
