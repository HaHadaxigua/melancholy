package persistence

import "fmt"

const (
	KB = 1024
	MB = 1048576
	GB = 1073741824
	TB = 1099511627776
	EB = 1125899906842624
)

// FormatFileSize 字节的单位转换 保留两位小数
func FormatFileSize(fileSize int64) (size string) {
	if fileSize < KB {
		return fmt.Sprintf("%.2fB", float64(fileSize)/float64(1))
	} else if fileSize < MB {
		return fmt.Sprintf("%.2fKB", float64(fileSize)/float64(KB))
	} else if fileSize < GB {
		return fmt.Sprintf("%.2fMB", float64(fileSize)/float64(MB))
	} else if fileSize < TB {
		return fmt.Sprintf("%.2fGB", float64(fileSize)/float64(GB))
	} else if fileSize < EB {
		return fmt.Sprintf("%.2fTB", float64(fileSize)/float64(TB))
	} else {
		return fmt.Sprintf("%.2fEB", float64(fileSize)/float64(EB))
	}
}
