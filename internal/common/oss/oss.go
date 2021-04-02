package oss

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

var AliyunOss MelancholyOSS

type MelancholyOSS interface {
	GetClient() *oss.Client
	ListBucketNames() ([]string, error)
	CreateBucket(bucketName string) *oss.Bucket
	DeleteBucket(bucketName string) error
	UploadBytes(bucketName, objectName string, data []byte) error
	UploadFileStream(objectName, bucketName string, fd *os.File) error

	// 通过流来下载文件,下载文件到缓存
	DownloadFileByStream(bucketName, objectName string) (*bytes.Buffer, error)
}
