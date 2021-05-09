package persistence

import (
	"bufio"
	"fmt"
	"github.com/HaHadaxigua/melancholy/internal/file/envir"
	"github.com/HaHadaxigua/melancholy/internal/file/msg"
	"io/ioutil"
	"os"
)

// Persistence 文件持久化相关，处理用户的本地文件
//type Persistence interface {
//	// 保存小文件：将文件持久化 p: 文件名，文件位置，要保存的数据
//	SaveSimpleFile(filename, location string, data []byte) error
//}
//
//// ResourceManager: 资源管理模块 处理本地文件持久化相关工作
//type ResourceManager struct {
//}
//
//func NewResourceManager() *ResourceManager {
//	return &ResourceManager{}
//}

// ResourceManager 创建文件夹, 给定路径和文件名创建文件夹
func CreateFolder(name, location string) {

}

// SaveFile 保存文件，需要给出文件名、文件位置、以及数据
func SaveSimpleFile(filename, location string, data []byte) error {
	file, err := os.Create(location + filename)
	if err != nil {
		return err
	}
	defer func() {
		file.Close()
	}()
	bufferWritter := bufio.NewWriter(file)
	bytesWritten, err := bufferWritter.Write(data)
	if err != nil {
		return err
	}
	if bytesWritten < 1 {
		return fmt.Errorf("write file from buffer to hard-disk-dirve")
	}
	//unflushedBufferedSize := bufferWritter.Buffered()
	if err := bufferWritter.Flush(); err != nil {
		return err
	}
	bufferWritter.Reset(bufferWritter) // 丢弃没缓存的内容
	return nil
}

// MergeFiles 将分片了的文件进行合并
func MergeFiles(dir, filename string) error {
	megedFile := fmt.Sprintf("%s/%s", dir, filename)
	exist, err := IsPathExist(megedFile)
	if err != nil {
		return err
	}
	if exist {
		return msg.ErrFileHasExisted // 文件已存在
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

// IsPathExist 判断文件夹或文件是否存在
func IsPathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//  CreateFolderIfNotExist 创建文件夹 如果不存在
func CreateFolderIfNotExist(path string) error {
	exist, err := IsPathExist(path)
	if err != nil {
		return err
	}
	if !exist {
		return os.Mkdir(path, os.ModePerm) // 创建文件夹
	}
	return nil
}
