package utils

import (
	"fmt"
	"github.com/HaHadaxigua/melancholy/internal/conf"
	"runtime"
)

// GetOSType 获取系统类型
func GetOSType() string {
	var location string
	switch runtime.GOOS {
	case "linux":
		location = conf.C.Application.LocationUnix
	case "windows":
		location = conf.C.Application.LocationWin
	}
	return location
}

// GetMultiFileName 获取分片文件的文件名
func GetMultiFileName(path, filename string) string {
	return fmt.Sprintf("%s%s", path, filename)
}

// GetMultiFilePath 获取分片文件的存储位置
func GetMultiFilePath(hash string) string {
	return fmt.Sprintf("%s%s", conf.C.Application.TmpFile, hash)
}

