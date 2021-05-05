package oss

import (
	"bytes"
	"fmt"
	"github.com/HaHadaxigua/melancholy/internal/conf"
	"github.com/HaHadaxigua/melancholy/internal/consts"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"runtime"
)

var AliyunOss MelancholyOSS

type MelancholyOSS interface {
	GetClient() *oss.Client
	ListBucketNames() ([]string, error)
	CreateBucket(bucketName string) (*oss.Bucket, error)
	DeleteBucket(bucketName string) error
	UploadBytes(bucketName, objectName string, data []byte) error
	UploadString(bucketName, objectName string, content string) error
	UploadFileStream(objectName, bucketName string, fd *os.File) error

	// 通过流来下载文件,下载文件到缓存
	DownloadFileByStream(bucketName, objectName string) (*bytes.Buffer, error)
}


// BuildBucketNameAndAddress 返回bucketName, address 和当前系统
func  BuildBucketNameAndAddress(userID int, filename string) (string,string,string){
	var location string
	switch runtime.GOOS {
	case "linux":
		location = conf.C.Application.LocationUnix
	case "windows":
		location = conf.C.Application.LocationWin
	}

	bucketName := fmt.Sprintf("%s%d", consts.OssBucketGeneratePrefix, userID)
	ossAddress := fmt.Sprintf("https://%s.%s/%s", bucketName, conf.C.Oss.EndPoint, filename)
	return bucketName, ossAddress, location
}