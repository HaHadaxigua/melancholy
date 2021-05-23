package envir

// 文件类型的ID
var (
	MapFileTypeToID map[string]int = map[string]int{
		FileTypeTxtStr: FileTypeTxtID,
		FileTypeShell:  FileTypeShID,
		FileTypePng:    FileTypePngID,
		FileTypeJpg:    FileTypeJpgID,
		FileTypeMp3:    FileTypeMp3ID,
		FileTypeFlac:   FileTypeFlacID,
		FileTypeMP4:    FileTypeMP4ID,
		FileTypeRMVB:   FileTypeRMVBID,
		FileTypeKMV:    FileTypeKMVID,
	}
	MapFileTypeToStr map[int]string = map[int]string{
		FileTypeTxtID:  FileTypeTxtStr,
		FileTypeShID:   FileTypeShell,
		FileTypePngID:  FileTypePng,
		FileTypeJpgID:  FileTypeJpg,
		FileTypeMp3ID:  FileTypeMp3,
		FileTypeFlacID: FileTypeFlac,
		FileTypeMP4ID:  FileTypeMP4,
		FileTypeRMVBID: FileTypeRMVB,
		FileTypeKMVID:  FileTypeKMV,
	}
)

// getFileSuffixID 获取文件的后缀id
func getFileSuffixID(suffix string) int {
	switch suffix {
	case ".txt":
		return FileTypeTxtID
	case ".sh":
		return FileTypeShID
	case FileTypePng:
		return FileTypePngID
	case FileTypeJpg:
		return FileTypeJpgID
	case FileTypeMp3:
		return FileTypeMp3ID
	case FileTypeFlac:
		return FileTypeFlacID
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
	case FileTypeFlacID, FileTypeMp3ID:
		return TypeMusic
	case FileTypeMP4ID, FileTypeRMVBID, FileTypeKMVID:
		return TypeVideo
	default:
		return TypeOther
	}
}
