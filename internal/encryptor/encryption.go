// @Title 用于计算文件hash和加密的包
package encryptor

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"math"
	"os"
)

// todo: 在计算hash时，面对大文件时，只计算部分

// CalcBytesHashInSHA1 计算byte数组的hash, 最通用的方法
func CalcBytesHashInSHA1(data []byte) string {
	r := sha1.Sum(data)
	return hex.EncodeToString(r[:])
}

// CalcStringHashInSHA1 使用sha算法来计算字符串的hash
func CalcStringHashInSHA1(s string) string {
	r := sha1.Sum([]byte(s))
	return hex.EncodeToString(r[:])
}

// CalcSimpleFileHashInSHA1 使用文件名获取文件来计算hash
func CalcSimpleFileHashInSHA1ByName(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()
	res, err := CalcSimpleFileHashInSHA1(f)
	if err != nil {
		return "", err
	}
	return res, nil
}

// CalcSimpleFileHashInSHA1 通过文件指针获取文件来计算Hash
func CalcSimpleFileHashInSHA1(file io.Reader) (string, error) {
	h := sha1.New()
	_, err := io.Copy(h, file)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// CalcLargeFileHashInSHA1ByName 使用文件名获取大文件来计算hash
func CalcLargeFileHashInSHA1ByName(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()
	res, err := CalcLargeFileHashInSHA1(f)
	if err != nil {
		return "", err
	}
	return res, nil
}

// CalcLargeFileHashInSHA1 通过文件指针来计算大文件的hash
func CalcLargeFileHashInSHA1(file *os.File) (string, error) {
	info, err := file.Stat()
	if err != nil {
		return "", err
	}
	if info.Size() < KBSize8 {
		return CalcSimpleFileHashInSHA1(file)
	}
	blocks := uint64(math.Ceil(float64(info.Size()) / float64(KBSize8)))
	hash := sha1.New()
	for i := uint64(0); i < blocks; i++ {
		blockSize := int(math.Min(KBSize8, float64(info.Size()-int64(i*KBSize8))))
		buf := make([]byte, blockSize)
		file.Read(buf)
		hash.Write(buf)
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
