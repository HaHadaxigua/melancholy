package utils

import (
	"fmt"
	"github.com/HaHadaxigua/melancholy/internal/conf"
	"github.com/HaHadaxigua/melancholy/internal/file/envir"
	"github.com/HaHadaxigua/melancholy/internal/file/msg"
	"io/ioutil"
	"os"
)

// GetMultiFileName 获取分片文件的文件名
func GetMultiFileName(path, filename string) string {
	return fmt.Sprintf("%s%s", path, filename)
}

// GetMultiFilePath 获取分片文件的存储位置
func GetMultiFilePath(hash string) string {
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

// 合并文件的方法
func MergeFiles(dir, filename string) error {
	megedFile := fmt.Sprintf("%s/%s", dir, filename)
	if PathExists(megedFile) {
		return msg.ErrMergeFileHasExist
	}
	complateFile, err := os.Create(megedFile)
	if err != nil {
		return err
	}
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, file := range files {
		if _, ok := envir.ExcludeFiles[file.Name()]; ok {
			continue
		}
		fileContent, err := ioutil.ReadFile(file.Name())
		if err != nil {
			return err
		}
		complateFile.Write(fileContent)
	}
	defer func() {
		complateFile.Close()
	}()
	return nil
}
