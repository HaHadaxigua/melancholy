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

// 需要在本地排除的文件
var (
	ExcludeFiles map[string]struct{} = map[string]struct{}{
		ExcludeFile1: {},
	}
)
