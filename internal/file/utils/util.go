package utils

import (
	"fmt"
	"github.com/HaHadaxigua/melancholy/internal/conf"
	"os"
)

func GetMultiFileName(path, filename string) string {
	return fmt.Sprintf("%s%s", path, filename)
}

func GetMultiFilePath(hash string) string{
	return fmt.Sprintf("%s%s", conf.C.Application.TmpFile, hash)
}
// PathExists 检查本地路径是否存在
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
