package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/HaHadaxigua/melancholy/internal/conf"
	"github.com/HaHadaxigua/melancholy/internal/file/envir"
	"github.com/HaHadaxigua/melancholy/internal/file/msg"
	"io"
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

// MergeFiles 合并文件的方法
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

// CalcStringHashInSHA 使用sha算法来计算字符串的hash
func CalcStringHashInSHA(s string) string {
	r := sha1.Sum([]byte(s))
	return hex.EncodeToString(r[:])
}

// CalcFileHashInSHA 使用文件名获取文件来计算hash
func CalcFileHashInSHAByName(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	res, err := CalcFileHashInSHA(f)
	if err != nil {
		return "", err
	}
	return res, nil
}

// CalcFileHashInSHA 通过文件指针获取文件来计算Hash
func CalcFileHashInSHA(file io.Reader) (string, error) {
	h := sha1.New()
	_, err := io.Copy(h, file)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
